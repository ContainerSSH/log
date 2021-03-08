package log

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"testing"
)

// Config describes the logging settings.
type Config struct {
	// Level describes the minimum level to log at
	Level Level `json:"level" yaml:"level" default:"5"`

	// Format describes the log message format
	Format Format `json:"format" yaml:"format" default:"ljson"`

	// Destination is the target to write the log messages to.
	Destination Destination `json:"destination" yaml:"destination" default:"stdout"`

	// File is the log file to write to if Destination is set to "file".
	File string `json:"file" yaml:"file" default:"/var/log/containerssh/containerssh.log"`

	// Syslog configures the syslog destination.
	Syslog SyslogConfig `json:"syslog" yaml:"syslog"`

	// T is the Go test for logging purposes.
	T *testing.T `json:"-" yaml:"-"`

	// Stdout is the standard output used by the DestinationStdout destination.
	Stdout io.Writer `json:"-" yaml:"-"`
}

// Validate validates the log configuration.
func (c *Config) Validate() error {
	if err := c.Level.Validate(); err != nil {
		return err
	}
	if err := c.Format.Validate(); err != nil {
		return err
	}
	if err := c.Destination.Validate(); err != nil {
		return err
	}
	if c.Destination == DestinationTest && c.T == nil {
		return fmt.Errorf("test log destination selected but no test case provided")
	}
	return nil
}

// region Level

// Level syslog-style log level identifiers
type Level int8

// Supported values for Level
const (
	LevelDebug     Level = 7
	LevelInfo      Level = 6
	LevelNotice    Level = 5
	LevelWarning   Level = 4
	LevelError     Level = 3
	LevelCritical  Level = 2
	LevelAlert     Level = 1
	LevelEmergency Level = 0
)

// UnmarshalJSON decodes a JSON level string to a level type.
func (level *Level) UnmarshalJSON(data []byte) error {
	var levelString LevelString
	if err := json.Unmarshal(data, &levelString); err != nil {
		unmarshalError := &json.UnmarshalTypeError{}
		if errors.As(err, &unmarshalError) {
			type levelAlias Level
			var l levelAlias
			if err = json.Unmarshal(data, &l); err != nil {
				return err
			}
			*level = Level(l)
		}
		return err
	}
	l, err := levelString.ToLevel()
	if err != nil {
		return err
	}
	*level = l
	return nil
}

// MarshalJSON marshals a level number to a JSON string
func (level Level) MarshalJSON() ([]byte, error) {
	levelString, err := level.Name()
	if err != nil {
		return nil, err
	}
	return json.Marshal(levelString)
}

// UnmarshalYAML decodes a YAML level string to a level type.
func (level *Level) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var levelString LevelString
	if err := unmarshal(&levelString); err != nil {
		return err
	}
	l, err := levelString.ToLevel()
	if err != nil {
		type levelAlias Level
		var l2 levelAlias
		if err2 := unmarshal(&l2); err2 != nil {
			return err
		}
		*level = Level(l2)
		return nil
	}
	*level = l
	return nil
}

// MarshalYAML creates a YAML text representation from a numeric level
func (level Level) MarshalYAML() (interface{}, error) {
	return level.Name()
}

// Name Convert the int level to the string representation
func (level *Level) Name() (LevelString, error) {
	switch *level {
	case LevelDebug:
		return LevelDebugString, nil
	case LevelInfo:
		return LevelInfoString, nil
	case LevelNotice:
		return LevelNoticeString, nil
	case LevelWarning:
		return LevelWarningString, nil
	case LevelError:
		return LevelErrorString, nil
	case LevelCritical:
		return LevelCriticalString, nil
	case LevelAlert:
		return LevelAlertString, nil
	case LevelEmergency:
		return LevelEmergencyString, nil
	}
	return "", fmt.Errorf("invalid log level (%d)", level)
}

// MustName Convert the int level to the string representation and panics if the name is not valid.
func (level Level) MustName() LevelString {
	name, err := level.Name()
	if err != nil {
		panic(err)
	}
	return name
}

// String Convert the int level to the string representation. Panics if the level is not valid.
func (level Level) String() string {
	return string(level.MustName())
}

// Validate if the log level has a valid value
func (level Level) Validate() error {
	if level < LevelEmergency || level > LevelDebug {
		return fmt.Errorf("invalid log level (%d)", level)
	}
	return nil
}

