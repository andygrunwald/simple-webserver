# Change Log

All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

### Added

* Disk-Attribute to marathon.json (native app)
* Disk-Attribute to marathon-docker.json (docker app)
* New deployment method: Marathon incl. Redis backend (docker container)
* Table of contents
* New deployment method: Docker Compose

### Fixed

* Changed Marathon definitions to new v1.0.0 version
* EXPOSE-Port of Dockerfile (from 8080 to 8082)
* New lines after body for `/ping` and `/version`

## [v1.0.0] - 2016-06-18

### Added

* JSON definition to deploy simple-webserver to Marathon (based on Apache Mesos)
* Dockerfile to deploy simple-webserver with Docker
* JSON definition to deploy simple-webserver in a Docker container to Marathon
* HTTP-Endpoint `/version` that will respond with the Name and Version of this app with HTTP status code 200
* HTTP-Endpoint `/kill` that will kill the webserver
* HTTP request logging to stdout
* This changelog
* "Contact" section to README
* Recommendation for bobrik/mesos-compose to README
* GoDoc
* Unit tests
* TravisCI
* "Things on the list to try" section to README
* Redis as /ping backend
* Instructions for native and Docker based start with a Redis backend

### Changed

* Deployment for Marathon Docker to Port 11001
* Minimized hardware req. for simple-webserver (native) on Marathon
* Removed \n from pong response

### Fixed

* Major refactoring (incl. more documentation) of README.md

## [v0.0.1] - 2015-09-09

### Added

* HTTP-Endpoint `/ping` that will respond static `pong` with HTTP status code 200
* HTTP-Endpoint `/` that redirects to `/ping` with HTTP status code 303 See Other
* `-listen` flag to define IP and Port for HTTP server binding

[Unreleased]: https://github.com/andygrunwald/simple-webserver/compare/v1.0.0...HEAD
[v1.0.0]: https://github.com/andygrunwald/simple-webserver/releases/tag/v1.0.0
[v0.0.1]: https://github.com/andygrunwald/simple-webserver/releases/tag/v0.0.1
