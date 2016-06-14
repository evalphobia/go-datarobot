package predict

import (
	"fmt"

	"github.com/evalphobia/go-datarobot/apiclient/config"
	"github.com/evalphobia/go-datarobot/apiclient/request"
)

const (
	endpointPredict = "/api/v1/%s/%s/predict"
	endpointTopK    = "/api/v1/%s/%s/predict_user_topk"
)

// Predict calls predict api and return the response
func Predict(conf config.Config, param Param) (*Response, error) {
	if ok, col := param.IsValid(); !ok {
		return nil, fmt.Errorf("invalid parameter: %s", col)
	}

	path := fmt.Sprintf(endpointPredict, param.ProjectID, param.ModelID)
	resp, err := request.Post(conf, path, param.Data)
	switch {
	case err != nil:
		return nil, err
	case resp.Error != nil:
		return nil, resp.Error
	case !resp.Ok:
		return nil, fmt.Errorf("status: %d, error: %s", resp.StatusCode, resp.String())
	}

	result := &Response{}
	err = resp.JSON(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
