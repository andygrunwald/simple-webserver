# simple-webserver

A simple webserver for testing purposes written in [Go](http://golang.org/).

With this webserver several tests can be done: Deployment, scheduling, scaling.
The main goal is to test and get experiments with new systems like [Docker](https://www.docker.com/), [Mesos](http://mesos.apache.org/), [Marathon](https://mesosphere.github.io/marathon/), [Kubernetes](http://kubernetes.io/) and others.

Below you will find how to deploy with:

* Docker
* Marathon
* Marathon + Docker

## Compile

```sh
$ go build -o simple-webserver
```

To compile it for a different OS (e.g. linux) use

```sh
$ GOOS=linux go build -o simple-webserver
```

### Build docker image

To build the Docker image execute ...

```sh
$ docker build -t andygrunwald/simple-webserver .
```

... or you use the existing images from [Docker Hub](https://hub.docker.com/r/andygrunwald/simple-webserver/).

## Usage

Starting the service:

```sh
$ ./simple-webserver
2015/09/09 14:41:26 Starting webserver and listen on :8082
```

Sending requests to it:

```sh
$ curl http://localhost:8082/ping
pong
```

or get the version:

```sh
$ curl http://localhost:8082/version
simple webserver v0.1.0-dev
```

## Binary

At [releases](https://github.com/andygrunwald/simple-webserver/releases) you will find this simple webserver as pre compiled binary for various operating systems.

This binary can be used for simple test purposes like running this via [Mesos](http://mesos.apache.org/)/[Marathon](https://github.com/mesosphere/marathon), [Docker](https://www.docker.com/), [rkt](https://github.com/coreos/rkt) or whatever you want to test.

## Deployment

### Docker

The docker image is available at [Docker Hub](https://hub.docker.com/r/andygrunwald/simple-webserver/).

```sh
$ docker pull andygrunwald/simple-webserver
$ docker run -d -p 8082:8082 andygrunwald/simple-webserver
9b20babef443420bf1728583fbfba......
```

After it you should be able to send a request:

```sh
$ curl http://docker-node:8082/ping
pong
```

### Marathon

To deploy this application to a [Marathon](https://github.com/mesosphere/marathon) cluster (scheduling framework for [Apache Mesos](http://mesos.apache.org/) use [marathon.json](./marathon.json) and call

```sh
$ curl -X POST -H "Content-Type: application/json" http://marathon:8080/v2/apps -d@marathon.json
{
    "id": "/simple-webserver",
    "cmd": "cd simple-webserver-v0.0.1-linux-amd64 && chmod +x simple-webserver && ./simple-webserver --listen \":$PORT0\"",
	...
}
```

If you use Marathon in combination with [haproxy-marathon-bridge](https://open.mesosphere.com/tutorials/service-discovery/) for [service discovery](https://mesosphere.github.io/marathon/docs/service-discovery-load-balancing.html) your service is available via

```sh
$ curl http://marathon-node:11000/ping
pong
```

### Marathon with Docker

To deploy this application in a docker container to a [Marathon](https://github.com/mesosphere/marathon) cluster (scheduling framework for [Apache Mesos](http://mesos.apache.org/) use [marathon-docker.json](./marathon-docker.json) and call

```sh
$ curl -X POST -H "Content-Type: application/json" http://marathon:8080/v2/apps -d@marathon-docker.json
{
    "id":"/simple-webserver-docker",
    "cmd":null,
	...
}
```

If you use Marathon in combination with [haproxy-marathon-bridge](https://open.mesosphere.com/tutorials/service-discovery/) for [service discovery](https://mesosphere.github.io/marathon/docs/service-discovery-load-balancing.html) your service is available via

```sh
$ curl http://marathon-node:11000/ping
pong
```

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).