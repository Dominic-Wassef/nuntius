package subscription

import "nuntius/connection"

// Subscription A Channel Subscription
type Subscription struct {
	Connection *connection.Connection
	ID         string
	Data       string
}

// New Create a new Subscription
func New(conn *connection.Connection, data string) *Subscription {
	return &Subscription{Connection: conn, Data: data}
}
