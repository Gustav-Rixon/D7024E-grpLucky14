FROM golang:1.19-alpine

WORKDIR /kademlia

COPY . .

#Creates the cli
RUN go build cmd/cli.go 
RUN go build -o /docker-gs-ping

ENV PATH /kademlia:$PATH

# Run
CMD [ "/docker-gs-ping" ]
