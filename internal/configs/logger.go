package configs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func SetupLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetOutput(os.Stdout)
	return logger
}
