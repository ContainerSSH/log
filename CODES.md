# Message / error codes

| Code | Explanation |
|------|-------------|
| `LOG_FILE_OPEN_FAILED` | ContainerSSH failed to open the specified log file. |
| `LOG_ROTATE_FAILED` | ContainerSSH cannot rotate the logs as requested because of an underlying error. |
| `LOG_WRITE_FAILED` | ContainerSSH cannot write to the specified log file. This usually happens because the underlying filesystem is full or the log is located on a non-local storage (e.g. NFS), which is not supported. |
| `TEST` | This is message that should only be seen in unit and component tests, never in production. |
| `UNKNOWN_ERROR` | This is an untyped error. If you see this in a log that is a bug and should be reported. |

