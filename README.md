# logrus_conf
logrusの設定を読み込む

## setup

``` bash
go get github.com/sirupsen/logrus
go get github.com/SampleUser0001/logrus_conf
mkdir log
```

## Sample Program

``` go
package main

import (
	logconf "github.com/SampleUser0001/logrus_conf"
	"github.com/sirupsen/logrus"
)

func main() {
	logconf.LogConf()
	logconf.Log.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")
}
```
