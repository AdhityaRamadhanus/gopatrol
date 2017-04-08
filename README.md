# checkupd

[![Go Report Card](https://goreportcard.com/badge/github.com/AdhityaRamadhanus/checkupd)](https://goreportcard.com/report/github.com/AdhityaRamadhanus/checkupd)  

self-hosted endpoint monitoring and status pages based on https://github.com/sourcegraph/checkup

<p>
  <a href="#installation">Installation |</a>
  <a href="#checklist">Checklist</a> |
  <a href="#checkupd">Checkupd</a> |
  <a href="#usage">Usage</a> |
  <a href="#licenses">License</a>
  <br><br>
  <blockquote>
	Checkupd is self-hosted health checks and status pages, written in Go using checkup and grpc as backend.

    It includes cli app called checklist to manage endpoint, setting up environment for status page and daemon to check endpoints called checkupd.

    There is much work to do for this project to be complete. Use it carefully.
  </blockquote>
</p>

Installation
------------
* git clone
* go get -v
* make (will create to executable on build/linux or build/mac)
* make build_docker (optional, create docker image checkup:v1.0.0)

Checklist
------------
* Checklist is command line tools to manage endpoint
* You can add http, tcp endpoint and delete it, list all the endpoints currently being checked and check them or setup status page
* This cli will connect to localhost:9009 by default since checkupd is running on that port, you can override it with --host flag in every command
```
NAME:
   checklist - Checkup server cli 

USAGE:
   checklist [global options] command [command options] [arguments...]

VERSION:
   1.0.0

AUTHOR:
   Adhitya Ramadhanus <adhitya.ramadhanus@gmail.com>

COMMANDS:
     add-http    Add endpoints to checkup
     add-tcp     Add tcp endpoints to checkup
     check       list and check endpoints
     list        list endpoint
     delete      delete endpoint
     setup-page  setup status page
     help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

Checkupd
------------
* Checkupd is a daemon that running in background to check all endpoints specified by users
* It needs config file as defined here https://github.com/sourcegraph/checkup
* it's recommended to run this as docker container through docker-compose i provide in this repo
```
./checkupd --config=<path to checkup.json>
```

Usage
------------
* The simplest way to run this is using docker container and setup the status page using checklist
```
1. make build_docker
2. make 
3. ./build/{linux,mac}/checklist setup-page
4. docker-compose up -d
```

License
----

MIT © [Adhitya Ramadhanus]

