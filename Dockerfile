FROM golang:1.20.3-alpine as builder

ENV GOPATH=/go/src
WORKDIR /go/src/pentag.kr/BuildinAuth

COPY ["./go.sum", "./go.mod", "./"]
RUN go mod download

COPY ./ ./
RUN go build -o main .

FROM alpine

# Install Doppler CLI
RUN wget -q -t3 'https://packages.doppler.com/public/cli/rsa.8004D9FF50437357.key' -O /etc/apk/keys/cli@doppler-8004D9FF50437357.rsa.pub && \
    echo 'https://packages.doppler.com/public/cli/alpine/any-version/main' | tee -a /etc/apk/repositories && \
    apk add doppler

COPY --from=builder ["/go/src/pentag.kr/BuildinAuth/main", "/"]

CMD ["doppler", "run", "--", "/main"]