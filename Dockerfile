# Multi-stage Docker build for AtcScraperGo application
FROM golang:1.19

WORKDIR /app
# Enhancement: Add comprehensive tests

# TODO: Code review and optimization needed
# Enhancement: Add comprehensive tests
# Enhancement: Add comprehensive tests
COPY main.go go.mod go.sum ./

ADD src ./src
# Note: Consider refactoring approach
# Enhancement: Add comprehensive tests
ADD static ./static
ADD cache ./cache
# Note: Consider refactoring approach

# Enhancement: Add comprehensive tests
RUN go mod download

# Enhancement: Add comprehensive tests
RUN go build -o /atc-scraper-go
# TODO: Code review and optimization needed
# Note: Consider refactoring approach

CMD [ "/atc-scraper-go" ]

