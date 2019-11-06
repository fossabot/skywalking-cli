package client

import (
	"context"
	"fmt"
	"github.com/machinebox/graphql"
)

func Services(duration Duration) []Service {
	ctx := context.Background()
	client := graphql.NewClient("http://122.112.182.72:8080/graphql")
	var services map[string][]Service
	request := graphql.NewRequest(`
		query ($duration: Duration!) {
			services: getAllServices(duration: $duration) {
				id
			 	name
			}
		}
	`)
	request.Var("duration", duration)
	if err := client.Run(ctx, request, &services); err != nil {
		fmt.Printf("%v\n", err)
		panic(err)
	}
	return services["services"]
}
