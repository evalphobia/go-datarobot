go-datarobot
====

[![Build Status](https://travis-ci.org/evalphobia/go-datarobot.svg?branch=master)](https://travis-ci.org/evalphobia/go-datarobot) [![codecov](https://codecov.io/gh/evalphobia/go-datarobot/branch/master/graph/badge.svg)](https://codecov.io/gh/evalphobia/go-datarobot)
 [![GoDoc](https://godoc.org/github.com/evalphobia/go-datarobot?status.svg)](https://godoc.org/github.com/evalphobia/go-datarobot)


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
    -model <model id> \
    -host <dedicated host> \
    -key <dedicated host key>

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
- `-block`: row data size to send api request (default: 1000)
- `-host` : dedicated hostname; will use shared host if none given (default: https://app.datarobot.com)
- `-key` : dedicated host key (default: )

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
