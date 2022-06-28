FROM golang:alpine
LABEL maintainer="Victor Tadashi <victor.tadashi@guru.com.vc>"

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update && \
    apk add gcc libc-dev && \
    apk add --no-cache ca-certificates libc6-compat

WORKDIR $GOPATH/src/github.com/guru-invest/guru.feeder.investor.corporate.actions

COPY . .

# RUN apk add tzdata && \
#     cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime && \
#     echo "America/Sao_Paulo" > /etc/timezone && \
#     date && \
#     apk del tzdata

RUN go build -ldflags '-linkmode=external' -mod vendor -o guru.feeder.investor.corporate.actions ./cmd/guru.feeder.investor.corporate.actions.go

EXPOSE 8080

CMD ["./guru.feeder.investor.corporate.actions"]