package graph

import (
	"github.com/tallstreet/graphql"
	"github.com/tallstreet/graphql/executor/resolver"
	"github.com/tallstreet/graphql/schema"
	"golang.org/x/net/context"
)

type Faction struct {
	Id string
	Name string
	Ships []*Ship
}


func (faction Faction) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "Faction",
		Description: "A Star Wars Faction",
		Fields: schema.GraphQLFieldSpecMap{
			"id": {
				Name:        "id",
				Description: "The id of faction.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, faction.Id, f)
				},
			},
			"name": {
				Name:        "name",
				Description: "The name of faction.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					return r.Resolve(ctx, faction.Name, f)
				},
			},
			"ships": {
				Name:        "ships",
				Description: "The ships belonging with the faction.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					ships := make([]interface{}, 0, len(faction.Ships))
					for _, ship := range faction.Ships {
						s, err := r.Resolve(ctx, ship, f)
						if err != nil {
							return nil, err
						}
						ships = append(ships, s)
					}
					return r.Resolve(ctx, ships, f)
				},
			},
		},
	}
}
