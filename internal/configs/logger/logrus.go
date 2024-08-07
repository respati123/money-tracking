package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type CustomLogger struct {
	logger *logrus.Logger
	module string
}

func NewLogger(viper *viper.Viper) *CustomLogger {
	log := logrus.New()

	log.SetLevel(logrus.Level(viper.GetInt32("LOG_LEVEL")))
	log.SetFormatter(&logrus.JSONFormatter{
		PrettyPrint:       true,
		DisableHTMLEscape: true,
	})

	return &CustomLogger{logger: log}
}

func (l *CustomLogger) Module(instance string) *CustomLogger {
	return &CustomLogger{
		logger: l.logger,
		module: instance,
	}
}

func (l *CustomLogger) Logger() *logrus.Logger {
	return l.logger
}

func (l *CustomLogger) ErrorWithFields(ctx *gin.Context, errorMessage string, err error) {
	l.logger.WithFields(logrus.Fields{
		"module":   l.module,
		"trace_id": ctx.GetString("trace_id"),
		"error":    errorMessage,
	}).Error(err.Error())
}
