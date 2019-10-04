package logger

import (
	"errors"
	"github.com/NickTaporuk/channels_booking_clients/config"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

type LocalLogger struct {
	logger  *logrus.Logger
	file    *os.File
	writers []io.Writer
}

func (l *LocalLogger) Writers() []io.Writer {
	return l.writers
}

func (l *LocalLogger) SetWriters(writers []io.Writer) {
	if len(l.writers) != 0 {
		l.writers = append(l.writers, writers...)
	} else {
		l.writers = writers
	}
}

func (l *LocalLogger) SetWriter(writer io.Writer) {
	l.writers = append(l.writers, writer)
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
func (l *LocalLogger) Init(data map[string]string, cfg *config.LoggerConfiguration) error {

	loggerLevel, ok := data["level"]
	if !ok {
		return errors.New("parameter level of logger initial data is required")
	}

	logger := logrus.New()

	// Log as JSON instead of the default ASCII formatter.
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	// Only logger the warning severity or above.
	loggingLevel, err := logrus.ParseLevel(loggerLevel)
	if err != nil {
		return err
	}

	logger.SetLevel(loggingLevel)

	err = l.initLogFileLogging(cfg)
	if err != nil {
		return err
	}

	l.initStdoutLogging(cfg)

	writers := l.Writers()
	mw := io.MultiWriter(writers...)

	logger.SetOutput(mw)

	l.SetLogger(logger)

	return nil
}

func (l *LocalLogger) initStdoutLogging(cfg *config.LoggerConfiguration) {
	if cfg.StdOut {
		l.SetWriter(os.Stdout)
	}
}

func (l *LocalLogger) initLogFileLogging(cfg *config.LoggerConfiguration) error {
	if cfg.File {
		var logLocation string
		var err error
		runID := time.Now().Format("cbg-2006-01-02-15-04-05")

		if cfg.FilePath == "" {
			cwd, err := os.Getwd()
			if err != nil {
				log.Fatalf("Failed to determine working directory: %s", err)
			}
			logLocation = filepath.Join(cwd, runID+".log")

		} else {
			filePath, err := filepath.Abs(cfg.FilePath)
			if err != nil {
				return err
			}

			logLocation = filepath.Join(filePath, runID+".log")

		}

		f, err := os.OpenFile(logLocation, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			return err
		}

		l.SetWriter(f)
	}

	return nil
}
