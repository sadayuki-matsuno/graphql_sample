package vuls

import (
	vulsModels "github.com/future-architect/vuls/models"
	"github.com/inconshreveable/log15"
	graphql "github.com/neelance/graphql-go"
)

// ScanResultSchema : ScanResultSchema
var ScanResultSchema = `
type ScanResult {
	ID         : String!
	OrganizationID : String!
	GroupID   : String!
	#	ScannedAt :Time!
	#	Lang      : String!
	#	ServerName :String!
	#	Family     :String!
	#	Release    :String!
	#	Errors     :[String!]!
}

input ScanResultInput {
	ScannedAt :Time!
	Lang      : String!
	ServerName :String!
	Family     :String!
	Release    :String!
	Errors     :[String!]!
}

type Container {
	ContainerID: String
    Name       : String
    Image      : String
    Type       : String
}

type Platform {
	Name      : String
	InstanceID: String
}

type VulnInfo {
	CveID            :String
	Confidence       :Confidence
	Packages         :[PackageInfo]
	DistroAdvisories :[DistroAdvisory]
	CpeNames         :[String]
}

type Confidence {
	Score           :Int
	DetectionMethod :String
}

type DistroAdvisory {
	AdvisoryID :String
	Severity   :String
	Issued     :Time
	Updated    :Time
}

type CveInfo {
	CveDetail : Cve
	VulnInfo: VulnInfo
}

type PackageInfo {
	Name       :String
	Version    :String
	Release    :String
	NewVersion :String
	NewRelease :String
	Repository :String
	Changelog  :Changelog
}

type Changelog {
	Contents :String
	Method   :String
}
`

func init() {
}

type scanResult struct {
	ID             string
	OrganizationID string
	GroupID        string
	vulsModels.ScanResult
}

type scanResultInput struct {
	OrganizationID string
	GroupID        string
	//	ScannedAt      time.Time
	ScannedAt graphql.Time

	Lang       string
	ServerName string // TOML Section key
	Family     string
	Release    string
	Container  vulsModels.Container
	Platform   vulsModels.Platform

	// Scanned Vulns via SSH + CPE Vulns
	ScannedCves []vulsModels.VulnInfo

	KnownCves   []vulsModels.CveInfo
	UnknownCves []vulsModels.CveInfo
	IgnoredCves []vulsModels.CveInfo

	Packages vulsModels.PackageInfoList
	Errors   []string
}

// ScanResultResolver : ScanResultResolver
type ScanResultResolver struct {
	s *scanResult
}

// ScanResults : scanResults
func (r *Resolver) ScanResults(
	args *struct {
		OrganizationID string
		GroupID        string
	}) []*ScanResultResolver {
	var t []*ScanResultResolver
	log15.Info("Stored Data", "method", "vuls.ScanResults", "args", args)
	return t
}

// AddScanResult : create scanResult
func (r *Resolver) AddScanResult(args *struct {
	OrganizationID  string
	GroupID         string
	ScanResultInput *scanResultInput
}) *ScanResultResolver {
	scanResult := &scanResult{
		OrganizationID: args.OrganizationID,
		GroupID:        args.GroupID,
	}
	//	log15.Info("Call CreateScanResult", "method", "vuls.CreateScanResult", "args", args)
	//
	log15.Info("Stored Data", "method", "vuls.CreateScanResult", "args", args)
	return &ScanResultResolver{scanResult}
	//return nil
}

// ID : ID
func (r *ScanResultResolver) ID() string {
	return r.s.ID
}

// OrganizationID : OrganizationID
func (r *ScanResultResolver) OrganizationID() string {
	return r.s.OrganizationID
}

// GroupID :GroupID
func (r *ScanResultResolver) GroupID() string {
	return r.s.GroupID
}

// // ScannedAt : ScannedAt
// func (r *ScanResultResolver) ScannedAt() string {
// 	return r.s.ScannedAt.String()
// }
//
// // Lang : Lang
// func (r *ScanResultResolver) Lang() string {
// 	return r.s.Lang
// }
//
// // ServerName : ServerName
// func (r *ScanResultResolver) ServerName() string {
// 	return r.s.ServerName
// }
//
// // Family : Family
// func (r *ScanResultResolver) Family() string {
// 	return r.s.Family
// }
//
// // Release : Release
// func (r *ScanResultResolver) Release() string {
// 	return r.s.Release
// }
//
// // Errors : Errors
// func (r *ScanResultResolver) Errors() []string {
// 	return r.s.Errors
// }
