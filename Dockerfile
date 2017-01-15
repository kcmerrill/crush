FROM golang:1.7
MAINTAINER kc merrill <kcmerrill@gmail.com>

RUN apt-get -y update

COPY . /code

RUN go get -u github.com/kcmerrill/queued

EXPOSE 80

ENTRYPOINT ["queued"]
CMD ["--port", "80"]
