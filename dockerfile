FROM debian

RUN apt-get update && apt-get install -y wget && apt-get install -y tar
RUN wget -c https://golang.org/dl/go1.15.8.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.15.8.linux-amd64.tar.gz
# RUN export PATH=$PATH:/usr/local/go/bin
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR go-app

COPY . .

RUN go build -o ./app cmd/main/main.go

ENTRYPOINT /go-app/app