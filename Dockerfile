FROM golang:1.19

WORKDIR /app

# Setup GO Stuff

COPY main.go go.mod go.sum ./

ADD src ./src
ADD static ./static
ADD cache ./cache

RUN go mod download

RUN go build -o /atc-scraper-go

# Setup Python Stuff

RUN apt update -y
RUN apt install -y python3-pip

ADD atc-decoder ./atc-decoder

RUN pip install -r ./atc-decoder/requirements.txt

# CMD [ "/atc-scraper-go" ]

