package log

// ELogWriteFailed is an error that occurs when writing the log destination failed. (e.g. the disk is full)
const ELogWriteFailed = "E_LOG_WRITE_FAILED"

// ELogRotateFailed is an error that happens when a log rotation is desired but failed for some reason.
const ELogRotateFailed = "E_LOG_ROTATE_FAILED"

// ELogFileOpenFailed indicates that the log file could not be opened.
const ELogFileOpenFailed = "E_LOG_FILE_OPEN_FAILED"

// EUnknownError is a non-conformant error.
const EUnknownError = "E_UNKNOWN_ERROR"
