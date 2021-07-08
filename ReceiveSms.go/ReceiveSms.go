package main

import (
	"fmt"

	ozeki "github.com/ozekisms/go_send_sms_http_rest_ozeki"
)

func main() {
	configuration := ozeki.NewConfiguration(
		"http_user",
		"qwe123",
		"http://127.0.0.1:9509/api",
	)

	api := ozeki.NewMessageApi(configuration)

	result := api.DownloadIncoming()

	fmt.Println(result)

	for i := 0; i < len(result.Messages); i++ {
		fmt.Println(result.Messages[i])
	}
}
