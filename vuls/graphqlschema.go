package vuls

// Schema : schema
var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	# The query type, represents all of the entry points into our object graph
	type Query {
		server(id: ID!): Server
		cves(cveIds: [String!]!): [Cve]
		cveList(cveIds: [String!]!): CveList!
	}
	# The mutation type, represents all updates we can make to our data
	type Mutation {
		createTask(cveId: String!,userId : String!, task: TaskInput!): Task
	}

	type Server {
		id: ID!
		name: String!
	}
	type Cve {
		cveId: String!
		nvd: Nvd
		jvn: Jvn
	}

	type Nvd {
		id: Int!
		summary: String
		score: Float
		accessVector: String
		accessComplexity: String
		authentication: String
		confidentialityImpact: String
		integrityImpact: String
		availabilityImpact: String
		cpes : [Cpe]
		cweID: String
		references: [Reference]
		publishedDate: String
	    lastModifiedDate: String
	}

    // Jvn is a model of JVN
    type Jvn {
		id: Int!
		cveDetailID: Int
		title  : String
    	summary: String
    	jvnLink: String
    	jvnID  : String
    	score  :  Float
		severity: String
    	vector  : String
		cpes : [Cpe]
		publishedDate: String
		lastModifiedDate: String
    }

	type CveList {
		# The total number of friends
		totalCount: Int!
		# The edges for each of the character's friends.
		edges: [CveListEdge]
		# A list of the friends, as a convenience when edges are not needed.
		cves: [Cve]
		# Information for paginating this connection
		pageInfo: PageInfo!
	}

	# An edge object for a character's friends
	type CveListEdge {
		# A cursor used for pagination
		cursor: ID!
		# The character represented by this friendship edge
		node: Cve
	}

	type Cpe {
		jvnID		 	:	Int
		nvdID 		 	:	Int
    	cpeName		 	:	String
    	part         	:	String
    	vendor       	:	String
    	product      	:	String
    	version      	:	String
    	vendorUpdate 	:	String
    	edition      	:	String
    	language     	:	String
	}

	type Reference {
		jvnID		: Int
		nvdID 		: Int
		source		: String
		link  		: String
	}

	type Task {
		serverName: String!
		organizationID: String!
		comments: String
	}

	input TaskInput {
		serverName: String!
		organizationID: String!
		comments: String
	}

	# Information for paginating this connection
	type PageInfo {
		startCursor: ID
		endCursor: ID
		hasNextPage: Boolean!
	}
`

// Resolver : resolver
type Resolver struct{}
