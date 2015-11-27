# Change Log

All notable changes to this project will be documented in this file.
This project adheres to [Semantic Versioning](http://semver.org/).

## [Unreleased]

### Added

* JSON definition to deploy simple-webserver to Marathon (based on Apache Mesos)
* Dockerfile to deploy simple-webserver with Docker
* JSON definition to deploy simple-webserver in a Docker container to Marathon
* HTTP-Endpoint `/version` that will respond with the Name and Version of this app with HTTP status code 200
* HTTP request logging to stdout
* This changelog

### Fixed

* More documentation in README.md

## [v0.0.1] - 2015-09-09

### Added

* HTTP-Endpoint `/ping` that will respond static `pong` with HTTP status code 200
* HTTP-Endpoint `/` that redirects to `/ping` with HTTP status code 303 See Other
* `-listen` flag to define IP and Port for HTTP server binding

[Unreleased]: https://github.com/andygrunwald/simple-webserver/compare/v0.0.1...HEAD
[v0.0.1]: https://github.com/andygrunwald/simple-webserver/releases/tag/v0.0.1