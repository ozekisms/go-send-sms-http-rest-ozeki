package go_send_sms_http_rest_ozeki

type Folder string

const (
	Inbox   Folder = "inbox"
	Outbox  Folder = "outbox"
	Sent    Folder = "sent"
	NotSent Folder = "notsent"
	Deleted Folder = "deleted"
)
