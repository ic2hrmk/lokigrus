package lokigrus

import (
	"fmt"

	"github.com/ic2hrmk/promtail"
	"github.com/sirupsen/logrus"
)

type PromtailHook struct {
	logrus.Hook
	client promtail.Client
}

//
// Initializes a Promtail hook for Logrus logger.
//	- lokiAddress - address of Grafana Loki server to push logs to (e.g. loki:3100)
//	- labels - is kinda tags for grepping in Loki's and Grafana's queries
func NewPromtailHook(lokiURL string, labels map[string]string) (*PromtailHook, error) {
	var (
		hook = &PromtailHook{}
		err  error
	)

	hook.client, err = promtail.NewJSONv1Client(lokiURL, labels)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize Promtail client: %s", err.Error())
	}

	return hook, err
}

func (rcv *PromtailHook) Fire(entry *logrus.Entry) error {
	if entry == nil {
		return fmt.Errorf("log entry is nil")
	}

	line, err := entry.String()
	if err != nil {
		return fmt.Errorf("unable to read log entry: %s", err)
	}

	switch entry.Level {
	case logrus.PanicLevel:
		rcv.client.Panicf(line)
	case logrus.FatalLevel:
		rcv.client.Fatalf(line)
	case logrus.ErrorLevel:
		rcv.client.Errorf(line)
	case logrus.WarnLevel:
		rcv.client.Warnf(line)
	case logrus.InfoLevel:
		rcv.client.Infof(line)
	case logrus.DebugLevel, logrus.TraceLevel:
		rcv.client.Debugf(line)
	default:
		return fmt.Errorf("entry has unmatched logrus log level [level=%s]",
			entry.Level.String())
	}

	return nil
}

func (rcv *PromtailHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (rcv *PromtailHook) LokiHealthCheck() error {
	_, err := rcv.client.Ping()
	if err != nil {
		return err
	}

	return nil
}

func matchLogLevels(logrusLevel logrus.Level) (promtail.Level, error) {
	var (
		matchedPromtailLevel promtail.Level
		err                  error
	)

	switch logrusLevel {
	case logrus.PanicLevel:
		matchedPromtailLevel = promtail.Panic
	case logrus.FatalLevel:
		matchedPromtailLevel = promtail.Fatal
	case logrus.ErrorLevel:
		matchedPromtailLevel = promtail.Error
	case logrus.WarnLevel:
		matchedPromtailLevel = promtail.Warn
	case logrus.InfoLevel:
		matchedPromtailLevel = promtail.Info
	case logrus.DebugLevel, logrus.TraceLevel:
		matchedPromtailLevel = promtail.Debug
	default:
		err = fmt.Errorf("entry has unmatched logrus log level [level=%s]",
			logrusLevel.String())
	}

	if err != nil {
		return matchedPromtailLevel, err
	}

	return matchedPromtailLevel, nil
}

// Compile time validation
var _ logrus.Hook = (*PromtailHook)(nil)
