package configs

import "github.com/sirupsen/logrus"

func SetupLogging(level string) {
	logLevel, err := logrus.ParseLevel(level)
	if err != nil {
		logrus.Warn("Invalid log level, defaulting to info")
		logLevel = logrus.InfoLevel
	}

	logrus.SetLevel(logLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	switch level {
	case "info":
		logrus.Info("Será logado: info e error")
	case "debug":
		logrus.Debug("Será logado: info, debug e error")
	case "error":
		logrus.Error("Será logado: error")
	default:
		logrus.Info("Será logado: info e error")
	}

	logrus.Debug("This is a debug message")
	logrus.Info("This is an info message")
	logrus.Error("This is an error message")
}
