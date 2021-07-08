package ozeki_libs_rest

import "fmt"

type MessageMarkResult struct {
	folder                string
	messageIdsMarkSucceed []string
	messageIdsMarkFailed  []string
	successCount          int
	failedCount           int
	totalCount            int
}

func NewMessageMarkResult(folder string) MessageMarkResult {
	var result MessageMarkResult = MessageMarkResult{}
	result.folder = folder
	result.messageIdsMarkSucceed = []string{}
	result.messageIdsMarkFailed = []string{}
	return result
}

func (m *MessageMarkResult) addIdMarkSucceed(id string) {
	m.messageIdsMarkSucceed = append(m.messageIdsMarkSucceed, id)
	m.successCount += 1
	m.totalCount += 1
}

func (m *MessageMarkResult) addIdMarkFailed(id string) {
	m.messageIdsMarkFailed = append(m.messageIdsMarkFailed, id)
	m.failedCount += 1
	m.totalCount += 1
}

func (m MessageMarkResult) String() string {
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