// endregion

// region LevelString

// LevelString is a type for supported log level strings
type LevelString string

// List of valid string values for log levels
const (
	LevelDebugString     LevelString = "debug"
	LevelInfoString      LevelString = "info"
	LevelNoticeString    LevelString = "notice"
	LevelWarningString   LevelString = "warning"
	LevelErrorString     LevelString = "error"
	LevelCriticalString  LevelString = "crit"
	LevelAlertString     LevelString = "alert"
	LevelEmergencyString LevelString = "emerg"
)

// ToLevel convert the string level to the int representation
func (level LevelString) ToLevel() (Level, error) {
	switch level {
	case LevelDebugString:
		return LevelDebug, nil
	case LevelInfoString:
		return LevelInfo, nil
	case LevelNoticeString:
		return LevelNotice, nil
	case LevelWarningString:
		return LevelWarning, nil
	case LevelErrorString:
		return LevelError, nil
	case LevelCriticalString:
		return LevelCritical, nil
	case LevelAlertString:
		return LevelAlert, nil
	case LevelEmergencyString:
		return LevelEmergency, nil
	}
	return -1, fmt.Errorf("invalid log level (%s)", level)
}

// endregion

// region Format

// Format is the logging format to use.
//swagger:enum
type Format string

const (
	// FormatLJSON is a newline-delimited JSON log format.
	FormatLJSON Format = "ljson"
	// FormatText prints the logs as plain text.
	FormatText Format = "text"
)

// Validate returns an error if the format is invalid.
func (format Format) Validate() error {
	switch format {
	case FormatLJSON:
	case FormatText:
	default:
		return fmt.Errorf("invalid log format: %s", format)
	}
	return nil
}

// endregion

// region Destination

// Destination is the output to write to.
//swagger:enum
type Destination string

const (
	// DestinationStdout is writing log messages to the standard output.
	DestinationStdout Destination = "stdout"
	// DestinationFile is writing the log messages to a file.
	DestinationFile Destination = "file"
	// DestinationSyslog is writing log messages to syslog.
	DestinationSyslog Destination = "syslog"
	// DestinationTest writes the logs to the *testing.T facility.
	DestinationTest Destination = "test"
)

// Validate validates the output target.
func (o Destination) Validate() error {
	switch o {
	case DestinationStdout:
	case DestinationFile:
	case DestinationSyslog:
	case DestinationTest:
	default:
		return fmt.Errorf("invalid destination: %s", o)
	}
	return nil
}

// endregion

// region Syslog

// Priority
type Facility int

const (
	// FacilityKern are kernel messages.
	FacilityKern Facility = 0
	// FacilityUser are user level messages.
	FacilityUser Facility = 1
	// FacilityMail are user mail log messages.
	FacilityMail Facility = 2
	// FacilityDaemon are daemon messages.
	FacilityDaemon Facility = 3
	// FacilityAuth are authentication messages.
	FacilityAuth Facility = 4
	// FacilitySyslog are syslog-specific messages.
	FacilitySyslog Facility = 5
	// FacilityLPR are printer messages.
	FacilityLPR Facility = 6
	// FacilityNews are news messages.
	FacilityNews Facility = 7
	// FacilityUUCP are UUCP subsystem messages.
	FacilityUUCP Facility = 8
	// FacilityCron are clock daemon messages.
	FacilityCron Facility = 9
	// FacilityAuthPriv are security/authorization messages.
	FacilityAuthPriv Facility = 10
	// FacilityFTP are FTP daemon messages.
	FacilityFTP Facility = 11
	// FacilityNTP are network time daemon messages.
	FacilityNTP Facility = 12
	// FacilityLogAudit are log audit messages.
	FacilityLogAudit Facility = 13
	// FacilityLogAlert are log alert messages.
	FacilityLogAlert Facility = 14
	// FacilityClock are clock daemon messages.
	FacilityClock Facility = 15

	// FacilityLocal0 are locally administered messages.
	FacilityLocal0 Facility = 16
	// FacilityLocal1 are locally administered messages.
	FacilityLocal1 Facility = 17
	// FacilityLocal2 are locally administered messages.
	FacilityLocal2 Facility = 18
	// FacilityLocal3 are locally administered messages.
	FacilityLocal3 Facility = 19
	// FacilityLocal4 are locally administered messages.
	FacilityLocal4 Facility = 20
	// FacilityLocal5 are locally administered messages.
	FacilityLocal5 Facility = 21
	// FacilityLocal6 are locally administered messages.
	FacilityLocal6 Facility = 22
	// FacilityLocal7 are locally administered messages.
	FacilityLocal7 Facility = 23
)

