# chat

## Overview

This is a Chat application using WebSocket. ( WebSocket is a communication standard for computer networks. It is a technical standard for realizing bidirectional communication in Web applications. )

![chat_demo](https://user-images.githubusercontent.com/38056766/69639993-f3d73300-10a0-11ea-963d-d9981249fe35.gif)

## Installation

### Requirements

- Golang ver1.12.6

### Setting

First of all, clone this repository.
Then excute `go get` the packages you need. Also create a file to write environment variables with the name `.env` and describe the following.

```
GO111MODULE=on
SECURITY_KEY=
FACEBOOK_CLIENT=
FACEBOOK_SECRET=
FACEBOOK_CALL_BACK_URL=
GITHUB_CLIENT=
GITHUB_SECRET=
GITHUB_CALL_BACK_URL=
GOOGLE_CLIENT=
GOOGLE_SECRET=
GOOGLE_CALL_BACK_URL=
``` 

## Usage

Build and Start web server. Let's open the URL specified by the host in your browser. 

```
❯ go build -o chat

❯ ./chat -host=":8080"
```
