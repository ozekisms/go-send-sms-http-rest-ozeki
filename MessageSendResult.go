package go-send-sms-http-rest-ozeki

import "fmt"

type send_response_struct struct {
	HttpCode    int              "json:\"http_code\""
	ReponseCode string           "json:\"reponse_code\""
	ResponseMsg string           "json:\"response_msg\""
	Data        send_data_struct "json:\"data\""
}

type send_data_struct struct {
	TotalCount   int       "json:\"total_count\""
	SuccessCount int       "json:\"success_count\""
	FailedCount  int       "json:\"failed_count\""
	Messages     []Message "json:\"messages\""
}

type MessageSendResult struct {
	message         Message
	status          DeliveryStatus
	responseMessage string
}

func (m MessageSendResult) String() string {
	return fmt.Sprintf("%s, %s", m.status, m.message)
}
