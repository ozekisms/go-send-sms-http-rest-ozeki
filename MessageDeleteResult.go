package go_send_sms_http_rest_ozeki

import "fmt"

type MessageDeleteResult struct {
	Folder                  string
	MessageIdsRemoveSucceed []string
	MessageIdsRemoveFailed  []string
	SuccessCount            int
	FailedCount             int
	TotalCount              int
}

func NewMessageDeleteResult(folder string) MessageDeleteResult {
	var result MessageDeleteResult = MessageDeleteResult{}
	result.Folder = folder
	result.MessageIdsRemoveSucceed = []string{}
	result.MessageIdsRemoveFailed = []string{}
	return result
}

func (m *MessageDeleteResult) addIdRemoveSucceed(id string) {
	m.MessageIdsRemoveSucceed = append(m.MessageIdsRemoveSucceed, id)
	m.SuccessCount += 1
	m.TotalCount += 1
}

func (m *MessageDeleteResult) addIdRemoveFailed(id string) {
	m.MessageIdsRemoveFailed = append(m.MessageIdsRemoveFailed, id)
	m.FailedCount += 1
	m.TotalCount += 1
}

func (m MessageDeleteResult) String() string {
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
