package bucket_test

import (
	"kademlia/internal/bucket"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBucket(t *testing.T) {
	// should create a new and empty bucket
	b := bucket.NewBucket()
	assert.NotNil(t, b)
	assert.Equal(t, b.Len(), 0)
}
