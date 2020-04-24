// +build unit

package lokigrus

import (
	"math"
	"testing"

	"github.com/ic2hrmk/promtail"
	"github.com/sirupsen/logrus"
)

func Test_matchLogLevels(t *testing.T) {
	type args struct {
		logrusLevel logrus.Level
	}
	tests := []struct {
		name    string
		args    args
		want    promtail.Level
		wantErr bool
	}{
		{
			name: "Panic level",
			args: args{
				logrusLevel: logrus.PanicLevel,
			},
			want: promtail.Panic,
		},
		{
			name: "Fatal level",
			args: args{
				logrusLevel: logrus.FatalLevel,
			},
			want: promtail.Fatal,
		},
		{
			name: "Error level",
			args: args{
				logrusLevel: logrus.ErrorLevel,
			},
			want: promtail.Error,
		},
		{
			name: "Warn level",
			args: args{
				logrusLevel: logrus.WarnLevel,
			},
			want: promtail.Warn,
		},
		{
			name: "Info level",
			args: args{
				logrusLevel: logrus.InfoLevel,
			},
			want: promtail.Info,
		},
		{
			name: "Trace level",
			args: args{
				logrusLevel: logrus.TraceLevel,
			},
			want: promtail.Debug,
		},
		{
			name: "Debug level",
			args: args{
				logrusLevel: logrus.DebugLevel,
			},
			want: promtail.Debug,
		},
		{
			name: "UNMATCHED level",
			args: args{
				logrusLevel: logrus.Level(math.MaxUint32),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := matchLogLevels(tt.args.logrusLevel)
			if (err != nil) != tt.wantErr {
				t.Errorf("matchLogLevels() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("matchLogLevels() got = %v, want %v", got, tt.want)
			}
		})
	}
}
