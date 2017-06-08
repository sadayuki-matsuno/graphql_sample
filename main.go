package main

// go:generate sqlboiler postgres

import (
	"log"
	"net/http"

	//	_ "github.com/lib/pq"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"github.com/sadayuki-matsuno/graphql_sample/db"
	"github.com/sadayuki-matsuno/graphql_sample/starwars"
	"github.com/sadayuki-matsuno/graphql_sample/vuls"
)

var schema *graphql.Schema
var schema2 *graphql.Schema
var configPath = ""

func init() {
	var err error
	schema, err = graphql.ParseSchema(starwars.Schema, &starwars.Resolver{})
	schema2, err = graphql.ParseSchema(vuls.Schema, &vuls.Resolver{})
	if err != nil {
		panic(err)
	}
}

func main() {

	db.OpenDBs()
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	//	http.Handle("/query", &relay.Handler{Schema: schema})
	http.Handle("/query", &relay.Handler{Schema: schema2})
	log.Fatal(http.ListenAndServe(":8888", nil))
}

var page = []byte(`
<!DOCTYPE html>
<html>
	<head>
		<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.7.8/graphiql.css" />
		<script src="https://cdnjs.cloudflare.com/ajax/libs/fetch/1.0.0/fetch.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.3.2/react.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/react/15.3.2/react-dom.min.js"></script>
		<script src="https://cdnjs.cloudflare.com/ajax/libs/graphiql/0.7.8/graphiql.js"></script>
	</head>
	<body style="width: 100%; height: 100%; margin: 0; overflow: hidden;">
		<div id="graphiql" style="height: 100vh;">Loading...</div>
		<script>
			function graphQLFetcher(graphQLParams) {
				graphQLParams.variables = graphQLParams.variables ? JSON.parse(graphQLParams.variables) : null;
				return fetch("/query", {
					method: "post",
					body: JSON.stringify(graphQLParams),
					credentials: "include",
				}).then(function (response) {
					return response.text();
				}).then(function (responseBody) {
					try {
						return JSON.parse(responseBody);
					} catch (error) {
						return responseBody;
					}
				});
			}

			ReactDOM.render(
				React.createElement(GraphiQL, {fetcher: graphQLFetcher}),
				document.getElementById("graphiql")
			);
		</script>
	</body>
</html>
`)
