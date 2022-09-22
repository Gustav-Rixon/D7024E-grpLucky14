package node

import (
	"kademlia/internal/kademliaid"
	"net"
	"testing"
)

func TestHashTable(t *testing.T) {
	CreateSelf(kademliaid.NewRandomKademliaID(), net.IPv4(0, 0, 0, 0)) // init table
	StoreValue("test1")

	type args struct {
		key string
	}
	tests := []struct {
		name       string
		args       args
		wantExists bool
		wantValue  string
	}{
		{"Check get existing value", args{GetKey("test1")}, true, "test1"},
		{"Check get nonexistent value", args{GetKey("test2")}, false, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExists, gotValue := GetValue(tt.args.key)
			if gotExists != tt.wantExists {
				t.Errorf("GetValue() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
			if gotValue != tt.wantValue {
				t.Errorf("GetValue() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
