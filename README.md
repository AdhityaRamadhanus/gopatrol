# gopatrol

[![Go Report Card](https://goreportcard.com/badge/github.com/AdhityaRamadhanus/gopatrol)](https://goreportcard.com/report/github.com/AdhityaRamadhanus/gopatrol)  [![Build Status](https://travis-ci.org/AdhityaRamadhanus/gopatrol.svg?branch=refactor-rest)](https://travis-ci.org/AdhityaRamadhanus/gopatrol)

self-hosted endpoint monitoring daemon with centralized events log using MongoDB based on https://github.com/sourcegraph/checkup

<p>
  <a href="#installation">Installation |</a>
  <a href="#setting-up-gopatrol-api-server">Setting API Server |</a>
  <a href="#setting-up-gopatrol-daemon">Setting Daemon |</a>
  <a href="#interacting-with-api-using-cli">CLI |</a>
  <a href="#interacting-with-api-using-dashboard">CLI |</a>
  <a href="#notifier-slack">Slack Notifier |</a>
  <a href="#notifier-email">Email Notifier |</a>
  <a href="#licenses">License</a>
  <br><br>
  <blockquote>
	gopatrol is self-hosted health checks, written in Go using checkup (instead of using them as dependency i decide to copy the file to this project) and restful api to interact with.

  There is much work to do for this project to be complete. Use it carefully.

  gopatrol currently supports:

  - Checking HTTP endpoints
  - Checking TCP endpoints (TLS supported)
  - Checking of DNS services & record existence  
  - Storing events in MongoDB
  - Add delete update checkers with dashboard/cli
  - Easy to setup and deploy
  - Get notified via slack and email (need help with email notifier)
  </blockquote>
</p>

Installation
----------- 
* git clone
* make
```bash
NAME:
   gopatrol-cli - gopatrol cli 

USAGE:
   gopatrol-cli [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Adhitya Ramadhanus <adhitya.ramadhanus@gmail.com>

COMMANDS:
     add-http  Add endpoints to checkup
     add-tcp   Add tcp endpoints to checkup
     add-dns   Add dns endpoints to checkup
     list      list endpoint
     delete    delete endpoint
     help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

```bash
NAME:
   gopatrol - gopatrol daemon 

USAGE:
   gopatrol [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Adhitya Ramadhanus <adhitya.ramadhanus@gmail.com>

COMMANDS:
     daemon   run gopatrol checking daemon
     api      run gopatrol api server
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

### Setting up gopatrol API Server

There's 2 different setup configuration for cli and dashboard, you can interact with the api using cli only if you run the api in unix domain socket

For these two setup you also need to set env var, i suggest using .env file to store your config

Example of .env
```
MONGODB_URI="mongodb://localhost:27017/gopatrol"
JWT_SECRET="Something"
ENV="Development"
```

1. **Setup For CLI**
  ```bash
  $ gopatrol api --log=<log output, stdout, stderr or filename> --proto=unix
  ```
  it will run on a socket file /tmp/gopatrol.sock

2. **Setup For Dashboard**
  ```bash
  $ gopatrol api --log=<log output, stdout, stderr or filename> --proto=http --address=:3000
  ```
  it will run on port 3000

### Setting up gopatrol daemon
You also need to set env var for daemon to work mostly for mongodb and slack notifier, i suggest using .env file to store your config

Example of .env
```
MONGODB_URI="mongodb://localhost:27017/gopatrol"
SLACK_TOKEN="xoxb-something-something" //bot users token
SLACK_CHANNEL="something" // channel ID 
ENV="Development"
```

```bash
$ gopatrol daemon --log=<log output, stdout, stderr or filename> interval (10s, 1m, etc)
```

### Interacting with API using CLI
1. Adding Tcp endpoint
```bash
NAME:
   gopatrol-cli add-tcp - Add tcp endpoints to checkup

USAGE:
   gopatrol-cli add-tcp [command options] name url

OPTIONS:
   --attempts value, -a value         how many times to check endpoint (default: 5)
   --thresholdrtt value, --rtt value  Threshold Rtt to define a degraded endpoint (default: 0)
   --tls-enabled                      Enable TLS connection to endpoint
   --tls-ca value                     Certificate file to established tls connection
   --tls-skip-verify                  Skip verify tls certificate
   --timeout value                    Timeout to established a tls connection (default: 3000000000)
```

2. Adding Http endpoint
```bash
NAME:
   gopatrol-cli add-http - Add endpoints to checkup

USAGE:
   gopatrol-cli add-http [command options] name url

OPTIONS:
   --attempts value, -a value         how many times to check endpoint (default: 5)
   --thresholdrtt value, --rtt value  Threshold Rtt to define a degraded endpoint (default: 0)
   --mustcontain value                HTML content that a page should contain to determine whether a page is up or down
   --mustnotcontain value             HTML content that a page should not contain to determine whether a page is up or down
   --headers value                    Http Headers to send along the check request
   --upstatus value                   Http status code to define a healthy page (default: 200)

```

3. Adding DNS endpoint
```bash
NAME:
   gopatrol-cli add-dns - Add dns endpoints to checkup

USAGE:
   gopatrol-cli add-dns [command options] name url hostname

OPTIONS:
   --tls                              Send request over tls
   --host value                       grpc server address (default: "/tmp/gopatrol.sock")
   --attempts value, -a value         how many times to check endpoint (default: 5)
   --thresholdrtt value, --rtt value  Threshold Rtt to define a degraded endpoint (default: 0)
   --timeout value                    Timeout to established a tls connection (default: 3000000000)
```

4. Deleting Endpoint
```bash
NAME:
   gopatrol-cli delete - delete endpoint

USAGE:
   gopatrol-cli delete [command options] slug

```

5. Listing Endpoint
```bash
NAME:
   gopatrol-cli list - list endpoint

USAGE:
   gopatrol-cli list [command options] [arguments...]

```

### Interacting with API using Dashboard
Still in progress

![dashboard](https://cloud.githubusercontent.com/assets/5761975/26282662/e2fd6ba2-3e3f-11e7-9619-dee0f770e0e3.png)


Notifier (Slack)
----------------
* To use this notifier you need bot integration in your team and channel id where this bot will notify you, refer to this link https://api.slack.com/bot-users 
* After you get the tokens and channel ID, set it in env var ot use .env file

![slack-notifier](https://cloud.githubusercontent.com/assets/5761975/26282665/ed703178-3e3f-11e7-8335-b9ee78b369e5.png)

Notifier (Email)
----------------
* NIP, not in progress, but definitely in to-do list

License
----

MIT © [Adhitya Ramadhanus]

