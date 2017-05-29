package vuls

// Schema : schema
var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	# The query type, represents all of the entry points into our object graph
	type Query {
		hero(episode: Episode = NEWHOPE): Character
		reviews(episode: Episode!): [Review]!
		search(text: String!): [SearchResult]!
		character(id: ID!): Character
		server(id: ID!): Server
		cves(cveIds: [String!]!): [Cve]
		cveList(cveIds: [String!]!): CveList!
		human(id: ID!): Human
		starship(id: ID!): Starship
	}
	# The mutation type, represents all updates we can make to our data
	type Mutation {
		createReview(episode: Episode!, review: ReviewInput!): Review
	}
	# The episodes in the Star Wars trilogy
	enum Episode {
		# Star Wars Episode IV: A New Hope, released in 1977.
		NEWHOPE
		# Star Wars Episode V: The Empire Strikes Back, released in 1980.
		EMPIRE
		# Star Wars Episode VI: Return of the Jedi, released in 1983.
		JEDI
	}
	# A character from the Star Wars universe
	interface Character {
		# The ID of the character
		id: ID!
		# The name of the character
		name: String!
		# The friends of the character, or an empty list if they have none
		friends: [Character]
		# The friends of the character exposed as a connection with edges
		friendsConnection(first: Int, after: ID): FriendsConnection!
		# The movies this character appears in
		appearsIn: [Episode!]!
	}
	# Units of height
	enum LengthUnit {
		# The standard unit around the world
		METER
		# Primarily used in the United States
		FOOT
	}
	# A humanoid creature from the Star Wars universe
	type Human implements Character {
		# The ID of the human
		id: ID!
		# What this human calls themselves
		name: String!
		# Height in the preferred unit, default is meters
		height(unit: LengthUnit = METER): Float!
		# Mass in kilograms, or null if unknown
		mass: Float
		# This human's friends, or an empty list if they have none
		friends: [Character]
		# The friends of the human exposed as a connection with edges
		friendsConnection(first: Int, after: ID): FriendsConnection!
		# The movies this human appears in
		appearsIn: [Episode!]!
		# A list of starships this person has piloted, or an empty list if none
		starships: [Starship]
	}
	# An autonomous mechanical character in the Star Wars universe
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

	# A connection object for a character's friends
	type FriendsConnection {
		# The total number of friends
		totalCount: Int!
		# The edges for each of the character's friends.
		edges: [FriendsEdge]
		# A list of the friends, as a convenience when edges are not needed.
		friends: [Character]
		# Information for paginating this connection
		pageInfo: PageInfo!
	}
	# An edge object for a character's friends
	type FriendsEdge {
		# A cursor used for pagination
		cursor: ID!
		# The character represented by this friendship edge
		node: Character
	}
	# Information for paginating this connection
	type PageInfo {
		startCursor: ID
		endCursor: ID
		hasNextPage: Boolean!
	}
	# Represents a review for a movie
	type Review {
		# The number of stars this review gave, 1-5
		stars: Int!
		# Comment about the movie
		commentary: String
	}
	# The input object sent when someone is creating a new review
	input ReviewInput {
		# 0-5 stars
		stars: Int!
		# Comment about the movie, optional
		commentary: String
	}
	type Starship {
		# The ID of the starship
		id: ID!
		# The name of the starship
		name: String!
		# Length of the starship, along the longest axis
		length(unit: LengthUnit = METER): Float!
	}
	union SearchResult = Human  | Starship
`
