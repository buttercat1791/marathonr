package event

import (
	"fmt"
)

func HandleEvent(msg []interface{}) (err error) {
	// Read event from the second array element.
	event, ok := msg[1].(Event)
	if !ok {
		return fmt.Errorf("expected event as second element of message, got %T", msg[1])
	}

	// Insert event into database.
	// TODO

	return nil
}
