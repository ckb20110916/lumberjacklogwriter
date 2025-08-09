package lumberjacklogwriter

import (
	"io"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/natefinch/lumberjack.v2"
)

func New(folder, filename string, maxBackups, maxSize, maxAge int) (logWriter io.Writer) {
	err := os.MkdirAll(folder, 0755)
	if err != nil {
		return
	}
	pathname := filepath.ToSlash(path.Join(folder, filename))
	logWriter = &lumberjack.Logger{
		LocalTime:  true,
		Filename:   pathname,
		MaxBackups: maxBackups,
		MaxSize:    maxSize,
		MaxAge:     maxAge,
		Compress:   true,
	}
	return
}
