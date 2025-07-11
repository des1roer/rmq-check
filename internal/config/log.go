package config

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Настройка logrus для вывода в формате JSON
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout) // Логи в stdout
}
