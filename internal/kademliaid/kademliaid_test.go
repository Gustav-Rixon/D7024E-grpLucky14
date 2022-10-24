package kademliaid

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIDGen(t *testing.T) {
	sampleID := "sample"
	newID := NewKademliaID(&sampleID)

	// Should receive a seeded ID (right side will change if we change hash type)
	assert.Equal(t, &newID, FromString("8151325dcdbae9e0ff95f9f9658432dbedfdb209"))

	// Should receive a 20 byte random ID
	newID = *NewRandomKademliaID()
	assert.Len(t, newID, 20)

	// Should be a string
	assert.Equal(t, reflect.TypeOf(newID.String()), reflect.TypeOf("string"))
}

func TestCompareID(t *testing.T) {
	// Test 1 is less than 2
	newID := FromString("0000000000000000000000000000000000000001")
	newID2 := FromString("0000000000000000000000000000000000000010")
	assert.True(t, newID.Less(newID2))
	assert.False(t, newID2.Less(newID))

	// Equal id's should not be less than one another
	newID2 = FromString("0000000000000000000000000000000000000001")
	assert.False(t, newID2.Less(newID))

	// Expected distance between newID and newID2 is 11
	newID2 = FromString("0000000000000000000000000000000000000010")
	dist := FromString("0000000000000000000000000000000000000011")
	assert.Equal(t, newID2.CalcDistance(newID), dist)

	// Should be equal
	newID2 = FromString("0000000000000000000000000000000000000001")
	assert.True(t, newID.Equals(newID2))

	// Should not be equal (chance 1 in a quadrillion they will be equal)
	newID = NewRandomKademliaID()
	assert.False(t, newID.Equals(newID2))
}
