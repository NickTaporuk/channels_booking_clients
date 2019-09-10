package logger

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

type LocalLogger struct {
	logger *logrus.Logger
	file   *os.File
}

func (l *LocalLogger) File() *os.File {
	return l.file
}

func (l *LocalLogger) SetFile(file *os.File) {
	l.file = file
}

func (l *LocalLogger) Logger() *logrus.Logger {
	return l.logger
}

func (l *LocalLogger) SetLogger(logger *logrus.Logger) {
	l.logger = logger
}

func (l *LocalLogger) Close() {
	var err = l.file.Close()

	if err != nil {
		log.Fatal(err)
	}
}

// Init init logger
func (l *LocalLogger) Init(data map[string]string) error {
	var (
		logger       *logrus.Logger
		loggingLevel logrus.Level
		loggerLevel  = data["level"]
		err          error
	)

	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to determine working directory: %s", err)
	}
	runID := time.Now().Format("run-2006-01-02-15-04-05")
	logLocation := filepath.Join(cwd, runID+".log")

	logger = logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	// Only logger the warning severity or above.
	loggingLevel, err = logrus.ParseLevel(loggerLevel)
	if err != nil {
		return err
	}

	logger.SetLevel(loggingLevel)

	f, err := os.OpenFile(logLocation, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	mw := io.MultiWriter(os.Stdout, f)

	logger.SetOutput(mw)

	l.SetFile(f)
	l.SetLogger(logger)

	return nil
}
