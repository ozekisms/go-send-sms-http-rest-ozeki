package go_send_sms_http_rest_ozeki

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID                        string              "json:\"message_id\""
	FromConnection            string              "json:\"from_connection\""
	FromAddress               string              "json:\"from_address\""
	FromStation               string              "json:\"from_station\""
	ToConnection              string              "json:\"to_connection\""
	ToAddress                 string              "json:\"to_address\""
	ToStation                 string              "json:\"to_station\""
	Text                      string              "json:\"text\""
	CreateDate                string              "json:\"create_date\""
	ValidUntil                string              "json:\"valid_until\""
	TimeToSend                string              "json:\"time_to_send\""
	IsSubmitReportRequested   bool                "json:\"submit_report_requested\""
	IsViewReportRequested     bool                "json:\"view_report_requested\""
	IsDeliveryReportRequested bool                "json:\"delivery_report_requested\""
	Tags                      []map[string]string "json:\"tags\""
	Status                    string              "json:\"status\""
}

func NewMessage() Message {
	var message Message = Message{}
	var uuid, error = uuid.NewUUID()
	if error == nil {
		message.ID = uuid.String()
	} else {
		message.ID = ""
	}
	message.CreateDate = time.Now().Format("2006-01-02T15:04:05")
	message.ValidUntil = time.Now().Add(time.Hour * 24 * 7 * time.Duration(1)).Format("2006-01-02T15:04:05")
	message.TimeToSend = time.Now().Format("2006-01-02T15:04:05")
	message.IsSubmitReportRequested = true
	message.IsDeliveryReportRequested = true
	message.IsViewReportRequested = true
	message.Tags = []map[string]string{}
	return message
}

func (m *Message) addTag(name string, value string) {
	pair := make(map[string]string)
	pair["name"] = name
	pair["value"] = value
	m.Tags = append(m.Tags, pair)
}

func (m Message) String() string {
	return fmt.Sprintf("%s->%s '%s'", m.FromAddress, m.ToAddress, m.Text)
}
