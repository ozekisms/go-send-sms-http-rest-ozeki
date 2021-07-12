# Go sms library to send sms with http/rest/json
This Go sms library enables you to **send** and **receive** SMS messages from GO with http requests. The library uses HTTP Post requests and JSON encoded content to send the text messages to the mobile network1. It connects to the HTTP SMS API of [Ozeki SMS gateway](https://ozeki-sms-gateway.com).

# What is Ozeki SMS Gateway?
Ozeki SMS Gateway is a powerful SMS Gateway software you can download and install on your Windows or Linux computer or to your Android mobile phone. It provides an HTTP SMS API, that allows you to connect to it from local or remote programs. The reason why companies use Ozeki SMS Gateway as their first point of access to the mobile network, is because it provides service provider independence. When you use Ozeki, the SMS contact lists and sms data is safe, because Ozeki is installed in their own computer (physical or virtual), and Ozeki provides direct access to the mobile network through wireless connections.

Download: [Ozeki SMS Gateway download page](https://ozeki-sms-gateway.com/p_727-download-sms-gateway.html)

Tutorial: [GO send sms sample and tutorial](https://ozeki-sms-gateway.com/p_874-go-send-multiple-sms-with-the-http-rest-api-code-sample.html)

## How to send sms from GO:

**To send sms from GO**
1. [Download Ozeki SMS Gateway](https://ozeki-sms-gateway.com/p_727-download-sms-gateway.html)
2. [Connect Ozeki SMS Gateway to the mobile network](https://ozeki-sms-gateway.com/p_70-mobile-network.html)
3. [Create an HTTP SMS API user](https://ozeki-sms-gateway.com/p_2102-create-an-http-sms-api-user-account.html)
4. Checkout the Github send SMS from GO repository
5. Open the Github SMS send example in Visual Studio Code
6. Install Go extension for Visual Studio Code
7. Compile the Send SMS console project
8. Check the logs in Ozeki SMS Gateway

# How to use the go_send_sms_http_rest_ozeki package?
 
 The go-send-sms-http-rest-ozeki package is a tool for Ozeki SMS Gateway. With this library you can send, delete, mark and receive SMS messages using the built in api of the SMS Gateway. You can use Visual Studio Code with the Go extension to modify the code.

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
## Manual / API reference
If you want to know more about the code, it is useful to visit the manual we provided. You can find detailed step-by-step tutorial and more example codes there.

Link: [How to send sms from C#](https://ozeki-sms-gateway.com/p_873-go-send-sms-with-the-http-rest-api-code-sample.html)

## How to send sms through your Android mobile phone

If you wish to [send SMS through your Android mobile phone from Go](https://android-sms-gateway.com/), you need to [install Ozeki SMS Gateway on your Android](https://ozeki-sms-gateway.com/p_2847-how-to-install-ozeki-sms-gateway-on-android.html) mobile phone. It is recommended to use an Android mobile phone with a minimum of 4GB RAM and a quad core CPU. Most devices today meet these specs. The advantage of using your mobile, is that it is quick to setup and it often allows you to [send sms free of charge](https://android-sms-gateway.com/p_246-how-to-send-sms-free-of-charge.html).
[Android SMS Gateway](https://android-sms-gateway.com)

## Get started now
Stop wasting time and start sending SMS messages with the help of Ozeki SMS Gateway!
