package connection

import (
	"nuntius/mocks"
	"testing"
)

func TestNewConnection(t *testing.T) {
	expectedSocketID := "socketID"
	expectedSocket := mocks.MockSocket{}

	c := New(expectedSocketID, expectedSocket)

	if c.SocketID != expectedSocketID {
		t.Errorf("c.SocketID == %s, wants %s", c.SocketID, expectedSocketID)
	}

	if c.Socket != expectedSocket {
		t.Errorf("c.Socket == %v, wants %v", c.Socket, expectedSocket)
	}

	if c.CreatedAt.IsZero() {
		t.Errorf("c.createdAt.IsZero() == %t, wants %t", c.CreatedAt.IsZero(), false)
	}
}
