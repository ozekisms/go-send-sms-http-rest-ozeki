package go-send-sms-http-rest-ozeki

type Folder string

const (
	Inbox   Folder = "inbox"
	Outbox  Folder = "outbox"
	Sent    Folder = "sent"
	NotSent Folder = "notsent"
	Deleted Folder = "deleted"
)
