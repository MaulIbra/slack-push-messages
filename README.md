# Push Notification Message To Slack

Simple REST API For push message to channel on slack

## ![](https://cdn-icons-png.flaticon.com/24/2694/2694997.png) **Stack**

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Slack](https://img.shields.io/badge/Slack-4A154B?style=for-the-badge&logo=slack&logoColor=white)

## ![](https://cdn-icons-png.flaticon.com/24/4319/4319207.png) **Framework**
![](https://badgen.net/badge/github/Fiber/blue?icon=github&link=http://youtube.com&logo=Github)
![](https://badgen.net/badge/github/godotenv/cyan?icon=github)

## ![](https://cdn-icons-png.flaticon.com/24/610/610363.png) Preparation On Local
if you want to running code on your local machine, follow instruction that explain below :

Prerequisite : Go was installed in your local machine
* clone this repository
* enter to path repository and run **`go mod tidy`**
* create .env file and put code below\
`PORT=`
* Running app with command `go run main.go local`

## ![](https://cdn-icons-png.flaticon.com/24/718/718064.png) **Endpoint**
* Send Message\
Path : `{{base_url}}/send-message`\
Method : `POST`\
Content-Type : `application/json`
    ```yaml
    {
        "slack_token": "",
        "slack_channel_id": "",
        "message": ""
    }
    ```

