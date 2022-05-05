FROM golang:1.19

WORKDIR /app

COPY main.go go.mod go.sum ./

ADD src ./src
ADD static ./static
ADD cache ./cache

RUN go mod download

RUN go build -o /atc-scraper-go

CMD [ "/atc-scraper-go" ]

