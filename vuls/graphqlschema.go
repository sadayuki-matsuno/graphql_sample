package vuls

import "github.com/inconshreveable/log15"

// Schema : schema
var Schema string

// CommonSchema : CommonSchema
var CommonSchema = `
schema {
	query: Query
	mutation: Mutation
}
# The query type, represents all of the entry points into our object graph
type Query {
	servers(id: ID!): Server
	cves(cveIds: [String!]!): [Cve]
	cveList(cveIds: [String!]!): CveList!
}
# The mutation type, represents all updates we can make to our data
type Mutation {
	createTask(cveId: String!,userId : String!, task: TaskInput!): Task
	addScanResult(organizationID: String!,groupID : String!, scanResultInput: ScanResultInput!): ScanResult
}
scalar Time

# Information for paginating this connection
type PageInfo {
	startCursor: ID
	endCursor: ID
	hasNextPage: Boolean!
}
`

func init() {

	Schema = CommonSchema + ScanResultSchema + CveSchema + TaskSchema + ServerSchema
	log15.Info("message", "method", ".", "Schema", Schema)
}

// Resolver : resolver
type Resolver struct{}
