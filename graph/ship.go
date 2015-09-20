package graph

import (
	"github.com/tallstreet/graphql"
	"github.com/tallstreet/graphql/executor/resolver"
	"github.com/tallstreet/graphql/schema"
	"golang.org/x/net/context"
)

type Ship struct {
	Name string
}

func (ship Ship) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "Ship",
		Description: "A Star Wars Ship",
		Fields: schema.GraphQLFieldSpecMap{
			"name": {
				Name:        "name",
				Description: "The name of ship.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, ship.Name, f)
				},
			},
		},
	}
}
