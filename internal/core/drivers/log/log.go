package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/zput/zxcTool/ztLog/zt_formatter"
)

func Init() {
	InitLogrus()
}

func InitLogrus() {
	var formater = &zt_formatter.ZtFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
		},
	}

	l := logrus.WithFields(logrus.Fields{})

	l.Logger.SetReportCaller(true)
	l.Logger.SetFormatter(formater)

	l.Logger.SetLevel(GetLoggerLevel())
}

func GetLoggerLevel() logrus.Level {
	loggerLevel := strings.ToLower(os.Getenv("LOGGER_LEVEL"))

	validLevels := map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
	}

	if level, exists := validLevels[loggerLevel]; exists {
		return level
	}

	logrus.Warnf("Invalid LOGGER_LEVEL: '%s', defaulting to DEBUG", loggerLevel)

	return logrus.DebugLevel
}

type LoggerMiddleware struct {
	SkipPaths []string
	Level     logrus.Level
}

func (m LoggerMiddleware) Use(r *gin.Engine) {
	middleware := func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		url := c.Request.URL.RequestURI()

		c.Next()

		track := true
		for _, p := range m.SkipPaths {
			if strings.Contains(path, p) {
				track = false

				break
			}
		}

		for _, p := range m.SkipPaths {
			if strings.Contains(url, p) {
				track = false

				break
			}
		}

		if track {
			end := time.Now()
			latency := end.Sub(start)

			msg := "Request"
			if len(c.Errors) > 0 {
				msg = c.Errors.String()
			}

			log := logrus.WithFields(logrus.Fields{
				"module":  "http",
				"method":  c.Request.Method,
				"path":    url,
				"status":  c.Writer.Status(),
				"latency": latency,
				"ip":      c.ClientIP(),
			})
			log.Logger.SetLevel(m.Level)

			if c.Writer.Status() > 299 {
				log.Error(msg)
			} else {
				log.Info(msg)
			}
		}
	}

	r.Use(middleware)
}

func NewLogger(skipPaths []string, level logrus.Level) *LoggerMiddleware {
	return &LoggerMiddleware{
		SkipPaths: skipPaths,
		Level:     level,
	}
}
