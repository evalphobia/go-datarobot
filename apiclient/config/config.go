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
func NewWithToken(user, token, hostName, key string) Config {
	       h := hostName
	       k := key
	       if h == "" {
		       h = endpointApp
	       }

	return Config{
		HostName: h,
		Key: k,
		User:  user,
		Token: token,
	}
}
