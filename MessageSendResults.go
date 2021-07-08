package ozeki_libs_rest

import "fmt"

type MessageSendResults struct {
	totalCount   int
	successCount int
	failedCount  int
	results      []MessageSendResult
}

func (m MessageSendResults) String() string {
	return fmt.Sprintf("Total: %d. Success: %d. Failed: %d.", m.totalCount, m.successCount, m.failedCount)
}
