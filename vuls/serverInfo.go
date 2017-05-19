package vuls

// ServerInfo : server info
type serverInfo struct {
	Name      string
	Release   string
	Platform  Platform
	Container Container
	Packages  []Package
}

// Platform has platform information
type Platform struct {
	Name       string // aws or azure or gcp or other...
	InstanceID string
}

// Container has Container information
type Container struct {
	ContainerID string
	Name        string
	Image       string
	Type        string
}

// Package has installed packages.
type Package struct {
	Name        string
	Version     string
	Release     string
	NewVersion  string
	NewRelease  string
	Repository  string
	Changelog   Changelog
	NotFixedYet bool // Ubuntu OVAL Only
}

func (r *Resolver) ServerInfo(args *struct{ name string }) *serverInfoResolver {
	if s := serverInfoData[args.name]; s != nil {
		return &serverInfoResolver{s}
	}
	return nil
}

type serverInfoResolver struct {
	s *serverInfo
}

func (r *serverInfoResolver) Name() strinf {
	return r.s.Name
}

func (r *serverInfoResolver) Release() string {
	return r.s.Release
}

func (r *serverInfoResolver) Platform(Name string) *[]*platformResolver {
	return platformResolver(r.s.Platform.Name)
}

func (r *serverInfoResolver) Container(Name string) *[]*containerResolver {
	return containerResolver(r.s.Container.ContainerID)
}

func (r *serverInfoResolver) Package(Name string) *[]*packageResolver {
	return packageResolver(r.s.Package.Name)
}

var serverInfoResolver = []*serverInfoResolver{
	{
		Name:     "2000",
		Release:  "C-3PO",
		Platform: Package{Name: "2001"},
	},
	{
		Name:     "2001",
		Release:  "C-2PO",
		Platform: Package{Name: "2004"},
	},
}
