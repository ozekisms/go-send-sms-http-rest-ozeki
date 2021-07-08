# How to use the go_send_sms_http_rest_ozeki package
 
 The go-send-sms-http-rest-ozeki package is a tool for Ozeki SMS Gateway. With this library you can send, delete, mark and receive SMS messages using the built in api of the SMS Gateway.

To use the example code called SendSms.go you have to set up an http_user in your SMS Gateway.

It is also important to mention, that you have to run the code on a computer where the SMS Gateway gateway is running.

## Installation

To install the go_send_sms_http_rest_ozeki package, you have to execute the following command:

    $ go get github.com/ozekisms/go_send_sms_http_rest_ozeki

## How to send a simple SMS message

To import the package into your Go application:

```go
    import ozeki "github.com/ozekisms/go_send_sms_http_rest_ozeki"
```

 To use the ozeki_libs_rest you have to create a Configuration object.
```go
    configuration := ozeki.NewConfiguration(
        "username",
        "password",
        "http://127.0.0.1:9509/api",
    )
```
To initialize a Message object we have to use the following code:

```go
    msg := ozeki.NewMessage()
    msg.ToAddress = "+36201111111"
    msg.Text = "Hello world!"
```
To send your message  we should also create a MessageApi object.
The MessageApi constructor takes only one parameter which is a configuration object.

```go
    api := ozeki.NewMessageApi(configuration)
    
    result := api.Send(msg) #We save our result in a varriable
```
