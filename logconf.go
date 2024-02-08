package logconf

import (
	"encoding/json"
	"os"

	"github.com/sirupsen/logrus"
)

// ログインスタンス生成
var Log = logrus.New()

func LogConf() {
	Log.SetReportCaller(true)
	setLogPath("./log/app.log")
}

type Config struct {
	ReportCaller bool   `json:"reportCaller"`
	LogFile      string `json:"logFile"`
}

func ConfLoad(confpath string) {
	// Read the JSON file
	data, err := os.ReadFile(confpath)
	if err != nil {
		Log.Warn("Failed to read the config file.")
		LogConf()
		return // ここで終了
	}

	var c Config
	err = json.Unmarshal(data, &c)
	if err != nil {
		Log.Warn("Failed to unmarshal the config file.")
		LogConf()
		return // ここで終了
	}

	Log.SetReportCaller(c.ReportCaller)
	setLogPath(c.LogFile)
}

func setLogPath(logpath string) {
	file, err := os.OpenFile(logpath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to log to file, using default stderr")
	}

}
