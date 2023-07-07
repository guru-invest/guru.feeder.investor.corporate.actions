FROM golang:alpine
LABEL maintainer="Tiago Sanches <tiago@guru.com.vc>"

RUN apk update && \
    apk add --no-cache git

WORKDIR /guru.feeder.investor.corporate.actions

COPY . .

# RUN apk add tzdata && \
#     cp /usr/share/zoneinfo/America/Sao_Paulo /etc/localtime && \
#     echo "America/Sao_Paulo" > /etc/timezone && \
#     date && \
#     apk del tzdata

RUN go build -o guru.feeder.investor.corporate.actions ./cmd/guru.feeder.investor.corporate.actions.go


CMD ["./guru.feeder.investor.corporate.actions"]