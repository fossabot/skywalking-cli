package service

import (
	"encoding/json"
	"fmt"
	"github.com/apache/skywalking-cli/graphql/client"
)

func (s *service) showList() error {
	services := client.Services(client.Duration{
		Start: "2019-11-03 2326",
		End:   "2019-11-03 2341",
		Step:  "MINUTE",
	})

	if bytes, e := json.Marshal(services); e != nil {
		return e
	} else {
		fmt.Printf("%v\n", string(bytes))
	}
	return nil
}
