package go_send_sms_http_rest_ozeki

type manipulate_response_struct struct {
	HttpCode    int                    "json:\"http_code\""
	ReponseCode string                 "json:\"reponse_code\""
	ResponseMsg string                 "json:\"response_msg\""
	Data        manipulate_data_struct "json:\"data\""
}

type manipulate_data_struct struct {
	Folder     string   "json:\"folder\""
	MessageIds []string "json:\"message_ids\""
}
