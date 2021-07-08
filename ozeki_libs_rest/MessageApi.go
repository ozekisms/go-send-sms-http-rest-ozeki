package ozeki_libs_rest

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type MessageApi struct {
	_configuration Configuration
}

func NewMessageApi(configuration Configuration) MessageApi {
	var api MessageApi = MessageApi{configuration}
	return api
}

func (api *MessageApi) createAuthorizationHeader(username string, password string) string {
	var usernamePassword string = username + ":" + password
	var usernamePasswordEncoded string = b64.StdEncoding.EncodeToString([]byte(usernamePassword))
	return "Basic " + usernamePasswordEncoded
}

func (api *MessageApi) createRequestBody(message interface{}) string {
	var jsonObject = map[string][]Message{}
	messagesArray := []Message{}
	switch messages := message.(type) {
	case Message:
		messagesArray = append(messagesArray, messages)
	case []Message:
		for i := 0; i < len(messages); i++ {
			messagesArray = append(messagesArray, messages[i])
		}
	}
	jsonObject["messages"] = messagesArray
	json, err := json.Marshal(jsonObject)
	if err == nil {
		return string(json)
	} else {
		return "{\"messages\": []}"
	}
}

func (m *MessageApi) createRequestBodyManipulate(folder Folder, message interface{}) string {
	obj := make(map[string]interface{})
	messagesArray := []string{}
	switch messages := message.(type) {
	case Message:
		messagesArray = append(messagesArray, messages.ID)
	case []Message:
		for i := 0; i < len(messages); i++ {
			messagesArray = append(messagesArray, messages[i].ID)
		}
	}
	obj["folder"] = folder
	obj["message_ids"] = messagesArray
	json, err := json.Marshal(obj)
	if err == nil {
		return string(json)
	} else {
		return fmt.Sprintf("{\"messages\": [], \"folder\": %s}", folder)
	}
}

func (api *MessageApi) createUriToSendSms(url string) string {
	return fmt.Sprintf("%s?action=sendmsg", strings.Split(url, "?")[0])
}

func (api *MessageApi) createUriToDeleteSms(url string) string {
	return fmt.Sprintf("%s?action=deletemsg", strings.Split(url, "?")[0])
}

func (api *MessageApi) createUriToMarkSms(url string) string {
	return fmt.Sprintf("%s?action=markmsg", strings.Split(url, "?")[0])
}

func (api *MessageApi) createUriToReceiveSms(url string, folder Folder) string {
	return fmt.Sprintf("%s?action=receivemsg&folder=%s", strings.Split(url, "?")[0], folder)
}

func (api *MessageApi) Send(message interface{}) interface{} {
	var authorizationHeader string = api.createAuthorizationHeader(api._configuration.Username, api._configuration.Password)
	var requestBody string = api.createRequestBody(message)
	var result MessageSendResults = api.getResponseSend(api.doRequestPost(api.createUriToSendSms(api._configuration.ApiUrl), authorizationHeader, requestBody))
	if len(result.results) == 1 {
		return result.results[0]
	} else {
		return result
	}
}

func (api *MessageApi) Delete(folder Folder, message interface{}) MessageDeleteResult {
	var authorizationHeader string = api.createAuthorizationHeader(api._configuration.Username, api._configuration.Password)
	var requestBody string = api.createRequestBodyManipulate(folder, message)
	return api.getResponseDelete(api.doRequestPost(api.createUriToDeleteSms(api._configuration.ApiUrl), authorizationHeader, requestBody), message)
}

func (api *MessageApi) Mark(folder Folder, message interface{}) MessageMarkResult {
	var authorizationHeader string = api.createAuthorizationHeader(api._configuration.Username, api._configuration.Password)
	var requestBody string = api.createRequestBodyManipulate(folder, message)
	return api.getResponseMark(api.doRequestPost(api.createUriToMarkSms(api._configuration.ApiUrl), authorizationHeader, requestBody), message)
}

func (api *MessageApi) DownloadIncoming() MessageReceiveResult {
	var authorizationHeader string = api.createAuthorizationHeader(api._configuration.Username, api._configuration.Password)
	return api.getResponseReceive(api.doRequestGet(api.createUriToReceiveSms(api._configuration.ApiUrl, Inbox), authorizationHeader))
}

