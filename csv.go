package main

import (
	"encoding/csv"
	"os"
)

// CSVHandler handles input and output csv files
type CSVHandler struct {
	*csv.Reader
	fpRead  *os.File
	rHeader []string
	lineNo  int

	*csv.Writer
	wHeader []string
	fpWrite *os.File
}

// Close closes file pointer of input and output csv files
func (c *CSVHandler) Close() {
	c.fpRead.Close()
	c.fpWrite.Close()
}

// NewCSVHandler returns initialized *CSVHandler
func NewCSVHandler(in, out string) (*CSVHandler, error) {
	r, err := os.Open(in)
	if err != nil {
		return nil, err
	}

	w, err := os.Create(out)
	if err != nil {
		return nil, err
	}

	c := &CSVHandler{
		fpRead:  r,
		fpWrite: w,
		Reader:  csv.NewReader(r),
		Writer:  csv.NewWriter(w),
	}

	// set header
	c.rHeader, err = c.Read()
	if err != nil {
		return nil, err
	}

	return c, nil
}

// GetPosition returns position(read line number)
func (c *CSVHandler) GetPosition() int {
	return c.lineNo
}

// Read returns []string and count up current position
func (c *CSVHandler) Read() ([]string, error) {
	line, err := c.Reader.Read()
	if err != nil {
		return nil, err
	}

	c.lineNo++
	return line, nil
}

// ReadMapItems reads lines from input csv file nad create map item, which has key=<header column name> val=<the value of the line>
func (c *CSVHandler) ReadMapItems(size int) ([]map[string]interface{}, error) {
	items := make([]map[string]interface{}, 0, size)
	header := c.rHeader
	for i := 0; i < size; i++ {
		line, err := c.Read()
		if err != nil {
			return items, err
		}

		item := make(map[string]interface{})
		for j, key := range header {
			item[key] = line[j]
		}
		items = append(items, item)
	}
	return items, nil
}

// Write writes a line into file
func (c *CSVHandler) Write(line []string) error {
	err := c.Writer.Write(line)
	if err != nil {
		return err
	}
	c.Flush()
	return nil
}
