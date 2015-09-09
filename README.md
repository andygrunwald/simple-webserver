# simple-webserver

A simple webserver for testing purposes written in [Go](http://golang.org/).

## Compile

```sh
$ go build
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

## License

This project is released under the terms of the [MIT license](http://en.wikipedia.org/wiki/MIT_License).