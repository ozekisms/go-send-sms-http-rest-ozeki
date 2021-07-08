package go_send_sms_http_rest_ozeki

import "fmt"

type MessageMarkResult struct {
	Folder                string
	MessageIdsMarkSucceed []string
	MessageIdsMarkFailed  []string
	SuccessCount          int
	FailedCount           int
	TotalCount            int
}

func NewMessageMarkResult(folder string) MessageMarkResult {
	var result MessageMarkResult = MessageMarkResult{}
	result.Folder = folder
	result.MessageIdsMarkSucceed = []string{}
	result.MessageIdsMarkFailed = []string{}
	return result
}

func (m *MessageMarkResult) addIdMarkSucceed(id string) {
	m.MessageIdsMarkSucceed = append(m.MessageIdsMarkSucceed, id)
	m.SuccessCount += 1
	m.TotalCount += 1
}

func (m *MessageMarkResult) addIdMarkFailed(id string) {
	m.MessageIdsMarkFailed = append(m.MessageIdsMarkFailed, id)
	m.FailedCount += 1
	m.TotalCount += 1
}

func (m MessageMarkResult) String() string {
	if m.TotalCount == 1 {
		if m.TotalCount == m.SuccessCount {
			return "true"
		} else {
			return "false"
		}
	} else {
		return fmt.Sprintf("Total: %d. Success: %d. Failed: %d.", m.TotalCount, m.SuccessCount, m.FailedCount)
	}
}
