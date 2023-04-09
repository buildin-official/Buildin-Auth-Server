FROM golang:1.20.3

ENV GOPATH=/go/src
WORKDIR /go/src/pentag.kr/BuildinAuth

COPY ["./go.sum", "./go.mod", "/go/src/pentag.kr/BuildinAuth/"]
RUN go mod download

COPY ./ ./

RUN go build -o main .

EXPOSE 3000
CMD ["/bin/sh", "-c", "/go/src/pentag.kr/BuildinAuth/main"]

# FROM scratch

# COPY --from=builder ["/go/src/pentag.kr/BuildinAuth/main", "/"]


# RUN chmod +x /main
# CMD ["sh", "-c", "/main"]