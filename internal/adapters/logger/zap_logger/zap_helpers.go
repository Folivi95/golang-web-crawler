package zap_logger

import "go.uber.org/zap"

type LoggingService struct {
	LoggingService *zap.Logger
}

func NewLogger() (*LoggingService, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	defer logger.Sync()

	return &LoggingService{
		LoggingService: logger,
	}, nil
}

func (l *LoggingService) LogError(msg string, err error) {
	l.LoggingService.Error(msg, zap.Error(err))
}

func (l *LoggingService) Log(msg string) {
	l.Log(msg)
}

func (l *LoggingService) LogInfo(msg string) {
	l.LoggingService.Info(msg)
}

func (l *LoggingService) LogWarn(msg string) {
	l.LoggingService.Warn(msg)
}

func (l *LoggingService) LogDebug(msg string) {
	l.LoggingService.Debug(msg)
}
