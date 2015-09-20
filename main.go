// Program basic_graphql_server shows a simple HTTP server that exposes a bare schema.
//
// Example:
//  $ go get github.com/tallstreet/graphql/example/basic_graphql_server
//  $ basic_graphql_server &
//  $ curl -g -XPOST 'http://localhost:8080/' --data-binary '{ "query": "{__schema{types{name,description,fields{name}}}}"}'
//  $ curl -g -XPOST 'http://localhost:8080/' --data-binary '{ "query": "query faction(id: \"2\"){ id, name } "}'
// {
//   "data": {
//     "faction": {
//       "Name": "Alliance to Restore the Republic",
//       "Ships": [
//         {
//           "Name": "X-Wing"
//         },
//         {
//           "Name": "Y-Wing"
//         },
//         {
//           "Name": "A-Wing"
//         },
//         {
//           "Name": "Millenium Falcon"
//         },
//         {
//           "Name": "Home One"
//         }
//       ]
//     }
//   }
// }
// Here we see the server showing the available root fields ("schema").

package main

import (
	"flag"
	"github.com/tallstreet/starwarsgraphqlgo/server"
)


var listenAddr = flag.String("l", ":8080", "listen addr")

func main() {
	flag.Parse()
	server.Application = server.NewApp(*listenAddr)
	server.Application.RunServer()
}