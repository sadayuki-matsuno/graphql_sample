package vuls

import "time"

// Schema : schema for graqhql
var Schema = `
	schema {
		query: Query
		mutation: Mutation
	}
	# The query type, represents all of the entry points into our object graph
	type Query {
	    serverInfo(id: ID!): ServerInfo
		# Name string
		# Platform Platform
        # Release string
		# Container Container
		# Packages []Package
        # Errors   []string
        # Optional [][]interface{}
		
		scanInfo(id: ID!): ScanInfo
		# ScannedAt time.Time
		# JSONVersion string
		# Lang string 
		# ServerInfo ServerInfo
		# vulnInfos []vulnInfo

		vulnInfo(cveID: String!): VulnInfo
        # CveID string
        # Confidence Confidence
        # Packages []Package
        # DistroAdvisories []DistroAdvisories
        # Cpes []Cpe
        # CveContents CveContents

		package(name: String!): Package
        # Name string
        # Version string
        # Release string
        # NewVersion string
        # NewRelease string
        # Repository string
        # Changelog Changelog
        # NotFixedYet bool

        platform(name: String): Platform
		# Name string
		# InstanceID string

		container(containerID: String): Container
		# ContainerID: string
		# Name string
		# Image string
		# Type string

		confidence(score: Int): Confidence
		# Score int
		# DetectionMethod string

		distroAdvisory(advisoryID: String): DistroAdvisory
		# AdvisoryID string
		# Severity string
		# Issued time.Time
		# Updated time.Time

		cveContent(cveID: String): CveContent
		# Type string
		# CveID string
		# title string
		# Summary string
        # Severity string
        # Cvss2Score float64
        # Cvss2Vector string
        # Cvss3Score float64
        # Cvss3Vector string
        # SourceLink string
        # Cpes []Cpe
        # References References
        # CweID string
        # Published  time.Time
        # LastModified time.Time
		
		reference(refID: String): Reference
		# RefID string
		# Source string
		# Link string

		cpe(name: String): Cpe
		# name string

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
	type Droid implements Character {
		# The ID of the droid
		id: ID!
		# What others call this droid
		name: String!
		# This droid's friends, or an empty list if they have none
		friends: [Character]
		# The friends of the droid exposed as a connection with edges
		friendsConnection(first: Int, after: ID): FriendsConnection!
		# The movies this droid appears in
		appearsIn: [Episode!]!
		# This droid's primary function
		primaryFunction: String
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
	union SearchResult = Human | Droid | Starship



    type ServerInfo {
		id: ID!
    	name: String
    	release: String
    	platform: Platform
    	container: Container
    	packages: [Package]
    }
    type ScanInfo {
		id: ID!
    	scannedAt  : String
    	jsonVersion: String
    	lang       : String
    	serverInfo : [ServerInfo]
    	vulnInfos  : [VulnInfo]
    	errors     : [String]
    	optional   : [String]
    }


    type Container {
		containerID: String
    	name       : String
    	image      : String
    	type       : String
    }
    
    
    type Platform {
		name      : String 
    	instanceID: String
    }
    
    type VulnInfo {
		cveID           : String
    	confidence      : Confidence
    	packageNames    : [String]
    	distroAdvisories: [DistroAdvisory]
    	cpeNames        : [String]
    	cveContents     : [CveContent]
    }
    
    
    
    type Confidence {
		score          : Int
    	detectionMethod: String
    }
    
    
    type DistroAdvisory {
		advisoryID: String
    	severity  : String
    	issued    : String
    	updated   : String
    }
    
    type CveContent {
		type        : String
    	cveID       : String
    	title       : String
    	summary     : String
    	severity    : String
    	cvss2Score  : Float
    	cvss2Vector : String
    	cvss3Score  : Float
    	cvss3Vector : String
    	sourceLink  : String
    	cpes        : [Cpe]
    	references  : [Reference]
    	cweID       : String
    	published   : String
    	lastModified: String
    }
    
    
    type Reference {
		source: String
    	link  : String
    	refID : String
    }
    
    
    type Cpe {
		cpeName: String
    }
    
    type Package {
		name       : String
    	version    : String
    	release    : String
    	newVersion : String
    	newRelease : String
    	repository : String
    	changelog  : Changelog
    	notFixedYet: Boolean
    }
    
    
    
    type Changelog {
		Contents: String
    	Method  : String
    }
`

// ScanInfo vuls sacn result
type ScanInfo struct {
	ScannedAt   time.Time
	JSONVersion string
	Lang        string
	ServerInfo  []ServerInfo
	VulnInfos   []VulnInfo
	Errors      []string
	Optional    []string
}

// VulnInfos is VulnInfo list, getter/setter, sortable methods.
type VulnInfos map[string]VulnInfo

// VulnInfo holds a vulnerability information and unsecure packages
type VulnInfo struct {
	CveID            string
	Confidence       Confidence
	PackageNames     []string
	DistroAdvisories []DistroAdvisory // for Aamazon, RHEL, FreeBSD
	CpeNames         []string
	CveContents      CveContents
}

// Confidence is a ranking how confident the CVE-ID was deteted correctly
// Score: 0 - 100
type Confidence struct {
	Score           int
	DetectionMethod string
}

// DistroAdvisory has Amazon Linux, RHEL, FreeBSD Security Advisory information.
type DistroAdvisory struct {
	AdvisoryID string
	Severity   string
	Issued     time.Time
	Updated    time.Time
}

// CveContents has CveContent
type CveContents map[CveContentType]CveContent

// CveContent has abstraction of various vulnerability information
type CveContent struct {
	Type         CveContentType
	CveID        string
	Title        string
	Summary      string
	Severity     string
	Cvss2Score   float64
	Cvss2Vector  string
	Cvss3Score   float64
	Cvss3Vector  string
	SourceLink   string
	Cpes         []Cpe
	References   References
	CweID        string
	Published    time.Time
	LastModified time.Time
}

// References is a slice of Reference
type References []Reference

// Reference has a related link of the CVE
type Reference struct {
	Source string
	Link   string
	RefID  string
}

// Cpe is Common Platform Enumeration
type Cpe struct {
	CpeName string
}

// CveContentType is a source of CVE information
type CveContentType string

// Changelog has contents of changelog and how to get it.
// Method: modesl.detectionMethodStr
type Changelog struct {
	Contents string
	Method   string
}

// Resolver : vuls resolver
type Resolver struct{}
