package libs

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

type Logger struct{}

// Print - Log Formatter
func (*Logger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		logrus.WithFields(
			logrus.Fields{
				"type":     "sql",
				"src":      v[1],
				"values":   v[4],
				"duration": v[2],
			},
		).Info(v[3])
	case "log":
		logrus.WithFields(logrus.Fields{"module": "gorm", "type": "log"}).Print(v[2])
	}
}
