package vuls

import (
	"github.com/inconshreveable/log15"
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

// Servers : Servers
func (r *Resolver) Servers(args *struct{ ID graphql.ID }) *ServerResolver {
	log15.Info("message", "method", "vuls.ID", "args.ID", args.ID)
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
	log15.Info("message", "method", "vuls.ID", "r.s.ID", r.s.ID)
	return r.s.ID
}

// Name : name
func (r *ServerResolver) Name() string {
	//	return r.s.Name
	return "Name"
}
