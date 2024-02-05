package logconf

import (
	"os"

	"github.com/sirupsen/logrus"
)

// ログインスタンス生成
var Log = logrus.New()

func LogConf() {
	//
	Log.SetReportCaller(true)

	// ファイル出力の場合の例
	file, err := os.OpenFile("./log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}

}