// Validate checks if the facility is valid.
func (f Facility) Validate() error {
	if _, ok := facilityToName[f]; !ok {
		return fmt.Errorf("invalid facility: %d", f)
	}
	return nil
}

// Name returns the facility name.
func (f Facility) Name() (FacilityString, error) {
	if name, ok := facilityToName[f]; ok {
		return name, nil
	}
	return "", fmt.Errorf("invalid facility: %d", f)
}

// MustName is identical to Name but panics if the facility is invalid.
func (f Facility) MustName() FacilityString {
	name, err := f.Name()
	if err != nil {
		panic(err)

	}
	return name
}

// FacilityString are facility names.
type FacilityString string

const (
	// FacilityStringKern are kernel messages.
	FacilityStringKern FacilityString = "kern"
	// FacilityStringUser are user level messages.
	FacilityStringUser FacilityString = "user"
	// FacilityStringMail are user mail log messages.
	FacilityStringMail FacilityString = "mail"
	// FacilityStringDaemon are daemon messages.
	FacilityStringDaemon FacilityString = "daemon"
	// FacilityStringAuth are authentication messages.
	FacilityStringAuth FacilityString = "auth"
	// FacilityStringSyslog are syslog-specific messages.
	FacilityStringSyslog FacilityString = "syslog"
	// FacilityStringLPR are printer messages.
	FacilityStringLPR FacilityString = "lpr"
	// FacilityStringNews are news messages.
	FacilityStringNews FacilityString = "news"
	// FacilityStringUUCP are UUCP subsystem messages.
	FacilityStringUUCP FacilityString = "uucp"
	// FacilityStringCron are clock daemon messages.
	FacilityStringCron FacilityString = "cron"
	// FacilityStringAuthPriv are security/authorization messages.
	FacilityStringAuthPriv FacilityString = "authpriv"
	// FacilityStringFTP are FTP daemon messages.
	FacilityStringFTP FacilityString = "ftp"
	// FacilityStringNTP are network time daemon messages.
	FacilityStringNTP FacilityString = "ntp"
	// FacilityStringLogAudit are log audit messages.
	FacilityStringLogAudit FacilityString = "logaudit"
	// FacilityStringLogAlert are log alert messages.
	FacilityStringLogAlert FacilityString = "logalert"
	// FacilityStringClock are clock daemon messages.
	FacilityStringClock FacilityString = "clock"

	// FacilityStringLocal0 are locally administered messages.
	FacilityStringLocal0 FacilityString = "local0"
	// FacilityStringLocal1 are locally administered messages.
	FacilityStringLocal1 FacilityString = "local1"
	// FacilityStringLocal2 are locally administered messages.
	FacilityStringLocal2 FacilityString = "local2"
	// FacilityStringLocal3 are locally administered messages.
	FacilityStringLocal3 FacilityString = "local3"
	// FacilityStringLocal4 are locally administered messages.
	FacilityStringLocal4 FacilityString = "local4"
	// FacilityStringLocal5 are locally administered messages.
	FacilityStringLocal5 FacilityString = "local5"
	// FacilityStringLocal6 are locally administered messages.
	FacilityStringLocal6 FacilityString = "local6"
	// FacilityStringLocal7 are locally administered messages.
	FacilityStringLocal7 FacilityString = "local7"
)

// Validate validates the facility string.
func (s FacilityString) Validate() error {
	if _, ok := nameToFacility[s]; !ok {
		return fmt.Errorf("invalid facility: %s", s)
	}
	return nil
}

// Number returns the facility number.
func (s FacilityString) Number() (Facility, error) {
	if val, ok := nameToFacility[s]; ok {
		return val, nil
	}
	return Facility(-1), fmt.Errorf("invalid facility: %s", s)
}

// MustNumber is identical to Number but panics instead of returning errors
func (s FacilityString) MustNumber() Facility {
	n, err := s.Number()
	if err != nil {
		panic(err)
	}
	return n
}