func (api *MessageApi) doRequestPost(url string, authorizationHeader string, requestBody string) string {
	client := &http.Client{}
	body := strings.NewReader(requestBody)
	request, error := http.NewRequest("POST", url, io.Reader(body))
	request.Header.Add("Authorization", authorizationHeader)
	request.Header.Add("Content-Type", "application/json")
	if error == nil {
		response, err := client.Do(request)
		if err == nil {
			defer response.Body.Close()

			if response.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(response.Body)
				if err != nil {
					println(err)
				}
				body := string(bodyBytes)
				return body
			}
		}
	}
	return ""
}

func (api *MessageApi) doRequestGet(url string, authorizationHeader string) string {
	client := &http.Client{}
	request, error := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", authorizationHeader)
	request.Header.Add("Accept", "application/json")
	if error == nil {
		response, err := client.Do(request)
		if err == nil {
			defer response.Body.Close()

			if response.StatusCode == http.StatusOK {
				bodyBytes, err := ioutil.ReadAll(response.Body)
				if err != nil {
					println(err)
				}
				body := string(bodyBytes)
				return body
			}
		}
	}
	return ""
}

func (api *MessageApi) getResponseSend(response string) MessageSendResults {
	var resp send_response_struct
	json.Unmarshal([]byte(response), &resp)
	var results MessageSendResults = MessageSendResults{}
	results.totalCount = resp.Data.TotalCount
	results.successCount = resp.Data.SuccessCount
	results.failedCount = resp.Data.FailedCount
	for i := 0; i < len(resp.Data.Messages); i++ {
		var result MessageSendResult = MessageSendResult{}
		var msg Message = resp.Data.Messages[i]
		result.message = msg
		if msg.Status == "SUCCESS" {
			result.status = Success
			result.responseMessage = ""
		} else {
			result.status = Success
			result.responseMessage = msg.Status
		}
		results.results = append(results.results, result)
	}
	return results
}

func (m *MessageApi) getResponseDelete(response string, messages interface{}) MessageDeleteResult {
	var resp manipulate_response_struct
	json.Unmarshal([]byte(response), &resp)
	var result MessageDeleteResult = NewMessageDeleteResult(resp.Data.Folder)
	switch messages := messages.(type) {
	case Message:
		if len(resp.Data.MessageIds) == 1 && resp.Data.MessageIds[0] == messages.ID {
			result.addIdRemoveSucceed(messages.ID)
		} else {
			result.addIdRemoveFailed(messages.ID)
		}
	case []Message:
		for i := 0; i < len(messages); i++ {
			var success bool = false
			for j := 0; j < len(resp.Data.MessageIds); j++ {
				if messages[i].ID == resp.Data.MessageIds[j] {
					success = true
				}
			}
			if success {
				result.addIdRemoveSucceed(messages[i].ID)
			} else {
				result.addIdRemoveFailed(messages[i].ID)
			}
		}
	}
	return result
}

func (m *MessageApi) getResponseMark(response string, messages interface{}) MessageMarkResult {
	var resp manipulate_response_struct
	json.Unmarshal([]byte(response), &resp)
	var result MessageMarkResult = NewMessageMarkResult(resp.Data.Folder)
	switch messages := messages.(type) {
	case Message:
		if len(resp.Data.MessageIds) == 1 && resp.Data.MessageIds[0] == messages.ID {
			result.addIdMarkSucceed(messages.ID)
		} else {
			result.addIdMarkFailed(messages.ID)
		}
	case []Message:
		for i := 0; i < len(messages); i++ {
			var success bool = false
			for j := 0; j < len(resp.Data.MessageIds); j++ {
				if messages[i].ID == resp.Data.MessageIds[j] {
					success = true
				}
			}
			if success {
				result.addIdMarkSucceed(messages[i].ID)
			} else {
				result.addIdMarkFailed(messages[i].ID)
			}
		}
	}
	return result
}

func (api *MessageApi) getResponseReceive(response string) MessageReceiveResult {
	var resp receive_response_struct
	json.Unmarshal([]byte(response), &resp)
	var result MessageReceiveResult = MessageReceiveResult{}
	result.folder = resp.Data.Folder
	result.limit = resp.Data.Limit
	result.messages = resp.Data.Data
	return result
}
