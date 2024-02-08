package logconf

import (
	"os"
	"strconv"
	"testing"
)

func Test_ConfLoad_Correct(t *testing.T) {
	ConfLoad("testdata/test_logconf.json")
	if Log.Out == nil {
		t.Errorf("Log.Out is nil")
	} else if Log.Out.(*os.File).Name() != "./log/test.log" {
		t.Errorf("Log.Out is not ./log/test.log")
	}

	// ファイルの設定値はfalse。
	if Log.ReportCaller == true {
		t.Errorf("Log.ReportCaller is not false")
	}
}

/**
 * ファイルが存在しない場合はデフォルト値が設定される。
 */
func Test_ConfLoad_FileNotFound(t *testing.T) {
	defaultConfig := Config{
		ReportCaller: true,
		LogFile:      "./log/app.log",
	}

	// ファイルが存在しない場合はデフォルト値が設定される。
	ConfLoad("testdata/notfound.json")

	if Log.Out == nil {
		t.Errorf("Log.Out is nil")
	} else if Log.Out.(*os.File).Name() != defaultConfig.LogFile {
		t.Errorf("Log.Out is not " + defaultConfig.LogFile)
	}

	// デフォルト値はtrue。
	if Log.ReportCaller != defaultConfig.ReportCaller {
		t.Errorf("Log.ReportCaller is not " + strconv.FormatBool(defaultConfig.ReportCaller))
	}
}

/**
 * 項目が存在しない場合はデフォルト値が設定される。
 */
func Test_ConfLoad_ItemNotFound(t *testing.T) {
	defaultConfig := Config{
		ReportCaller: true,
		LogFile:      "./log/app.log",
	}

	ConfLoad("testdata/empty.json")
	if Log.Out == nil {
		t.Errorf("Log.Out is nil")
	} else if Log.Out.(*os.File).Name() != defaultConfig.LogFile {
		t.Errorf("Log.Out is not " + defaultConfig.LogFile)
	}

	// デフォルト値はtrue。
	if Log.ReportCaller != defaultConfig.ReportCaller {
		t.Errorf("Log.ReportCaller is not " + strconv.FormatBool(defaultConfig.ReportCaller))
	}

}
