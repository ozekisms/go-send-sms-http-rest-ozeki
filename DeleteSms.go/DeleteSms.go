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
	msg.ID = "eb22f84e-dfe7-11eb-93ad-74d4355e997d"

	api := ozeki.NewMessageApi(configuration)

	result := api.Delete(ozeki.Inbox, msg)

	fmt.Println(result)
}
