# simple-webserver

[![GoDoc](https://godoc.org/github.com/andygrunwald/simple-webserver?status.svg)](https://godoc.org/github.com/andygrunwald/simple-webserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/andygrunwald/simple-webserver)](https://goreportcard.com/report/github.com/andygrunwald/simple-webserver)

A small webserver for testing various technologies, techniques and concepts written in [Go](http://golang.org/).

## Batteries Included

Here you will find a list of tested technologies, techniques and concepts within this project:

* [Go](http://golang.org/): How to write a small webserver in this programing language
* [Docker](https://www.docker.com/): How to isolate a Go app in a docker container
* [Marathon](https://mesosphere.github.io/marathon/) @ [Apache Mesos](http://mesos.apache.org/): How to deploy this app on a Marathon cluster (native and in a docker container)

## Why this project?

The information technology and computer science world is getting crazy.
Every day new tooling will be released, new techniques comes up and some concepts are getting best practice.

**It is important to stay up to date and this is the really hard part here.**

This project is my personal playground to try out all the new things i am excited of.
I am aware that this project is not representative and comparable with a real products that solves problems of the world.
Those projects are much more complex and have several other dependencies.
But the simplicity of this projects is on purpose and enables me to get a first idea and feeling of the topic i want to try in a small amount of time.

## Build instructions

### Native application

For building the native webserver you need a running [Go installation](https://golang.org/doc/install).

```sh
$ go build .
```

To compile it for a different OS (e.g. linux) use

```sh
$ GOOS=linux go build
```

Precompiled binaries for various operating systems can be found in our [releases](https://github.com/andygrunwald/simple-webserver/releases).

### Docker

For building the docker image you need a running [Docker engine](https://docs.docker.com/engine/installation/).

```sh
$ docker build -t andygrunwald/simple-webserver .
```

Or you use the existing images from [Docker Hub](https://hub.docker.com/r/andygrunwald/simple-webserver/).

## Tryout everything

### Native application

Start the webserver with

```sh
$ ./simple-webserver
2016/06/05 14:57:18 Starting webserver and listen on :8082
```

and send your first request

```sh
$ curl http://localhost:8082/version
simple webserver v0.1.0-dev
```

### Docker

Start the container with

```sh
$ docker run -d -p 8082:8082 andygrunwald/simple-webserver
9b20babef443420bf1728583fbfba......
```

and send your first request

```sh
$ curl http://DOCKER-IP:8082/ping
pong
```

For further information the docker image is available at [Docker Hub](https://hub.docker.com/r/andygrunwald/simple-webserver/).

### Marathon (native application)

Start the native application on a [Marathon](https://github.com/mesosphere/marathon) cluster (scheduling framework for [Apache Mesos](http://mesos.apache.org/) with

```sh
$ curl -X POST -H "Content-Type: application/json" http://marathon:8080/v2/apps -d@marathon.json
{
    "id": "/simple-webserver",
    ...
}
```

Please visit your Marathon UI at http://marathon:8080/ to see if the deployment works out.

If you use Marathon in combination with [haproxy-marathon-bridge](https://open.mesosphere.com/tutorials/service-discovery/) or another [service discovery](https://mesosphere.github.io/marathon/docs/service-discovery-load-balancing.html) solution your service is available via the **service port 11000**

```sh
$ curl http://marathon-node:11000/ping
pong
```

### Marathon (docker container)

Start the docker container on a [Marathon](https://github.com/mesosphere/marathon) cluster (scheduling framework for [Apache Mesos](http://mesos.apache.org/) with

```sh
$ curl -X POST -H "Content-Type: application/json" http://marathon:8080/v2/apps -d@marathon-docker.json
{
    "id":"/simple-webserver-docker",
    ...
}
```

Please visit your Marathon UI at http://marathon:8080/ to see if the deployment works out.

If you use Marathon in combination with [haproxy-marathon-bridge](https://open.mesosphere.com/tutorials/service-discovery/) or another [service discovery](https://mesosphere.github.io/marathon/docs/service-discovery-load-balancing.html) solution your service is available via the **service port 11001**

```sh
$ curl http://marathon-node:11001/ping
pong
```

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
