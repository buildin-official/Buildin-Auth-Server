FROM golang:1.20.2-alpine AS builder

ENV GOPATH=/go/src
WORKDIR /go/src/pentag.kr/BuildinAuth

COPY ["./go.sum", "./go.mod", "/go/src/pentag.kr/BuildinAuth/"]
RUN go mod download

COPY ./ ./
RUN go build -o main .

FROM scratch

COPY --from=builder ["/go/src/pentag.kr/BuildinAuth/main", "/"]

EXPOSE 3000

ENTRYPOINT ["/main"]