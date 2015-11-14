FROM scratch

MAINTAINER Andy Grunwald <andygrunwald@gmail.com>

ADD simple-webserver simple-webserver

EXPOSE 8082

ENTRYPOINT ["./simple-webserver"]
