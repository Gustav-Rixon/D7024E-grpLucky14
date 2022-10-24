package contact

import (
	"kademlia/internal/address"
	"kademlia/internal/kademliaid"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContact(t *testing.T) {
	// Should be able to create a contact
	contact := NewContact(kademliaid.FromString("0000000000000000000000000000000000000001"), address.New("127.0.0.1:8888"))
	assert.IsType(t, Contact{}, contact)

	// Should be able to set distance from given ID
	contact.CalcDistance(kademliaid.FromString("0000000000000000000000000000000000000002"))
	assert.NotNil(t, contact.distance)

	// Should be able to see if one contact is closer than another
	contact2 := NewContact(kademliaid.FromString("0000000000000000000000000000000000000010"), address.New("127.0.0.1:8888"))
	contact2.CalcDistance(contact.GetDistance())
	assert.True(t, contact.Less(&contact2))

	// Should be able to make into string
	assert.IsType(t, "", contact.String())
}

func TestSerialize(t *testing.T) {
	contacts := []Contact{
		NewContact(kademliaid.FromString("0000000000000000000000000000000000000001"), address.New("127.0.0.1:8888")),
		NewContact(kademliaid.FromString("0000000000000000000000000000000000000002"), address.New("127.0.0.2:8888")),
		NewContact(kademliaid.FromString("0000000000000000000000000000000000000003"), address.New("127.0.0.3:8888")),
	}

	serializedContacts := SerializeContacts(contacts)
	splitContacts := strings.Split(serializedContacts, "%")
	var deserializedContacts []Contact
	var err error

	for _, con := range splitContacts {
		var deserCon *Contact
		err, deserCon = Deserialize(&con)
		deserializedContacts = append(deserializedContacts, *deserCon)
	}

	// Deserialized contacts should be 3
	assert.NoError(t, err)
	assert.Len(t, deserializedContacts, 3)

	// Should return an error when attempting to deserialize something that's not a contact
	notARealContact := "blah"
	err, _ = Deserialize(&notARealContact)
	assert.Error(t, err)
}

func TestContactCandidates(t *testing.T) {
	candidates := ContactCandidates{}

	contacts := []Contact{
		NewContact(kademliaid.FromString("0000000000000000000000000000000000000003"), address.New("127.0.0.3:8888")),
		NewContact(kademliaid.FromString("0000000000000000000000000000000000000002"), address.New("127.0.0.2:8888")),
		NewContact(kademliaid.FromString("0000000000000000000000000000000000000001"), address.New("127.0.0.1:8888")),
		NewContact(kademliaid.FromString("0000000000000000000000000000000000000004"), address.New("127.0.0.4:8888")),
		NewContact(kademliaid.FromString("0000000000000000000000000000000000000006"), address.New("127.0.0.6:8888")),
		NewContact(kademliaid.FromString("0000000000000000000000000000000000000005"), address.New("127.0.0.5:8888")),
	}

	for i := range contacts {
		contacts[i].CalcDistance(kademliaid.FromString("0000000000000000000000000000000000000000"))
		println(contacts[i].distance)
	}

	// Should be able to append a certain number of contacts and get a certain number via GetContacts
	candidates.Append(contacts)
	assert.Len(t, candidates.GetContacts(candidates.Len()), 6)

	// Should be able to sort contacts by distance, making contact #3 in contacts become #1 in candidates
	candidates.Sort()
	assert.Equal(t, candidates.Contacts[0], contacts[2])
}
