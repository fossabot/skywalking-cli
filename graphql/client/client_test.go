package client

import (
	"fmt"
	"testing"
)

func TestClient(t *testing.T) {
	services := Services(Duration{
		Start: "2019-11-03 2326",
		End:   "2019-11-03 2341",
		Step:  "MINUTE",
	})

	fmt.Printf("%v", services)
}
