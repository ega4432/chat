# chat

## Overview

This is a Chat application using WebSocket. ( WebSocket is a communication standard for computer networks. It is a technical standard for realizing bidirectional communication in Web applications. )

![chat_demo](https://user-images.githubusercontent.com/38056766/69636388-d2bf1400-1099-11ea-8045-1d8a69df38a0.gif)

## Installation

### Requirements

- Golang ver1.12.6

### Setting

First of all, clone this repository.
Then excute `go get` the packages you need. Also create a file to write environment variables with the name `.envrc` and describe the following.

```
export GO111MODULE=on
export SECURITY_KEY=
export FACEBOOK_CLIENT=
export FACEBOOK_SECRET=
export FACEBOOK_CALL_BACK_URL=
export GITHUB_CLIENT=
export GITHUB_SECRET=
export GITHUB_CALL_BACK_URL=
export GOOGLE_CLIENT=
export GOOGLE_SECRET=
export GOOGLE_CALL_BACK_URL=
``` 

## Usage

Build and Start web server. Let's open the URL specified by the host in your browser. 

```
❯ go build -o chat

❯ ./chat -host=":8080"
```
