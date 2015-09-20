# StarWars GraphQL Go

Program basic_graphql_server shows a simple HTTP server that exposes a bare schema.

Example:

	$ go get github.com/tallstreet/graphql/example/basic_graphql_server
	$ basic_graphql_server &
	$ curl -g -XPOST 'http://localhost:8080/' --data-binary '{ "query": "{__schema{types{name,description,fields{name}}}}"}'
	$ curl -g -XPOST 'http://localhost:8080/' --data-binary '{ "query": "query gary { factions(id: \"2\") {  name, ships { name }  } } "}'
	{
		"data": {
			"factions": {
			"name": "Galactic Empire",
			"ships": [
				{
				"name": "Executor"
				},
				{
				"name": "TIE Fighter"
				},
				{
				"name": "TIE Interceptor"
				}
			]
			}
		}
	}

Here we see the server showing the available root fields ("schema").