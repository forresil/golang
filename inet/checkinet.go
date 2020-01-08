package inet

import (
	"net/http"
)

// Connected - Tells you if there is an internet connection or not.
// Returns true if connected, false if not.
//------------------------------------------------------------------
func Connected() (ok bool) {
	_, err := http.Get("http://clients3.google.com/generate_204")
	if err != nil {
		return false
	}
	return true
}
