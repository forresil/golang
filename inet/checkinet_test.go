package inet

import (
	"log"
	"testing"
)

func TestConnected(t *testing.T) {
	status := Connected()

	if status != true {
		t.Errorf("Internet connection is down!")
	}
	log.Println("Your Internet connection is GOOD!")

}
