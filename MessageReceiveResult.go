package go_send_sms_http_rest_ozeki

import "fmt"

type receive_response_struct struct {
	HttpCode     int                 "json:\"http_code\""
	ResponseCode int                 "json:\"response_code\""
	ResponseMsg  string              "json:\"response_msg\""
	Data         receive_data_struct "json:\"data\""
}

type receive_data_struct struct {
	Folder string    "json:\"folder\""
	Limit  string    "json:\"limit\""
	Data   []Message "json:\"data\""
}

type MessageReceiveResult struct {
	Folder   string
	Limit    string
	Messages []Message
}

func (m MessageReceiveResult) String() string {
	return fmt.Sprintf("Message count: %d", len(m.Messages))
}
