package log

import (
	"os"
	"sync"
)

func newFileWriter(filename string, format Format) (Writer, error) {
	lock := &sync.Mutex{}
	fh, err := openLogFile(filename)
	if err != nil {
		return nil, err
	}
	return &fileWriter{
		fileHandleWriter: newFileHandleWriter(fh, format, lock),
		filename:         filename,
		lock:             lock,
		fh:               fh,
	}, nil
}

// fileWriter inherits the write method from fileHandleWriter and writes to a file. It adds the ability to rotate
// logs and close the log file.
type fileWriter struct {
	*fileHandleWriter

	filename string
	lock     *sync.Mutex
	fh       *os.File
}

func (f *fileWriter) Rotate() error {
	f.lock.Lock()
	defer f.lock.Unlock()
	fh, err := openLogFile(f.filename)
	if err != nil {
		return Wrap(
			err,
			ELogRotateFailed,
			"failed to rotate logs",
		)
	}
	oldFh := f.fh
	f.fh = fh
	f.fileHandleWriter.fh = fh
	if err := oldFh.Close(); err != nil {
		return Wrap(
			err,
			ELogRotateFailed,
			"failed to close old log file",
		)
	}
	return nil
}

func (f *fileWriter) Close() error {
	return f.fh.Close()
}

func openLogFile(filename string) (*os.File, error) {
	fh, err := os.OpenFile(filename, os.O_CREATE&os.O_APPEND, 0644)
	if err != nil {
		return nil, Wrap(err, ELogFileOpenFailed, "failed to open log file %s", filename)
	}
	return fh, nil
}
