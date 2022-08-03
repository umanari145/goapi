package libs

import (
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct{}

// Print - Log Formatter
func (*Logger) Print(v ...interface{}) {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	//logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "../logs/app.log", // ファイル名
		MaxSize:    500,               // ローテーションするファイルサイズ(megabytes)
		MaxBackups: 3,                 // 保持する古いログの最大ファイル数
		MaxAge:     365,               // 古いログを保持する日数
		LocalTime:  true,              // バックアップファイルの時刻フォーマットをサーバローカル時間指定
		Compress:   true,              // ローテーションされたファイルのgzip圧縮
	})

	switch v[0] {
	case "sql":
		logrus.WithFields(
			logrus.Fields{
				"time":     time.Now().In(jst),
				"type":     "sql",
				"src":      v[1],
				"values":   v[4],
				"duration": v[2],
			},
		).Info(v[3])
	case "log":
		//fmt.Print(v[1])
		logrus.WithFields(logrus.Fields{"time": time.Now().In(jst), "type": "log"}).Print(v[2])
	}
}
