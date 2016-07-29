FROM alpine:latest

MAINTAINER Fernando Nogueira

WORKDIR "/opt"

ADD .docker_build/golly /opt/bin/golly
ADD ./templates /opt/templates
ADD ./static /opt/static

CMD ["/opt/bin/golly"]