var facilityToName = map[Facility]FacilityString{
	FacilityKern:     FacilityStringKern,
	FacilityUser:     FacilityStringUser,
	FacilityMail:     FacilityStringMail,
	FacilityDaemon:   FacilityStringDaemon,
	FacilityAuth:     FacilityStringAuth,
	FacilitySyslog:   FacilityStringSyslog,
	FacilityLPR:      FacilityStringLPR,
	FacilityNews:     FacilityStringNews,
	FacilityUUCP:     FacilityStringUUCP,
	FacilityCron:     FacilityStringCron,
	FacilityAuthPriv: FacilityStringAuthPriv,
	FacilityFTP:      FacilityStringFTP,
	FacilityNTP:      FacilityStringNTP,
	FacilityLogAudit: FacilityStringLogAudit,
	FacilityLogAlert: FacilityStringLogAlert,
	FacilityClock:    FacilityStringClock,

	FacilityLocal0: FacilityStringLocal0,
	FacilityLocal1: FacilityStringLocal1,
	FacilityLocal2: FacilityStringLocal2,
	FacilityLocal3: FacilityStringLocal3,
	FacilityLocal4: FacilityStringLocal4,
	FacilityLocal5: FacilityStringLocal5,
	FacilityLocal6: FacilityStringLocal6,
	FacilityLocal7: FacilityStringLocal7,
}

var nameToFacility = map[FacilityString]Facility{
	FacilityStringKern:     FacilityKern,
	FacilityStringUser:     FacilityUser,
	FacilityStringMail:     FacilityMail,
	FacilityStringDaemon:   FacilityDaemon,
	FacilityStringAuth:     FacilityAuth,
	FacilityStringSyslog:   FacilitySyslog,
	FacilityStringLPR:      FacilityLPR,
	FacilityStringNews:     FacilityNews,
	FacilityStringUUCP:     FacilityUUCP,
	FacilityStringCron:     FacilityCron,
	FacilityStringAuthPriv: FacilityAuthPriv,
	FacilityStringFTP:      FacilityFTP,
	FacilityStringNTP:      FacilityNTP,
	FacilityStringLogAudit: FacilityLogAudit,
	FacilityStringLogAlert: FacilityLogAlert,
	FacilityStringClock:    FacilityClock,

	FacilityStringLocal0: FacilityLocal0,
	FacilityStringLocal1: FacilityLocal1,
	FacilityStringLocal2: FacilityLocal2,
	FacilityStringLocal3: FacilityLocal3,
	FacilityStringLocal4: FacilityLocal4,
	FacilityStringLocal5: FacilityLocal5,
	FacilityStringLocal6: FacilityLocal6,
	FacilityStringLocal7: FacilityLocal7,
}

// SyslogConfig is the configuration for syslog logging.
//goland:noinspection GoVetStructTag
type SyslogConfig struct {
	// Destination is the socket to send logs to. Can be a local path to unix sockets as well as UDP destinations.
	Destination string `json:"destination" yaml:"destination" default:"/dev/log"`
	// Facility logs to the specified syslog facility.
	Facility FacilityString `json:"facility" yaml:"facility" default:"auth"`
	// Tag is the syslog tag to log with.
	Tag string `json:"tag" yaml:"tag" default:"ContainerSSH"`
	// Pid is a setting to append the current process ID to the tag.
	Pid bool `json:"pid" yaml:"pid" default:"false"`

	// connection is the connection to the Syslog server. Internal usage only.
	connection net.Conn `json:"-" yaml:"-"`
	// tag is the real syslog tag for the message
	tag string `json:"-" yaml:"-"`
}

// Validate validates the syslog configuration
func (c *SyslogConfig) Validate() error {
	destination := "/dev/log"
	if c.Destination != "" {
		destination = c.Destination
	}
	if strings.HasPrefix(c.Destination, "/") {
		connection, err := net.Dial("unix", destination)
		if err != nil {
			connection, err = net.Dial("unixgram", destination)
			if err != nil {
				return fmt.Errorf("failed to open UNIX socket to %s (%w)", c.Destination, err)
			}
		}
		c.connection = connection
	} else {
		connection, err := net.Dial("udp", destination)
		if err != nil {
			return fmt.Errorf("failed to open UDP socket to %s (%w)", c.Destination, err)
		}
		c.connection = connection
	}
	if err := c.Facility.Validate(); err != nil {
		return err
	}
	c.tag = "ContainerSSH"
	if c.Tag != "" {
		c.tag = c.Tag
	}
	return nil
}

// endregion
