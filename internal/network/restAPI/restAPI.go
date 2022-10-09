package restAPI

import (
	"fmt"
	"io/ioutil"
	cmdparser "kademlia/internal/command/parser"
	"kademlia/internal/kademliaid"
	"kademlia/internal/node"
	"net/http"
	"path"

	"github.com/rs/zerolog/log"
)

type RPCHandler struct {
	Node *node.Node
}

// Get endpoint
func (handler *RPCHandler) Get(w http.ResponseWriter, r *http.Request) {
	log.Trace().Msg("GET request received")
	hash := path.Base(r.URL.String())

	//If the Get command only retrives the endpoint the hash is missing.
	if hash == "objects" {
		log.Info().Msg("GET request received: Hash missing")
		w.Write([]byte("Missing hash"))
		return
	}

	cmd := cmdparser.ParseCmd(fmt.Sprintf("get %s", hash))
	value, err := cmd.Execute(handler.Node)
	if err != nil {
		log.Error().Msg("Failed to execute get command")
	}
	w.Write([]byte(value))
}

// Post endpoint
func (handler *RPCHandler) Post(w http.ResponseWriter, r *http.Request) {
	log.Trace().Msg("POST request received")
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error().Msgf("Error reading request body %s", err.Error())
		return
	}

	cmd := cmdparser.ParseCmd(fmt.Sprintf("put %s", string(data)))
	value, err := cmd.Execute(handler.Node)
	if err != nil {
		log.Error().Str("Error", err.Error()).Msg("Failed to execute put command")
		w.Write([]byte(err.Error()))
		return
	}

	sData := string(data)
	hash := kademliaid.NewKademliaID(&sData)
	w.Header().Add("Location", "/objects/"+hash.String())
	w.WriteHeader(201) // reply with 201

	w.Write([]byte(value))
}

func Listen(node *node.Node) error {
	requesthandler := RPCHandler{Node: node}

	http.HandleFunc("/objects/", requesthandler.Get)
	http.HandleFunc("/objects", requesthandler.Post)

	err := http.ListenAndServe(":8090", nil)
	return err
}
