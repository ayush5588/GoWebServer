FROM golang:1.10
WORKDIR $GOPATH/src/github.com/ayush5588/GoWebServer/
ADD src src
EXPOSE 8080
CMD ["go","run","src/main.go"]
