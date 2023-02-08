package zap_logger

import "go.uber.org/zap"

type Logger struct {
	LoggingService *zap.Logger
}

func NewLogger() (*Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()

	return &Logger{
		LoggingService: logger,
	}, nil
}

func (l *Logger) LogError(msg string, err error) {
	l.LoggingService.Error(msg, zap.Error(err))
}

func (l *Logger) LogInfo(msg string) {
	l.LoggingService.Info(msg)
}

func (l *Logger) LogWarn(msg string) {
	l.LoggingService.Warn(msg)
}

func (l *Logger) LogDebug(msg string) {
	l.LoggingService.Debug(msg)
}
