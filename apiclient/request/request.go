package request

import (
	"gopkg.in/h2non/gentleman.v1"
	"gopkg.in/h2non/gentleman.v1/plugins/auth"
	"gopkg.in/h2non/gentleman.v1/plugins/body"

	"github.com/evalphobia/go-datarobot/apiclient/config"
)

const (
	endpointApp = "https://app.datarobot.com"
)

// Post sends POST request to datarobot api
func Post(c config.Config, path string, param interface{}) (*Response, error) {
	cli := gentleman.New()
	cli.Use(auth.Basic(c.User, c.Token))
	cli.URL(endpointApp)

	req := cli.Request()
	req.Path(path)
	req.Use(body.JSON(param))
	req.Method("POST")

	res, err := req.Send()
	if err != nil {
		return nil, err
	}
	return &Response{res}, nil
}

// Response is wrapper struct of http response
type Response struct {
	*gentleman.Response
}
