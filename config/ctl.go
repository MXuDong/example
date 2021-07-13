package config

import (
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"os"
)

type innerValue struct {
	Config    config                `json:"config"`
	ClientSet *kubernetes.Clientset `json:"client_set"`
	Logrus    logrus.Logger         `json:"logrus"`
}

var Ctl *innerValue = nil

func InitInnerValue() *innerValue {
	i := innerValue{}
	i.Config = Config
	// do init
	Ctl = &i

	i.Logrus = logrus.Logger{
		Out:          os.Stdout,
		Formatter:    &logrus.TextFormatter{},
		Level:        getLogLevel(i.Config.ProgramLogConfig.LogLevel),
		ReportCaller: false,
	}

	// init kubernetes from config
	InitKubernetesConfig()

	return &i
}

func getLogLevel(level string) logrus.Level {
	switch level {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	case "trace":
		return logrus.TraceLevel
	default:
		return logrus.InfoLevel
	}
}
