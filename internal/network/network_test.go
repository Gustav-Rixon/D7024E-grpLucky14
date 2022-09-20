package network

import (
	"kademlia/internal/node"
	"net"
	"reflect"
	"testing"
)

func TestGetOutboundIP(t *testing.T) {
	tests := []struct {
		name string
		want net.IP
	}{
		{"Check type", net.IPv4(255, 255, 255, 255)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOutboundIP(); !reflect.DeepEqual(reflect.TypeOf(got), reflect.TypeOf(tt.want)) {
				t.Errorf("GetOutboundIP() = %t, want %t", reflect.TypeOf(got), reflect.TypeOf(tt.want))
			}
		})
	}
}

func TestInitNetwork(t *testing.T) {
	type args struct {
		listenPort int
		sendPort   int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"Bad listen port", args{-1, 0}, true},
		{"Bad send port", args{0, -1}, true},
		{"Good ports", args{0, 0}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitNetwork(tt.args.listenPort, tt.args.sendPort); (err != nil) != tt.wantErr {
				t.Errorf("InitNetwork() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPing(t *testing.T) {
	NetInfo.listenPort = ":80"
	NetInfo.sendPort = "localhost:90"

	go Listen()

	node.CreateSelf([20]byte{}, net.IP{})
	SendPing(net.IPv4(255, 255, 255, 255))
}
