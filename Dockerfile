# Multi-stage Docker build for AtcScraperGo application
FROM golang:1.19

WORKDIR /app

# Enhancement: Add comprehensive tests
COPY main.go go.mod go.sum ./

ADD src ./src
ADD static ./static
ADD cache ./cache
# Note: Consider refactoring approach

RUN go mod download

RUN go build -o /atc-scraper-go

CMD [ "/atc-scraper-go" ]

