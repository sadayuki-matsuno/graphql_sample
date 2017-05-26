// Package starwars provides a example schema and resolver based on Star Wars characters.
//
// Source: https://github.com/graphql/graphql.github.io/blob/source/site/_core/swapiSchema.js
package starwars

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	graphql "github.com/neelance/graphql-go"
)

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
		droid(id: ID!): Droid
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
`

type human struct {
	ID        graphql.ID
	Name      string
	Friends   []graphql.ID
	AppearsIn []string
	Height    float64
	Mass      int
	Starships []graphql.ID
}

var humans = []*human{
	{
		ID:        "1000",
		Name:      "Luke Skywalker",
		Friends:   []graphql.ID{"1002", "1003", "2000", "2001"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    1.72,
		Mass:      77,
		Starships: []graphql.ID{"3001", "3003"},
	},
	{
		ID:        "1001",
		Name:      "Darth Vader",
		Friends:   []graphql.ID{"1004"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    2.02,
		Mass:      136,
		Starships: []graphql.ID{"3002"},
	},
	{
		ID:        "1002",
		Name:      "Han Solo",
		Friends:   []graphql.ID{"1000", "1003", "2001"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    1.8,
		Mass:      80,
		Starships: []graphql.ID{"3000", "3003"},
	},
	{
		ID:        "1003",
		Name:      "Leia Organa",
		Friends:   []graphql.ID{"1000", "1002", "2000", "2001"},
		AppearsIn: []string{"NEWHOPE", "EMPIRE", "JEDI"},
		Height:    1.5,
		Mass:      49,
	},
	{
		ID:        "1004",
		Name:      "Wilhuff Tarkin",
		Friends:   []graphql.ID{"1001"},
		AppearsIn: []string{"NEWHOPE"},
		Height:    1.8,
		Mass:      0,
	},
}

var humanData = make(map[graphql.ID]*human)

func init() {
	for _, h := range humans {
		humanData[h.ID] = h
	}
}

type droid struct {
	ID              graphql.ID
	Name            string
	Friends         []graphql.ID
	AppearsIn       []string
	PrimaryFunction string
}

var droids = []*droid{
	{
		ID:              "2000",
		Name:            "C-3PO",
		Friends:         []graphql.ID{"1000", "1002", "1003", "2001"},
		AppearsIn:       []string{"NEWHOPE", "EMPIRE", "JEDI"},
		PrimaryFunction: "Protocol",
	},
	{
		ID:              "2001",
		Name:            "R2-D2",
		Friends:         []graphql.ID{"1000", "1002", "1003"},
		AppearsIn:       []string{"NEWHOPE", "EMPIRE", "JEDI"},
		PrimaryFunction: "Astromech",
	},
}

var droidData = make(map[graphql.ID]*droid)

func init() {
	for _, d := range droids {
		droidData[d.ID] = d
	}
}

type starship struct {
	ID     graphql.ID
	Name   string
	Length float64
}

var starships = []*starship{
	{
		ID:     "3000",
		Name:   "Millennium Falcon",
		Length: 34.37,
	},
	{
		ID:     "3001",
		Name:   "X-Wing",
		Length: 12.5,
	},
	{
		ID:     "3002",
		Name:   "TIE Advanced x1",
		Length: 9.2,
	},
	{
		ID:     "3003",
		Name:   "Imperial shuttle",
		Length: 20,
	},
}

var starshipData = make(map[graphql.ID]*starship)

func init() {
	for _, s := range starships {
		starshipData[s.ID] = s
	}
}

type review struct {
	stars      int32
	commentary *string
}

var reviews = make(map[string][]*review)

// Resolver : resolver
type Resolver struct{}

// Hero : hero
func (r *Resolver) Hero(args *struct{ Episode string }) *CharacterResolver {
	if args.Episode == "EMPIRE" {
		return &CharacterResolver{&HumanResolver{humanData["1000"]}}
	}
	return &CharacterResolver{&DroidResolver{droidData["2001"]}}
}

// Reviews : reviews
func (r *Resolver) Reviews(args *struct{ Episode string }) []*ReviewResolver {
	var l []*ReviewResolver
	for _, review := range reviews[args.Episode] {
		l = append(l, &ReviewResolver{review})
	}
	return l
}

// Search : search
func (r *Resolver) Search(args *struct{ Text string }) []*SearchResultResolver {
	var l []*SearchResultResolver
	for _, h := range humans {
		if strings.Contains(h.Name, args.Text) {
			l = append(l, &SearchResultResolver{&HumanResolver{h}})
		}
	}
	for _, d := range droids {
		if strings.Contains(d.Name, args.Text) {
			l = append(l, &SearchResultResolver{&DroidResolver{d}})
		}
	}
	for _, s := range starships {
		if strings.Contains(s.Name, args.Text) {
			l = append(l, &SearchResultResolver{&StarshipResolver{s}})
		}
	}
	return l
}

// Character : character
func (r *Resolver) Character(args *struct{ ID graphql.ID }) *CharacterResolver {
	if h := humanData[args.ID]; h != nil {
		return &CharacterResolver{&HumanResolver{h}}
	}
	if d := droidData[args.ID]; d != nil {
		return &CharacterResolver{&DroidResolver{d}}
	}
	return nil
}

// Human : human
func (r *Resolver) Human(args *struct{ ID graphql.ID }) *HumanResolver {
	if h := humanData[args.ID]; h != nil {
		return &HumanResolver{h}
	}
	return nil
}

// Droid : droid
func (r *Resolver) Droid(args *struct{ ID graphql.ID }) *DroidResolver {
	if d := droidData[args.ID]; d != nil {
		return &DroidResolver{d}
	}
	return nil
}

// Starship : starship
func (r *Resolver) Starship(args *struct{ ID graphql.ID }) *StarshipResolver {
	if s := starshipData[args.ID]; s != nil {
		return &StarshipResolver{s}
	}
	return nil
}

// CreateReview : create review
func (r *Resolver) CreateReview(args *struct {
	Episode string
	Review  *reviewInput
}) *ReviewResolver {
	review := &review{
		stars:      args.Review.Stars,
		commentary: args.Review.Commentary,
	}
	reviews[args.Episode] = append(reviews[args.Episode], review)
	return &ReviewResolver{review}
}

type friendsConenctionArgs struct {
	First *int32
	After *graphql.ID
}

type character interface {
	ID() graphql.ID
	Name() string
	Friends() *[]*CharacterResolver
	FriendsConnection(*friendsConenctionArgs) (*FriendsConnectionResolver, error)
	AppearsIn() []string
}

// CharacterResolver : characterResolver
type CharacterResolver struct {
	character
}

// ToHuman : to human
func (r *CharacterResolver) ToHuman() (*HumanResolver, bool) {
	c, ok := r.character.(*HumanResolver)
	return c, ok
}

// ToDroid : to droid
func (r *CharacterResolver) ToDroid() (*DroidResolver, bool) {
	c, ok := r.character.(*DroidResolver)
	return c, ok
}

// HumanResolver : HumanResolver
type HumanResolver struct {
	h *human
}

// ID : id
func (r *HumanResolver) ID() graphql.ID {
	return r.h.ID
}

// Name : name
func (r *HumanResolver) Name() string {
	return r.h.Name
}

// Height : height
func (r *HumanResolver) Height(args *struct{ Unit string }) float64 {
	return convertLength(r.h.Height, args.Unit)
}

// Mass : mass
func (r *HumanResolver) Mass() *float64 {
	if r.h.Mass == 0 {
		return nil
	}
	f := float64(r.h.Mass)
	return &f
}

// Friends : friends
func (r *HumanResolver) Friends() *[]*CharacterResolver {
	return resolveCharacters(r.h.Friends)
}

// FriendsConnection : friendsConnection
func (r *HumanResolver) FriendsConnection(args *friendsConenctionArgs) (*FriendsConnectionResolver, error) {
	return newFriendsConnectionResolver(r.h.Friends, args)
}

// AppearsIn : AppearsIn
func (r *HumanResolver) AppearsIn() []string {
	return r.h.AppearsIn
}

// Starships : starships
func (r *HumanResolver) Starships() *[]*StarshipResolver {
	l := make([]*StarshipResolver, len(r.h.Starships))
	for i, id := range r.h.Starships {
		l[i] = &StarshipResolver{starshipData[id]}
	}
	return &l
}

// DroidResolver : DroidResolver
type DroidResolver struct {
	d *droid
}

// ID : id
func (r *DroidResolver) ID() graphql.ID {
	return r.d.ID
}

// Name : name
func (r *DroidResolver) Name() string {
	return r.d.Name
}

// Friends : friends
func (r *DroidResolver) Friends() *[]*CharacterResolver {
	return resolveCharacters(r.d.Friends)
}

// FriendsConnection : friendsConnection
func (r *DroidResolver) FriendsConnection(args *friendsConenctionArgs) (*FriendsConnectionResolver, error) {
	return newFriendsConnectionResolver(r.d.Friends, args)
}

// AppearsIn : appearsIn
func (r *DroidResolver) AppearsIn() []string {
	return r.d.AppearsIn
}

// PrimaryFunction : primaryFunction
func (r *DroidResolver) PrimaryFunction() *string {
	if r.d.PrimaryFunction == "" {
		return nil
	}
	return &r.d.PrimaryFunction
}

// StarshipResolver : StarshipResolver
type StarshipResolver struct {
	s *starship
}

// ID : id
func (r *StarshipResolver) ID() graphql.ID {
	return r.s.ID
}

// Name : name
func (r *StarshipResolver) Name() string {
	return r.s.Name
}

// Length : length
func (r *StarshipResolver) Length(args *struct{ Unit string }) float64 {
	return convertLength(r.s.Length, args.Unit)
}

// SearchResultResolver : SearchResultResolver
type SearchResultResolver struct {
	result interface{}
}

// ToHuman : toHuman
func (r *SearchResultResolver) ToHuman() (*HumanResolver, bool) {
	res, ok := r.result.(*HumanResolver)
	return res, ok
}

// ToDroid : toDroid
func (r *SearchResultResolver) ToDroid() (*DroidResolver, bool) {
	res, ok := r.result.(*DroidResolver)
	return res, ok
}

// ToStarship : toStarship
func (r *SearchResultResolver) ToStarship() (*StarshipResolver, bool) {
	res, ok := r.result.(*StarshipResolver)
	return res, ok
}

func convertLength(meters float64, unit string) float64 {
	switch unit {
	case "METER":
		return meters
	case "FOOT":
		return meters * 3.28084
	default:
		panic("invalid unit")
	}
}

func resolveCharacters(ids []graphql.ID) *[]*CharacterResolver {
	var characters []*CharacterResolver
	for _, id := range ids {
		if c := resolveCharacter(id); c != nil {
			characters = append(characters, c)
		}
	}
	return &characters
}

func resolveCharacter(id graphql.ID) *CharacterResolver {
	if h, ok := humanData[id]; ok {
		return &CharacterResolver{&HumanResolver{h}}
	}
	if d, ok := droidData[id]; ok {
		return &CharacterResolver{&DroidResolver{d}}
	}
	return nil
}

// ReviewResolver : ReviewResolver
type ReviewResolver struct {
	r *review
}

// Stars : stars
func (r *ReviewResolver) Stars() int32 {
	return r.r.stars
}

// Commentary : commentary
func (r *ReviewResolver) Commentary() *string {
	return r.r.commentary
}

// FriendsConnectionResolver : FriendsConnectionResolver
type FriendsConnectionResolver struct {
	ids  []graphql.ID
	from int
	to   int
}

func newFriendsConnectionResolver(ids []graphql.ID, args *friendsConenctionArgs) (*FriendsConnectionResolver, error) {
	from := 0
	if args.After != nil {
		b, err := base64.StdEncoding.DecodeString(string(*args.After))
		if err != nil {
			return nil, err
		}
		i, err := strconv.Atoi(strings.TrimPrefix(string(b), "cursor"))
		if err != nil {
			return nil, err
		}
		from = i
	}

	to := len(ids)
	if args.First != nil {
		to = from + int(*args.First)
		if to > len(ids) {
			to = len(ids)
		}
	}

	return &FriendsConnectionResolver{
		ids:  ids,
		from: from,
		to:   to,
	}, nil
}

// TotalCount : totalCount
func (r *FriendsConnectionResolver) TotalCount() int32 {
	return int32(len(r.ids))
}

// Edges : edges
func (r *FriendsConnectionResolver) Edges() *[]*FriendsEdgeResolver {
	l := make([]*FriendsEdgeResolver, r.to-r.from)
	for i := range l {
		l[i] = &FriendsEdgeResolver{
			cursor: encodeCursor(r.from + i),
			id:     r.ids[r.from+i],
		}
	}
	return &l
}

// Friends : friends
func (r *FriendsConnectionResolver) Friends() *[]*CharacterResolver {
	return resolveCharacters(r.ids[r.from:r.to])
}

// PageInfo : pageInfo
func (r *FriendsConnectionResolver) PageInfo() *PageInfoResolver {
	return &PageInfoResolver{
		startCursor: encodeCursor(r.from),
		endCursor:   encodeCursor(r.to - 1),
		hasNextPage: r.to < len(r.ids),
	}
}

func encodeCursor(i int) graphql.ID {
	return graphql.ID(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i+1))))
}

// FriendsEdgeResolver : FriendsEdgeResolver
type FriendsEdgeResolver struct {
	cursor graphql.ID
	id     graphql.ID
}

// Cursor : cursor
func (r *FriendsEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

// Node : node
func (r *FriendsEdgeResolver) Node() *CharacterResolver {
	return resolveCharacter(r.id)
}

// PageInfoResolver : PageInfoResolver
type PageInfoResolver struct {
	startCursor graphql.ID
	endCursor   graphql.ID
	hasNextPage bool
}

// StartCursor : startCursor
func (r *PageInfoResolver) StartCursor() *graphql.ID {
	return &r.startCursor
}

// EndCursor : endCursor
func (r *PageInfoResolver) EndCursor() *graphql.ID {
	return &r.endCursor
}

// HasNextPage : hasNextPage
func (r *PageInfoResolver) HasNextPage() bool {
	return r.hasNextPage
}

type reviewInput struct {
	Stars      int32
	Commentary *string
}
