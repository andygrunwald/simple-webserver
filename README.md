# simple-webserver

[![Build Status](https://travis-ci.org/andygrunwald/simple-webserver.svg?branch=master)](https://travis-ci.org/andygrunwald/simple-webserver)
[![GoDoc](https://godoc.org/github.com/andygrunwald/simple-webserver?status.svg)](https://godoc.org/github.com/andygrunwald/simple-webserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/andygrunwald/simple-webserver)](https://goreportcard.com/report/github.com/andygrunwald/simple-webserver)

A small webserver (written in [Go](http://golang.org/)) for testing various technologies, techniques and concepts like [Docker](https://www.docker.com/), [Marathon](https://mesosphere.github.io/marathon/) / [Apache Mesos](http://mesos.apache.org/), [API Blueprint](https://apiblueprint.org/) and others.

All guides how to use, play and use a technique or concept are documented here. Feel free to play around and learn as much as i do.

---

## Table of Contents

1. [Batteries Included](#batteries-included)
2. [Why this project?](#why-this-project)
3. [Build instructions](#build-instructions)
	1. [Native application](#native-application)
	2. [Docker](#docker)
4. [Tryout everything](#tryout-everything)
	1. [Native application](#native-application-1)
	2. [Docker](#docker-1)
	3. [Marathon (native application)](marathon-native-application)
	4. [Marathon (docker container)](#marathon-docker-container)
	5. [Marathon incl. Redis backend (docker container)](#marathon-incl-redis-backend-docker-container)
5. [Things on the list to try](#things-on-the-list-to-try)
6. [Contact](#contact)
7. [License](#license)

## Batteries Included

Here you will find a list of tested technologies, techniques and concepts within this project:

* [Go](http://golang.org/): How to write a small webserver with a minimal [Redis](http://redis.io/) backend in this programming language (incl. [GoDoc](https://godoc.org/github.com/andygrunwald/simple-webserver) and [unit tests](./main_test.go))
* [Travis CI](https://travis-ci.org/): Execute unit tests on every push and pull request
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
$ go build -o simple-webserver .
```

To compile it for a different OS (e.g. linux) use

```sh
$ GOOS=linux go build -o simple-webserver .
```

Precompiled binaries for various operating systems can be found in our [releases](https://github.com/andygrunwald/simple-webserver/releases).

If you want to execute the unit tests, run

```sh
$ go test -v ./...
```

### Docker

For building the docker image you need a running [Docker engine](https://docs.docker.com/engine/installation/).

```sh
$ docker build -t andygrunwald/simple-webserver .
```

Or you use the existing images from [Docker Hub](https://hub.docker.com/r/andygrunwald/simple-webserver/).

## Tryout everything

### Native application

The general usage of *simple-webserver* is

```sh
$ ./simple-webserver -help
Usage of ./simple-webserver:
  -listen string
    	Address + Port to listen on. Format ip:port. (default ":8082")
  -redis string
    	Address + Port where a redis server is listening. (default ":6379")
```

Start the webserver with

```sh
$ ./simple-webserver
2016/06/05 14:57:18 Starting webserver and listen on :8082
```

and send your first request

```sh
$ curl http://localhost:8082/version
simple webserver v1.0.0
```

If you have a Redis server available, you can apply it via command line flag:

```sh
$ ./simple-webserver -redis ":6379"
2016/06/05 14:57:18 Starting webserver and listen on :8082
```

and call a `/ping` request

```sh
$ curl http://localhost:8082/ping
PONG
```

### Docker

Start the container with

```sh
$ docker run -d -p 8082:8082 andygrunwald/simple-webserver
9b20babef443420bf1728583fbfba......
```

and send your first request

```sh
$ curl http://DOCKER-IP:8082/version
simple webserver v1.0.0
```

Or launch *simple-webserver* with a Redis backend

```sh
$ docker run --name simple-webserver-redis -d redis
$ docker run -d -p 8082:8082 \
				--link simple-webserver-redis:redis \
				andygrunwald/simple-webserver -redis "redis:6379"
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
$ curl http://marathon-node:11000/version
simple webserver v1.0.0
```

If you don`t have a Mesos / Marathon cluster, i highly recommend [bobrik/mesos-compose](https://github.com/bobrik/mesos-compose).

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
$ curl http://marathon-node:11001/version
simple webserver v1.0.0
```

If you don`t have a Mesos / Marathon cluster, i highly recommend [bobrik/mesos-compose](https://github.com/bobrik/mesos-compose).

### Marathon incl. Redis backend (docker container)

Start the docker container with a Redis backend on a [Marathon](https://github.com/mesosphere/marathon) cluster (scheduling framework for [Apache Mesos](http://mesos.apache.org/) with

```sh
$ curl -X POST -H "Content-Type: application/json" http://marathon:8080/v2/groups -d@marathon-redis-docker.json
{
    "version":"2016-06-18T16:01:45.739Z",
    ...
}
```

Please visit your Marathon UI at http://marathon:8080/ to see if the deployment works out.

If it doesn't work out, please check `constraints` value in [marathon-redis-docker.json](./marathon-redis-docker.json) if this is equal the IP of your Mesos slave. If not, adjust it accordingly.

If you use Marathon in combination with [haproxy-marathon-bridge](https://open.mesosphere.com/tutorials/service-discovery/) or another [service discovery](https://mesosphere.github.io/marathon/docs/service-discovery-load-balancing.html) solution your service is available via the **service port 11001**

```sh
$ curl http://marathon-node:11001/ping
PONG
```

If you don`t have a Mesos / Marathon cluster, i highly recommend [bobrik/mesos-compose](https://github.com/bobrik/mesos-compose).

## Things on the list to try

* [Docker Compose](https://docs.docker.com/compose/overview/)
* [API Blueprint](https://apiblueprint.org/) specs
* [Blue-Green Deployment with Marathon](https://mesosphere.github.io/marathon/docs/blue-green-deploy.html)

## Contact

You are more than welcome to [open issues](https://github.com/andygrunwald/simple-webserver/issues) and / or [send pull requests](https://github.com/andygrunwald/simple-webserver/pulls) if you found a typo, a bug, need help to try something or want to request a new feature.

**If you appreciate this project please feel free to drop me a line and tell me!*
It's always nice to hear from people who have benefitted from this work and pushes the motivation again.**

* Twitter: [@andygrunwald](https://twitter.com/andygrunwald)
* Email: Checkout the mail at [my github profile](https://github.com/andygrunwald)

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).
