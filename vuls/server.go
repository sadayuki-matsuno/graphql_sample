package vuls

import (
	graphql "github.com/neelance/graphql-go"
)

var serverData = make(map[graphql.ID]*server)

func init() {
	for _, s := range servers {
		serverData[s.ID] = s
	}
}

type server struct {
	ID   graphql.ID
	Name string
}

// type cve struct {
// 	CveID graphql.ID
// 	Name  string
// }

var servers = []*server{
	{
		ID:   "2000",
		Name: "C-3PO",
	},
	{
		ID:   "2001",
		Name: "R2-D2",
	},
}

// Server : Server
func (r *Resolver) Server(args *struct{ ID graphql.ID }) *ServerResolver {
	if s := serverData[args.ID]; s != nil {
		return &ServerResolver{s}
	}
	return nil
}

// ServerResolver :
type ServerResolver struct {
	s *server
}

// ID : id
func (r *ServerResolver) ID() graphql.ID {
	return r.s.ID
}

// Name : name
func (r *ServerResolver) Name() string {
	//	return r.s.Name
	return "Name"
}
