FROM debian AS build

RUN apt-get update && apt-get install -y wget && apt-get install -y tar
RUN wget -c https://golang.org/dl/go1.15.8.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.15.8.linux-amd64.tar.gz
ENV PATH="/usr/local/go/bin:${PATH}"

COPY . .
RUN go build -o ./app cmd/main/main.go

FROM debian AS final

WORKDIR go-app

COPY --from=build /app .

ENTRYPOINT /go-app/app