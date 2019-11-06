package client

import (
	"github.com/graphql-go/graphql",
)

func main() {
	graphql.Do(graphql.Params{
		Schema:         starwars.Schema,
		RequestString:  query,
		VariableValues: params,
	})
}
