package request

import (
	"gopkg.in/h2non/gentleman.v1"
	"gopkg.in/h2non/gentleman.v1/plugins/auth"
	"gopkg.in/h2non/gentleman.v1/plugins/body"

	"github.com/evalphobia/go-datarobot/apiclient/config"
)
// Post sends POST request to datarobot api
func Post(c config.Config, path string, param interface{}) (*Response, error) {
	cli := gentleman.New()
	cli.Use(auth.Basic(c.User, c.Token))
	cli.URL(c.HostName)

	req := cli.Request()
	req.Path(path)
	req.Use(body.JSON(param))
	req.Method("POST")
	if c.Key != "" {
		req.AddHeader("datarobot-key", c.Key)
	}

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
