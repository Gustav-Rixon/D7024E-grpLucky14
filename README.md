# D7024e - kademlia-lab
## 2022 Lab group 14 
This is a lab for the course [Mobile and distributed computing systems](https://www.ltu.se/edu/course/D70/D7024E/D7024E-Mobila-och-distribuerade-datorsystem-1.67844) at [Luleå University of Technology](https://www.ltu.se/)

# Starting the project
## Prerequisites
* [Go](https://go.dev/)
* [Docker](https://www.docker.com/)

## Installation
```sh
   git clone https://github.com/your_username_/Project-Name.git
   ```
## Running
From the project directory run

```sh
   docker build --tag kademlia .
   docker compose up
   ```
### Initilizing the network
To initilize the network executing the following script
```sh
   ./init.sh
   ```
This scripts executes the join command onto a random chosen node.

## Interacting with the network
All nodes have an CLI that can be interacted with thru docker 
```sh
   docker exec -ti <container-id> cli <command>
   ```
## REST API
All nodes have a RESTFUL API with the following method

Method: GET, endpoint: /objects/{hash}/

The {hash} portion of the HTTP path is to be substituted for the hash of the object. A successful call should result in the contents of the object being responded with.

Method: POST, endpoint: /objects

Each message sent to the endpoint must contain a data object in its HTTP body. If the operation is successful, the node will reply with 201 CREATED containing both the contents of the uploaded object and a Location: /objects/{hash} header, where {hash} will be substituted for the hash of the uploaded object.

## Test
```sh
    go test ./... -coverprofile=coverage.out
    go tool cover -func coverage.out
   ```
