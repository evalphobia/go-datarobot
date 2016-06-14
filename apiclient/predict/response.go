package predict

import "github.com/evalphobia/go-datarobot/apiclient/request"

// Response is api response from datarobot
type Response struct {
	request.BaseAPIResponse

	ModelID       string       `json:"model_id"`
	ExecutionTime float64      `json:"execution_time"`
	Task          string       `json:"task"`
	Predictions   []Prediction `json:"predictions"`
}

// Prediction is struct for predict result
type Prediction struct {
	RowID              int                `json:"row_id"`
	Prediction         string             `json:"prediction"`
	ClassProbabilities map[string]float64 `json:"class_probabilities,omitempty"`
}

// GetProbability returns the class probabillity from key
func (p Prediction) GetProbability(key string) (float64, bool) {
	if v, ok := p.ClassProbabilities[key]; ok {
		return v, true
	}
	if v, ok := p.ClassProbabilities[`"`+key+`"`]; ok {
		return v, true
	}
	return 0, false
}

// MustGetProbability returns the class probabillity from key and return 0 if the key does not exist.
func (p Prediction) MustGetProbability(key string) float64 {
	if v, ok := p.ClassProbabilities[key]; ok {
		return v
	}
	if v, ok := p.ClassProbabilities[`"`+key+`"`]; ok {
		return v
	}
	return 0
}
