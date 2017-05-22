package vuls

import graphql "github.com/neelance/graphql-go"

// ServerInfo : server info
type serverInfo struct {
	ID      graphql.ID
	Name    string
	Release string
	//	Platform platform
}

// Platform has platform information
//type platform struct {
//	Name       string // aws or azure or gcp or other...
//	InstanceID string
//}

// ServerInfoResolver : aa
//type ServerInfoResolver interface {
//	ServerInfo() *serverInfoResolver
//}

// ServerInfoResolver : aa
type ServerInfoResolver struct {
	s *serverInfo
}

// PlatformResolver : aa
type PlatformResolver interface {
	Platform() *PlatformResolver
}

var serverInfoData = make(map[graphql.ID]*serverInfo)

func init() {
	for _, s := range serverInfos {
		serverInfoData[s.ID] = s
	}
}

//// ServerInfo Server Info
//func (r *Resolver) ServerInfo(args *struct{ ID graphql.ID }) *serverInfoResolver {
//	if d := serverInfoData[args.ID]; d != nil {
//		return &serverInfoResolver{d}
//	}
//	return nil
//}

// ServerInfo : server info
func (r *Resolver) ServerInfo(args *struct{ ID graphql.ID }) *ServerInfoResolver {
	if s := serverInfoData[args.ID]; s != nil {
		return &ServerInfoResolver{s}
	}
	return nil
}

// Name : aa
func (r *ServerInfoResolver) Name() string {
	return r.s.Name
}

// Release : aa
func (r *ServerInfoResolver) Release() string {
	return r.s.Release
}

var serverInfos = []*serverInfo{
	{
		ID:      "2000",
		Name:    "2000",
		Release: "C-3PO",
	},
	{
		ID:      "2001",
		Name:    "2001",
		Release: "C-2PO",
	},
}
