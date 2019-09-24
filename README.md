# logrus-graylog2-hook
Graylog2 hook for logrus

=====


[![Build Status](https://travis-ci.org/yumimobi/logrus-graylog2-hook.svg?branch=master)](https://travis-ci.org/yumimobi/logrus-graylog2-hook)
[![Go Report Card](http://goreportcard.com/badge/yumimobi/logrus-graylog2-hook)](http://goreportcard.com/report/yumimobi/logrus-graylog2-hook)
[![codecov](https://codecov.io/gh/yumimobi/logrus-graylog2-hook/branch/master/graph/badge.svg)](https://codecov.io/gh/yumimobi/logrus-graylog2-hook)


## Installation

```bash
$ go get github.com/yumimobi/logrus-graylog2-hook
```

## Usage

```go
package main

import (
	"errors"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/yumimobi/logrus-graylog2-hook"
)

func main() {
	hook := graylog.New("127.0.0.1:12201", "ad_server", nil, logrus.DebugLevel)
	logrus.AddHook(hook)
	logrus.SetOutput(ioutil.Discard)

	// 根据设定过滤规则进行信息发送
	// 1. 设置 facility 为 ad_server 的匹配规则
	// logrus.WithError(errors.New("this is an error")).Info("get an error")
}
```

```go
package main

import (
	"errors"
	"io/ioutil"

	"github.com/Sirupsen/logrus"
	"github.com/yumimobi/logrus-graylog2-hook"
)

func main() {
	hook := graylog.New("127.0.0.1:12201", "", nil, logrus.DebugLevel)
	logrus.AddHook(hook)
	logrus.SetOutput(ioutil.Discard)

	// 根据设定过滤规则进行信息发送
	// 2. 设置 message(short_message) 为 debug_data 的匹配规则
	logrus.WithField("user":"张三").Info("debug_data")
	或者
	logrus.WithField("user":"张三").WithField("message":"debug_data").Info("user info")
}
```