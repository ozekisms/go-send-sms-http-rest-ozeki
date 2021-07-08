package go_send_sms_http_rest_ozeki

import "fmt"

type MessageSendResults struct {
	TotalCount   int
	SuccessCount int
	FailedCount  int
	Results      []MessageSendResult
}

func (m MessageSendResults) String() string {
	return fmt.Sprintf("Total: %d. Success: %d. Failed: %d.", m.TotalCount, m.SuccessCount, m.FailedCount)
}
