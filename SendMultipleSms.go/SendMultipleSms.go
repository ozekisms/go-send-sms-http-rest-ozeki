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

	msg1 := ozeki.NewMessage()
	msg1.ToAddress = "+36201111111"
	msg1.Text = "Hello world 1"

	msg2 := ozeki.NewMessage()
	msg2.ToAddress = "+36202222222"
	msg2.Text = "Hello world 2"

	msg3 := ozeki.NewMessage()
	msg3.ToAddress = "+36203333333"
	msg3.Text = "Hello world 3"

	api := ozeki.NewMessageApi(configuration)

	result := api.Send([]ozeki.Message{msg1, msg2, msg3})

	fmt.Println(result)
}
