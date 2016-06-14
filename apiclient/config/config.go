package config

// Config has auth data for datarobot api
type Config struct {
	User     string
	Token    string
	Password string
}

// NewWithToken returns initialized Config with api token
func NewWithToken(user, token string) Config {
	return Config{
		User:  user,
		Token: token,
	}
}
