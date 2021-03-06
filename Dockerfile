FROM golang:1.9
MAINTAINER kc merrill <kcmerrill@gmail.com>
COPY . /go/src/github.com/kcmerrill/crush
WORKDIR /go/src/github.com/kcmerrill/crush
RUN  go build -ldflags "-X main.Commit=`git rev-parse HEAD` -X main.Version=0.1.`git rev-list --count HEAD`" -o /usr/local/bin/crush
EXPOSE 80
EXPOSE 8080
ENTRYPOINT ["crush"]
