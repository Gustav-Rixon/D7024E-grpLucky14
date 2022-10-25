package rpcparser

import (
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"kademlia/internal/rpc"
	"kademlia/internal/rpc/commands/findnodeRPC"
	"kademlia/internal/rpc/commands/findnoderesponse"
	"kademlia/internal/rpc/commands/findvalueRPC"
	findvalueresp "kademlia/internal/rpc/commands/findvalueresponse"
	"kademlia/internal/rpc/commands/ping"
	"kademlia/internal/rpc/commands/pong"
	"kademlia/internal/rpc/commands/refresh"
	"kademlia/internal/rpc/commands/store"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseRPC(t *testing.T) {
	con := contact.Contact{ID: kademliaid.FromString("0")}
	con.CalcDistance(kademliaid.NewRandomKademliaID())

	type args struct {
		requestor *contact.Contact
		rpc       *rpc.RPC
	}
	tests := []struct {
		name    string
		args    args
		want    RPCCommand
		wantErr bool
	}{
		{name: "Test PING",
			args:    args{&con, &rpc.RPC{Content: "PING", RPCId: kademliaid.NewRandomKademliaID()}},
			want:    ping.Ping{},
			wantErr: false},

		{name: "Test PONG",
			args:    args{&con, &rpc.RPC{Content: "PONG", RPCId: kademliaid.NewRandomKademliaID()}},
			want:    pong.Pong{},
			wantErr: false},

		{name: "Test STORE",
			args:    args{&con, &rpc.RPC{Content: "STORE", RPCId: kademliaid.NewRandomKademliaID()}},
			want:    &store.Store{},
			wantErr: false},

		{name: "Test FIND_NODE",
			args:    args{&con, &rpc.RPC{Content: "FIND_NODE", RPCId: kademliaid.NewRandomKademliaID()}},
			want:    &findnodeRPC.FindNodeRPC{},
			wantErr: false},

		{name: "Test FIND_NODE_RESP",
			args:    args{&con, &rpc.RPC{Content: "FIND_NODE_RESP", RPCId: kademliaid.NewRandomKademliaID()}},
			want:    &findnoderesponse.FindNodeResponse{},
			wantErr: false},

		{name: "Test FIND_VALUE",
			args:    args{&con, &rpc.RPC{Content: "FIND_VALUE", RPCId: kademliaid.NewRandomKademliaID()}},
			want:    &findvalueRPC.FindValueRPC{},
			wantErr: false},

		{name: "Test FIND_VALUE_RESP",
			args:    args{&con, &rpc.RPC{Content: "FIND_VALUE_RESP", RPCId: kademliaid.NewRandomKademliaID()}},
			want:    &findvalueresp.FindValueResp{},
			wantErr: false},

		{name: "Test REFRESH",
			args:    args{&con, &rpc.RPC{Content: "REFRESH", RPCId: kademliaid.NewRandomKademliaID()}},
			want:    &refresh.RefreshRPC{},
			wantErr: false},

		{name: "Test unknown",
			args:    args{&con, &rpc.RPC{Content: "not a real RPC", RPCId: kademliaid.NewRandomKademliaID()}},
			want:    nil,
			wantErr: true},

		{name: "Test empty",
			args:    args{&con, &rpc.RPC{}},
			want:    nil,
			wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRPC(tt.args.requestor, tt.args.rpc)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRPC() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.IsType(t, got, tt.want)
		})
	}
}
