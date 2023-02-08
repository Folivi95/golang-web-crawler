package ports

type Logger interface {
	LogError(msg string, err error)
	Log(msg string)
	LogInfo(msg string)
	LogDebug(msg string)
	LogWarn(msg string)
}
