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

	msg := ozeki.NewMessage()
	msg.ToAddress = "+36201111111"
	msg.Text = "Hello world!"

	api := ozeki.NewMessageApi(configuration)

	result := api.Send(msg)

	fmt.Println(result)
}
