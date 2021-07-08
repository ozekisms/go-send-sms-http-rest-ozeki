package go_send_sms_http_rest_ozeki

type Configuration struct {
	Username string
	Password string
	ApiUrl   string
}

func NewConfiguration(Username string, Password string, ApiUrl string) Configuration {
	var configuration Configuration = Configuration{Username, Password, ApiUrl}
	return configuration
}
