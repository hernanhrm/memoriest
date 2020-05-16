package logger

import (
	"fmt"
	"github.com/orandin/lumberjackrus"
	log "github.com/sirupsen/logrus"
	"time"
)

func InitLogger(path string, pretty bool) {
	formatter := &log.JSONFormatter{
		TimestampFormat: time.RFC3339,
		PrettyPrint:     pretty,
	}

	log.SetReportCaller(true)
	log.SetFormatter(formatter)

	hook, err := lumberjackrus.NewHook(
		settingsLogFile(path, "general"),
		log.InfoLevel,
		formatter,
		&lumberjackrus.LogFileOpts{
			log.InfoLevel:  settingsLogFile(path, "info"),
			log.ErrorLevel: settingsLogFile(path, "error"),
			log.WarnLevel:  settingsLogFile(path, "warn"),
			log.TraceLevel: settingsLogFile(path, "trace"),
		},
	)

	if err != nil {
		log.Fatalf("unable to set lumberjack hook %v", err)
	}

	log.AddHook(hook)
}

func settingsLogFile(path, level string) *lumberjackrus.LogFile {
	return &lumberjackrus.LogFile{
		Filename:   fmt.Sprintf("%s/%s.log", path, level),
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 1,
		LocalTime:  true,
		Compress:   true,
	}
}
