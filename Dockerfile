FROM golang:onbuild

MAINTAINER Andy Grunwald <andygrunwald@gmail.com>

EXPOSE 8080
ENTRYPOINT ["app"]
