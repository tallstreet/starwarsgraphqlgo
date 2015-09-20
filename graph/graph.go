package graph

import (
	"fmt"
	"github.com/tallstreet/graphql"
	"github.com/tallstreet/graphql/executor/resolver"
	"github.com/tallstreet/graphql/schema"
	"golang.org/x/net/context"
)

type Graph struct {
	Ships        map[string]*Ship
	Factions     map[string]*Faction
}

func NewGraph() *Graph {
	graph := &Graph{}
	graph.Ships = make(map[string]*Ship)
	graph.Factions = make(map[string]*Faction)
	
	graph.Ships["1"] = &Ship{
        Name: "X-Wing",
    }
	
	graph.Ships["2"] = &Ship{
        Name: "Y-Wing",
    }
	
	graph.Ships["3"] = &Ship{
        Name: "A-Wing",
    }
	
	graph.Ships["4"] = &Ship{
        Name: "Millenium Falcon",
    }
	
	graph.Ships["5"] = &Ship{
        Name: "Home One",
    }
	
	graph.Ships["6"] = &Ship{
        Name: "TIE Fighter",
    }
	
	graph.Ships["7"] = &Ship{
        Name: "TIE Interceptor",
    }
	
	graph.Ships["8"] = &Ship{
        Name: "Executor",
    }
	
	graph.Factions["1"] = &Faction{
        Name: "Alliance to Restore the Republic",
		Ships: []*Ship{
			graph.Ships["1"], 
			graph.Ships["2"], 
			graph.Ships["3"], 
			graph.Ships["4"], 
			graph.Ships["5"],
		},
    }
	
	graph.Factions["2"] = &Faction{
        Name: "Galactic Empire",
		Ships: []*Ship{
			graph.Ships["6"], 
			graph.Ships["7"], 
			graph.Ships["8"],
		},
    }
	

	return graph
}


func (graph Graph) GraphQLTypeInfo() schema.GraphQLTypeInfo {
	return schema.GraphQLTypeInfo{
		Name:        "Star Wars",
		Description: "A Star Wars Graph",
		Fields: schema.GraphQLFieldSpecMap{
			"ship": {
				Name:        "ship",
				Description: "Star Wars Ship",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					if o, ok := f.Arguments.Get("id"); ok {
						if idStr, ok := o.(*graphql.Value).Value.(string); ok {
							g := graph.Ships[idStr]
							
							if g != nil {
								return r.Resolve(ctx, g, f)
							}
							return nil, fmt.Errorf("Ship for %s not found", idStr)
						} else {
							return nil, fmt.Errorf("'id' argument should be a string. Got %#v", o)
						}
					}
					return nil, fmt.Errorf("Argument 'id' must be specified")
				},
				IsRoot: true,
			},
			"factions": {
				Name:        "factions",
				Description: "The name of faction.",
				Func: func(ctx context.Context, r resolver.Resolver, f *graphql.Field) (interface{}, error) {
					if o, ok := f.Arguments.Get("id"); ok {
						if idStr, ok := o.(*graphql.Value).Value.(string); ok {
							g := graph.Factions[idStr]
							
							if g != nil {
								return r.Resolve(ctx, g, f)
							}
							return nil, fmt.Errorf("Faction for %s not found", idStr)
						} else {
							return nil, fmt.Errorf("'id' argument should be a string. Got %#v", o)
						}
					}
					return nil, fmt.Errorf("Argument 'id' must be specified")
				},
				IsRoot: true,
			},
		},
	}
}
