package main

import (
	"flag"
	"fmt"
	"io"
	"sync"

	"github.com/evalphobia/go-datarobot/apiclient/config"
	"github.com/evalphobia/go-datarobot/apiclient/predict"
)

var (
	input     string
	output    string
	user      string
	token     string
	projectID string
	modelID   string
	blockSize int

	lock = sync.Mutex{}
)

func initFlag() {
	flag.StringVar(&input, "input", "", "read file")
	flag.StringVar(&output, "output", "", "write file")
	flag.StringVar(&user, "user", "", "datarobot login user")
	flag.StringVar(&token, "token", "", "datarobot api token")
	flag.StringVar(&projectID, "project", "", "datarobot project_id")
	flag.StringVar(&modelID, "model", "", "datarobot model_id")
	flag.IntVar(&blockSize, "block", 1000, "block data size to send api request")
	flag.Parse()

	fmt.Println("======== running setting ==========")
	fmt.Printf("input: %s\noutput: %s\nuser: %s\nproject: %s\nmodel: %s\nblock size: %d\n",
		input, output,
		user, projectID, modelID,
		blockSize)
	fmt.Println("====================================")
}

func main() {
	fmt.Println("init")
	initFlag()

	csv, err := NewCSVHandler(input, output)
	if err != nil {
		panic(err)
	}
	defer csv.Close()

	// read
	readCh := readLines(csv)

	// send req
	c := config.NewWithToken(user, token)
	maxReq := make(chan bool, 5)
	writeCh := make(chan apiResult)
	reqCount := 0
	readCount := 0
	for i := 0; ; i++ {
		items, ok := <-readCh
		if !ok {
			readCount = i
			fmt.Println("[INFO] finished reading all lines")
			break
		}
		go func(i int, items []map[string]interface{}) {
			maxReq <- true
			fmt.Printf("[INFO] send request: part=%d\n", i)
			resp, err := predict.Predict(c, predict.Param{
				ProjectID: projectID,
				ModelID:   modelID,
				Data:      items,
			})
			<-maxReq
			reqCount++

			if err != nil {
				fmt.Printf("[ERROR] request: part=%d, err=%s\n", i, err.Error())
				panic("error on request")
			}
			fmt.Printf("[INFO] finished request: part=%d\n", i)

			writeCh <- apiResult{
				idx:         i,
				items:       items,
				predictions: resp.Predictions,
				err:         err,
			}
			if reqCount == readCount {
				fmt.Println("[INFO] finished all requests")
				close(writeCh)
			}
		}(i, items)
	}

	err = writeLines(csv, writeCh)
	if err != nil {
		fmt.Printf("[ERROR] writeLine end: %s", err.Error())
		return
	}

	fmt.Println("end")
}

func readLines(c *CSVHandler) chan []map[string]interface{} {
	ch := make(chan []map[string]interface{})
	go func() {
		for {
			items, err := c.ReadMapItems(blockSize)
			switch {
			case err == io.EOF:
				if len(items) > 0 {
					ch <- items
				}
				close(ch)
				return
			case err != nil:
				panic(err)
			default:
				ch <- items
			}
		}
	}()
	return ch
}

func writeLines(csv *CSVHandler, ch chan apiResult) error {
	for {
		data, ok := <-ch
		if !ok {
			return nil
		}
		lock.Lock()
		fmt.Printf("[INFO] write line: part=%d\n", data.idx)

		items := data.items
		for _, p := range data.predictions {
			if p.RowID >= len(items) {
				fmt.Printf("[ERROR] index out of range for result and items: part=%d, rowID=%d\n", data.idx, p.RowID)
				continue
			}
			item := items[p.RowID]

			// add class probabilities
			newKeys := make([]string, 0, len(p.ClassProbabilities))
			for key, val := range p.ClassProbabilities {
				item[key] = val
				newKeys = append(newKeys, key)
			}
			// write header once
			if len(csv.wHeader) == 0 {
				csv.wHeader = append(csv.rHeader, newKeys...)
				csv.Write(csv.wHeader)
			}

			// write line
			line := make([]string, len(csv.wHeader))
			for i, key := range csv.wHeader {
				line[i] = fmt.Sprint(item[key])
			}
			csv.Write(line)
		}
		lock.Unlock()
	}
}

type apiResult struct {
	idx         int
	items       []map[string]interface{}
	predictions []predict.Prediction
	err         error
}
