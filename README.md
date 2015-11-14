# simple-webserver

A simple webserver for testing purposes written in [Go](http://golang.org/).

## Compile

```sh
$ go build -o simple-webserver
```

To compile it for a different OS (e.g. linux) use

```sh
$ GOOS=linux go build -o simple-webserver
```

## Usage

```sh
$ ./simple-webserver
2015/09/09 14:41:26 Starting webserver and listen on :8082

$ curl http://localhost:8082/ping
pong
```

## Binary

Under [releases](https://github.com/andygrunwald/simple-webserver/releases) you will find this simple webserver as pre compiled binary for various operating systems.

This binary can be used for simple test purposes like running this via [Mesos](http://mesos.apache.org/)/[Marathon](https://github.com/mesosphere/marathon), [Docker](https://www.docker.com/), [rkt](https://github.com/coreos/rkt) or whatever you want to test.

## Deployment

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

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).