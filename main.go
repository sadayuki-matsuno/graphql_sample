package main

// go:generate sqlboiler postgres

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"./starwars"
	"./vuls"
	"github.com/inconshreveable/log15"
	_ "github.com/lib/pq"
	"github.com/neelance/graphql-go"
	"github.com/neelance/graphql-go/relay"
	"github.com/spf13/viper"
	"github.com/vattle/sqlboiler/boil"
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

	if err := prepareViper(); err != nil {
		log15.Warn("Failed to read viper config file", "method", "main.main", "err", err)
	}
	var dsn = fmt.Sprintf(
		"host=%s user=%s dbname=%s sslmode=%s password=%s",
		viper.GetString("postgres.host"),
		viper.GetString("postgres.user"),
		viper.GetString("postgres.dbname"),
		viper.GetString("postgres.sslmode"),
		viper.GetString("postgres.pass"),
	)
	log15.Info("check dsn", "method", "main.main", "dsn", dsn)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log15.Error("Failed to open database", "err", err)

	}
	boil.SetDB(db)
	log15.Info("connectd to db", "dsn", dsn)

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(page)
	}))

	//	http.Handle("/query", &relay.Handler{Schema: schema})
	http.Handle("/query", &relay.Handler{Schema: schema2})

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func prepareViper() error {
	log15.Info("call prepareViper", "method", "main.prepareViper")
	viper.SetConfigName("sqlboiler")

	configHome := os.Getenv("XDG_CONFIG_HOME")
	homePath := os.Getenv("HOME")
	wd, err := os.Getwd()
	if err != nil {
		wd = "./"
	}

	configPaths := []string{wd}
	if len(configHome) > 0 {
		configPaths = append(configPaths, filepath.Join(configHome, "sqlboiler"))
	} else {
		configPaths = append(configPaths, filepath.Join(homePath, ".config/sqlboiler"))
	}

	for _, p := range configPaths {
		viper.AddConfigPath(p)
	}

	err = viper.ReadInConfig()
	return err
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
