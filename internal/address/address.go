package address

import (
	"os"
	"strconv"
)

type Address struct {
	//hostAddress string
	hostPort string
}

func Newaddress(address string) string {
	listenport := os.Getenv("LISTEN_PORT")

	return listenport
}

func (address *Address) GetPort() (int, error) {
	return strconv.Atoi(address.hostPort)
}
