package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func Init(debug bool) {
	// Форматируем вывод с полным timestamp
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	// Если debug = true — ставим уровень Debug, иначе Info
	if debug {
		Log.SetLevel(logrus.DebugLevel)
	} else {
		Log.SetLevel(logrus.InfoLevel)
	}

	// Выводим логи в консоль (stdout)
	Log.SetOutput(os.Stdout)
}
