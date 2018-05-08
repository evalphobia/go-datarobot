go-datarobot
----

[![GoDoc][1]][2] [![License: Apache 2.0][23]][24] [![Release][5]][6] [![Build Status][7]][8] [![Codecov Coverage][11]][12] [![Go Report Card][13]][14] [![Code Climate][19]][20] [![BCH compliance][21]][22]

[1]: https://godoc.org/github.com/evalphobia/go-datarobot?status.svg
[2]: https://godoc.org/github.com/evalphobia/go-datarobot
[3]: https://img.shields.io/badge/License-MIT-blue.svg
[4]: LICENSE.md
[5]: https://img.shields.io/github/release/evalphobia/go-datarobot.svg
[6]: https://github.com/evalphobia/go-datarobot/releases/latest
[7]: https://travis-ci.org/evalphobia/go-datarobot.svg?branch=master
[8]: https://travis-ci.org/evalphobia/go-datarobot
[9]: https://coveralls.io/repos/evalphobia/go-datarobot/badge.svg?branch=master&service=github
[10]: https://coveralls.io/github/evalphobia/go-datarobot?branch=master
[11]: https://codecov.io/github/evalphobia/go-datarobot/coverage.svg?branch=master
[12]: https://codecov.io/github/evalphobia/go-datarobot?branch=master
[13]: https://goreportcard.com/badge/github.com/evalphobia/go-datarobot
[14]: https://goreportcard.com/report/github.com/evalphobia/go-datarobot
[15]: https://img.shields.io/github/downloads/evalphobia/go-datarobot/total.svg?maxAge=1800
[16]: https://github.com/evalphobia/go-datarobot/releases
[17]: https://img.shields.io/github/stars/evalphobia/go-datarobot.svg
[18]: https://github.com/evalphobia/go-datarobot/stargazers
[19]: https://codeclimate.com/github/evalphobia/go-datarobot/badges/gpa.svg
[20]: https://codeclimate.com/github/evalphobia/go-datarobot
[21]: https://bettercodehub.com/edge/badge/evalphobia/go-datarobot?branch=master
[22]: https://bettercodehub.com/
[23]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[24]: LICENSE.md


- `go-datarobot` is golang cli app to get prediction result from [Datarobot](https://www.datarobot.com/) API and add class probabilities into the csv file.

- `go-datarobot/apiclient` is [Datarobot](https://www.datarobot.com/) API client for Golang.

# Quick Usage

## command line

```sh
# install
$ go get github.com/evalphobia/go-datarobot


# check input data
$ cat ./example/input.csv

user_id,flag
1,true
2,false

# execute to predict from input csv data
$ go-datarobot \
    -input ./example/input.csv \
    -output ./example/output.csv \
    -user example@example.com \
    -token <your api token> \
    -project <project id> \
    -model <model id>

======== running setting ==========
input: ./example/input.csv
output: ./example/output.csv
user: example@example.com
project: <project id>
model: <model id>
block size: 1000
====================================
[INFO] send request: part=0
[INFO] send request: part=1
[INFO] finished reading all lines
[INFO] finished request: part=1
[INFO] write line: part=1
[INFO] finished request: part=0
[INFO] write line: part=0
[INFO] finished all requests
end


# show predict result
$ cat ./example/output.csv

user_id,flag,true,false
1,true,0.9929134249687195,0.007086575031280518
2,false,0.0012341737747192383,0.9987658262252808
```

### Options

- `-input`: input csv file path (required)
- `-output`: output csv file path (required)
- `-user`: datarobot account user name (required)
- `-token`: datarobot account api token (required)
- `-project`: datarobot account project id (required)
- `-model`: datarobot account model id (required)
- `-block`: row data size to send api request (default:1000)

## API client

### predict API

```go
import(
	"github.com/evalphobia/go-datarobot/apiclient/config"
	"github.com/evalphobia/go-datarobot/apiclient/predict"
)

func main(){
	user := "example@example.com"
	token := "<your api token>"
	c := config.NewWithToken(user, token)

	row1 := map[string]interface{}{
		"user_id": 1,
		"flag": true,
	}
	row2 := map[string]interface{}{
		"user_id": 2,
		"flag": false,
	}
	items := []map[string]interface{}{
		row1,
		row2,
	}

	resp, err := predict.Predict(c, predict.Param{
		ProjectID: "<project id>",
		ModelID:   "<model id>",
		Data:      items,
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("predictions: %v\n", resp.Predictions)
}
```


# License

Apache License, Version 2.0
