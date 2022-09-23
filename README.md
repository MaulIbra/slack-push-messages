# Push Notification Message To Slack

Simple REST API For push message to channel on slack

## ![](https://cdn-icons-png.flaticon.com/24/2694/2694997.png) **Stack**

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
![Slack](https://img.shields.io/badge/Slack-4A154B?style=for-the-badge&logo=slack&logoColor=white)

## ![](https://cdn-icons-png.flaticon.com/24/4319/4319207.png) **Framework**
![](https://badgen.net/badge/github/Fiber/blue?icon=github)
![](https://badgen.net/badge/github/godotenv/cyan?icon=github)
![](https://img.shields.io/badge/JWT-000000?style=for-the-badge&logo=JSON%20web%20tokens&logoColor=white)

## ![](https://cdn-icons-png.flaticon.com/24/4471/4471714.png) **Deployment**
![](https://img.shields.io/badge/Heroku-430098?style=for-the-badge&logo=heroku&logoColor=white)


## ![](https://cdn-icons-png.flaticon.com/24/610/610363.png) Preparation On Local
if you want to run code on your local machine, follow instruction that explain below :

Prerequisite : Go was installed in your local machine
* clone this repository
* enter to path repository and run **`go mod tidy`**
* create .env file and put code below\
`PORT=8080`\
`SECRET_KEY`\
`EXPIRED_TIME_TOKEN=1` in minutes
* Running app with command `go run main.go local`

## ![](https://cdn-icons-png.flaticon.com/24/718/718064.png) **Endpoint**
Heroku base url : https://slack-push-message.herokuapp.com
* Get Token\
Token will create using JWT with ip address as a payload data
Path : `{{base_url}}/token`\
Method : `POST`\
Response : 
  ```yaml
  {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjM5Mjc2NjQsImlwIjoiMTI3LjAuMC4xIn0.uH30Jmd3YOVf9W-Mslhs_kRIkbfcRbCUmJlFlS7pr3k"
  }
  ```

* Send Message\
Path : `{{base_url}}/send-message`\
Method : `POST`\
Headers :\
  key : `Authorization`
  value : `{{token}}`\
Content-Type : `application/json`
    ```yaml
    {
        "slack_token": "",
        "slack_channel_id": "",
        "message": ""
    }
    ```

