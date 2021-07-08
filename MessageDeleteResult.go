package go-send-sms-http-rest-ozeki

import "fmt"

type MessageDeleteResult struct {
	folder                  string
	messageIdsRemoveSucceed []string
	messageIdsRemoveFailed  []string
	successCount            int
	failedCount             int
	totalCount              int
}

func NewMessageDeleteResult(folder string) MessageDeleteResult {
	var result MessageDeleteResult = MessageDeleteResult{}
	result.folder = folder
	result.messageIdsRemoveSucceed = []string{}
	result.messageIdsRemoveFailed = []string{}
	return result
}

func (m *MessageDeleteResult) addIdRemoveSucceed(id string) {
	m.messageIdsRemoveSucceed = append(m.messageIdsRemoveSucceed, id)
	m.successCount += 1
	m.totalCount += 1
}

func (m *MessageDeleteResult) addIdRemoveFailed(id string) {
	m.messageIdsRemoveFailed = append(m.messageIdsRemoveFailed, id)
	m.failedCount += 1
	m.totalCount += 1
}

func (m MessageDeleteResult) String() string {
	if m.totalCount == 1 {
		if m.totalCount == m.failedCount {
			return "true"
		} else {
			return "false"
		}
	} else {
		return fmt.Sprintf("Total: %d. Success: %d. Failed: %d.", m.totalCount, m.successCount, m.failedCount)
	}
}
