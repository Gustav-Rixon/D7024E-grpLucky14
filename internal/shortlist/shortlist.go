package shortlist

import (
	"kademlia/internal/contact"
	"kademlia/internal/kademliaid"
	"sort"
)

const k = 5 // k-closest

type Entry struct {
	Contact       contact.Contact
	Probed        bool
	Active        bool
	ReturnedValue bool
}

type Shortlist struct {
	Entries [k]*Entry
	Closest *contact.Contact
	Target  *kademliaid.KademliaID
}

func (sl *Shortlist) Swap(i, j int) {
	sl.Entries[i], sl.Entries[j] = sl.Entries[j], sl.Entries[i]
}

func (sl *Shortlist) Less(i, j int) bool {
	if sl.Entries[j] == nil {
		return true
	}
	if sl.Entries[i] == nil {
		return false
	}
	return sl.Entries[i].Contact.Less(&sl.Entries[j].Contact)
}

func (shortlist *Shortlist) Len() int {
	length := 0
	for _, entry := range shortlist.Entries {
		if entry != nil {
			length++
		}
	}
	return length
}

func NewShortlist(target *kademliaid.KademliaID, candidates []contact.Contact) *Shortlist {
	shortlist := &Shortlist{}
	shortlist.Closest = &candidates[0]
	shortlist.Target = target
	for i, contact := range candidates {
		shortlist.Entries[i] = &Entry{contact, false, false, false}
	}
	return shortlist
}

func (sl *Shortlist) GetContacts() []contact.Contact {
	contacts := []contact.Contact{}
	for _, entry := range sl.Entries {
		if entry != nil {
			contacts = append(contacts, entry.Contact)
		}
	}
	return contacts
}

func (sl *Shortlist) Add(c *contact.Contact) {
	for _, entry := range sl.Entries {
		if entry != nil {
			if entry.Contact.ID.Equals(c.ID) {
				return
			}
		}
	}

	c.CalcDistance(sl.Target)

	if sl.Len() == k {
		if c.Less(&sl.Entries[k-1].Contact) {
			sl.Entries[k-1] = &Entry{Contact: *c, Active: false, Probed: false}
		}
	} else {
		for i := 0; i < len(sl.Entries); i++ {
			if sl.Entries[i] == nil {
				sl.Entries[i] = &Entry{Contact: *c, Active: false, Probed: false}
				break
			}
		}
	}

	sort.Sort(sl)
	sl.Closest = &sl.Entries[0].Contact
}

func (sl *Shortlist) Drop(c *contact.Contact) {
	for i := 0; i < len(sl.Entries); i++ {
		if sl.Entries[i] != nil {
			if sl.Entries[i].Contact.ID.Equals(c.ID) {
				sl.Entries[i] = nil
			}
		}
	}
}
