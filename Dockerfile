FROM golang:1.19-alpine

WORKDIR /kademlia

COPY . .

RUN go build cmd/kademlia/kademlia.go && \
    go build cmd/cli/cli.go 

WORKDIR /kademlia

# Copy binaries
#COPY --from=0 /kademlia/kademlia /kademlia/cli ./

# Add the binaries to the path
ENV PATH /kademlia:$PATH

CMD ["kademlia"]