package restAPI_test

import (
	"kademlia/internal/contact"
	"kademlia/internal/datastore"
	"kademlia/internal/kademliaid"
	restAPI "kademlia/internal/network/restAPI"
	"kademlia/internal/node"
	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
)

type HttpWriterMock struct {
	mock.Mock
}

func (m *HttpWriterMock) Write(data []byte) (int, error) {
	args := m.Called(data)
	return args.Int(0), nil
}

func (mock *HttpWriterMock) WriteHeader(statusCode int) {
}

func (mock *HttpWriterMock) Header() http.Header {
	return nil
}

func TestGet(t *testing.T) {
	var writerMock *HttpWriterMock
	var n node.Node
	var handler restAPI.RPCHandler
	var req *http.Request

	//Should tell user about missing hash
	writerMock = new(HttpWriterMock)
	req, _ = http.NewRequest("GET", "/objects/", nil)

	n = node.Node{}
	handler = restAPI.RPCHandler{Node: &n}
	writerMock.On("Write", []byte("Missing hash")).Return(0, nil)
	handler.Get(writerMock, req)
	writerMock.AssertExpectations(t)

	// Should return the value if found
	n = node.Node{}
	writerMock = new(HttpWriterMock)
	writerMock.On("Write", mock.Anything).Return(0, nil)
	n.DataStore = datastore.New()
	msg := "test"
	hash := kademliaid.NewKademliaID(&msg)
	req, _ = http.NewRequest("POST", "/objects/"+hash.String(), nil)
	contacts := &[]contact.Contact{}
	n.DataStore.Insert(msg, contacts, nil, true)
	handler = restAPI.RPCHandler{Node: &n}
	handler.Get(writerMock, req)
	writerMock.AssertCalled(t, "Write", []byte(msg+", from local node"))
}
