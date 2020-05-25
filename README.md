# Messenger backend with rsa-1024 and AES-128

This repository contains a full backend to build a messenger with option to send messages and files with secure encryption at the same time
it means your partner can download each chunk you have sent. The security strategy is almost same as <a href="https://scontent.whatsapp.net/v/t61.22868-34/68135620_760356657751682_6212997528851833559_n.pdf/WhatsApp-Security-Whitepaper.pdf?_nc_sid=41cc27&_nc_oc=AQkJcUSq9G6NGOZ0ZuRULw_icW7l1Fq-HLh2i2KELDnTOnkpgfZCeAUpQU3pa7mkSLQ&_nc_ht=scontent.whatsapp.net&oh=1a7c6a379166027b672825a7f84c6ec9&oe=5ECCDC93">Whatsapp</a>.

## Getting Started

You have two options to run the server, I myself highly recommend using its docker in dockerhub and docker-compose which is uploaded.
If you want to build this with customized options you should change two pivotal files: main.go and redis.go they both have settings for
mysql and redis connection.

### Prerequisites

Encryption for messages is based on AES-128 with HEX encoding.
Install all requirments by

```
go mod download
```

### Run

As its clear in main.go by serving main.go the server starts at port 4300
### Requests format
There is only one request to the server and the format structure is presented at request.go.
Here is a simple example for handshaking.
```
{
  "info": {
    "token_requested": true,
    "src_username": "aicam",
    "dest_username": "sijal",
    "public_key": "-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCi8sMm2CUQrxju5TVdF7u6iAxxa4sdIkBPARR0TaZTZUcmp+ySmCGXVM68h02M2nui/vCHmxLxbPH34cg3Ul8LoS8MNZIqDk2qekbaRA1s/ehL3kU4chw8050x15XCp9k9y4Q2LrZr/pkoXY23tk7PYT0LOcjKCOe7z/cf1SMcTwIDAQAB\n-----END PUBLIC KEY-----"
  },
  "message": {
    "text": "",
    "offset": 2,
    "hmac": ""
  },
  "file": {
    "data": "",
    "chunk": 0,
    "finished": false,
    "file_name": "",
    "hmac": ""
  }
}

```
