package predict

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrediction(t *testing.T) {
	assert := assert.New(t)

	p := &Prediction{}
	body := jsonPredictions
	err := json.Unmarshal([]byte(body), &p)
	assert.NoError(err)
	assert.Equal(1, p.RowID)
	assert.Equal(`"2"`, p.Prediction)
	assert.Equal(0.13461583977122482, p.ClassProbabilities[`"1"`])
	assert.Equal(0.8653841602287752, p.ClassProbabilities[`"2"`])
}

func TestResponse(t *testing.T) {
	assert := assert.New(t)

	r := &Response{}
	body := jsonReponse
	err := json.Unmarshal([]byte(body), &r)
	assert.NoError(err)
	assert.Equal("", r.Status)
	assert.Equal("1234567890abcd", r.ModelID)
	assert.Equal("Binary", r.Task)
	assert.Equal("v1", r.Version)
	assert.Equal(200, r.Code)
	assert.Equal(28449.733018875122, r.ExecutionTime)
	assert.Len(r.Predictions, 3)

	p1 := r.Predictions[1]
	assert.Equal(1, p1.RowID)
	assert.Equal(`"1"`, p1.Prediction)
	assert.Equal(0.766407100127209, p1.ClassProbabilities[`"1"`])
	assert.Equal(0.23359289987279097, p1.ClassProbabilities[`"2"`])
}

var jsonPredictions = `
{
  "row_id": 1,
  "prediction": "\"2\"",
  "class_probabilities": {
    "\"1\"": 0.13461583977122482,
    "\"2\"": 0.8653841602287752
  }
}
`

var jsonReponse = `{"status": "", "model_id": "1234567890abcd", "code": 200, "execution_time": 28449.733018875122, "predictions": [{"row_id": 0, "prediction": "\"2\"", "class_probabilities": {"\"1\"": 0.13461583977122482, "\"2\"": 0.8653841602287752}}, {"row_id": 1, "prediction": "\"1\"", "class_probabilities": {"\"1\"": 0.766407100127209, "\"2\"": 0.23359289987279097}}, {"row_id": 2, "prediction": "\"1\"", "class_probabilities": {"\"1\"": 0.859911414488622, "\"2\"": 0.14008858551137807}}], "task": "Binary", "version": "v1"}`
