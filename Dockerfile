FROM golang:1.20.3

ENV GOPATH=/go/src
WORKDIR /go/src/pentag.kr/BuildinAuth

# Install Doppler CLI
RUN apt-get update && apt-get install -y apt-transport-https ca-certificates curl gnupg && \
    curl -sLf --retry 3 --tlsv1.2 --proto "=https" 'https://packages.doppler.com/public/cli/gpg.DE2A7741A397C129.key' | apt-key add - && \
    echo "deb https://packages.doppler.com/public/cli/deb/debian any-version main" | tee /etc/apt/sources.list.d/doppler-cli.list && \
    apt-get update && \
    apt-get -y install doppler

COPY ["./go.sum", "./go.mod", "/go/src/pentag.kr/BuildinAuth/"]
RUN go mod download

COPY ./ ./

RUN go build -o main .

CMD ["doppler", "run", "--", "/go/src/pentag.kr/BuildinAuth/main"]