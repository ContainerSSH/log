# Error/message codes

| Code | Explanation |
|------|-------------|
| `E_LOG_WRITE_FAILED` | ContainerSSH failed to write to the log output. This is fatal error and will cause ContainerSSH to stop. |
| `E_LOG_ROTATE_FAILED` | ContainerSSH failed to close and reopen the logs. This is a fatal error and will cause ContainerSSH to stop. |
| `E_LOG_FILE_OPEN_FAILED` | ContainerSSH failed to open a log file. This is a fatal error and will cause ContainerSSH to stop. |
| `E_UNKNOWN_ERROR` | This is a generic error without any specificity. Please report this as a bug. |
