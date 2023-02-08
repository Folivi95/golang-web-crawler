package ports

type Logger interface {
	LogError(msg string, err error)
	LogInfo(msg string)
	LogDebug(msg string)
	LogWarn(msg string)
}
