package config

const (
	endpointApp = "https://app.datarobot.com"
)

// Config has auth data for datarobot api
type Config struct {
	HostName string
	Key      string
	User     string
	Token    string
	Password string
}

// NewWithToken returns initialized Config with api token
func NewWithTokenAndHost(user, token, hostName, key string) Config {
	return Config{
		HostName: hostName,
		Key: key,
		User:  user,
		Token: token,
	}
}

func NewWithToken(user, token string) Config {
	return NewWithTokenAndHost(user, token, endpointApp, "")
}
