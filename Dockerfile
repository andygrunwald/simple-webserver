FROM golang:onbuild

MAINTAINER Andy Grunwald <andygrunwald@gmail.com>

EXPOSE 8082
ENTRYPOINT ["app"]
