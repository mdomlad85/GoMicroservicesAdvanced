FROM golang:1.9.2-alpine

RUN apk update && apk upgrade && apk add --no-cache bash git
RUN go get github.com/gin-gonic/gin
RUN go get github.com/mdomlad85/GoMicroservices

ENV SOURCES /GitProjects/GoProjects/src/github.com/mdomlad85/GoMicroservicesAdvanced/Frameworks/
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go build

WORKDIR ${SOURCES}
CMD ${SOURCES}GoMicroservicesAdvanced
EXPOSE 8080
