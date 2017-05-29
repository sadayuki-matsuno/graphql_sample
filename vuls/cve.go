package vuls

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"../models"
	"github.com/inconshreveable/log15"
	graphql "github.com/neelance/graphql-go"
	"github.com/vattle/sqlboiler/queries/qm"
)

var log = log15.New()

func init() {
}

// Cves : Cves
func (r *Resolver) Cves(args *struct{ CveIDs []string }) *[]*CveResolver {

	var cveResolvers []*CveResolver
	uniqueCveIDs := getUniqueSlice(args.CveIDs)
	log15.Info("Select data from cves", "method", "cve.Cves", "uniqueCveIDs", uniqueCveIDs)
	for _, id := range uniqueCveIDs {
		if c, err := models.CveDetailsG(qm.Where("cve_id=?", id)).One(); c != nil {
			cveResolvers = append(cveResolvers, &CveResolver{c})
		} else if err != nil {
			log15.Warn("Failed to select data", "CveID", id, "err", err)
		}
	}
	log15.Info("Selected data from cves", "method", "cve.Cves", "cveResolvers", cveResolvers)
	return &cveResolvers
}

//// Cve : Cve
//func (r *Resolver) Cve(args *struct{ CveID string }) *CveResolver {
//	if c, err := models.CveDetailsG(qm.Where("CveID=?", args.CveID)).One(); c != nil {
//		log15.Info("Selected data from cve", "method", "cve.Cve", "c", c)
//		return &CveResolver{c}
//	} else if err != nil {
//		log15.Warn("Failed to select data", "CveID", args.CveID)
//	}
//	return nil
//}

// CveList : cvesConnection
func (r *Resolver) CveList(args *CveListArgs) (*CveListResolver, error) {

	log15.Info("Select data from cvelist", "method", "cve.CveList", args.CveIDs, args.CveIDs)
	uniqueCveIDs := getUniqueSlice(args.CveIDs)
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

	to := len(uniqueCveIDs)
	if args.First != nil {
		to = from + int(*args.First)
		if to > len(uniqueCveIDs) {
			to = len(uniqueCveIDs)
		}
	}

	return &CveListResolver{
		cveIDs: uniqueCveIDs,
		from:   from,
		to:     to,
	}, nil
}

// CveListResolver : CveListResolver
type CveListResolver struct {
	cveIDs []string
	from   int
	to     int
}

// CveResolver :
type CveResolver struct {
	//	c *cve.CveDetail
	c *models.CveDetail
}

// NvdResolver :
type NvdResolver struct {
	n *models.NVD
}

// JvnResolver :
type JvnResolver struct {
	j *models.JVN
}

// CpeResolver : CpeResolver
type CpeResolver struct {
	c *models.Cpe
}

// ReferenceResolver : ReferenceResolver
type ReferenceResolver struct {
	r *models.Reference
}

// CveListEdgeResolver : CveListEdgeResolver
type CveListEdgeResolver struct {
	cursor graphql.ID
	cveID  string
}

// CveListArgs : CveListArgs
type CveListArgs struct {
	CveIDs []string
	First  *int32
	After  *graphql.ID
}

// ListArgs : ListArgs
type ListArgs struct {
	First *int32
	After *graphql.ID
}

// CveID : id
func (r *CveResolver) CveID() string {
	return r.c.CveID.String
}

// Nvd : Nvd
func (r *CveResolver) Nvd() *NvdResolver {

	id := &r.c.ID
	log15.Info("Select data from nvd", "method", "cve.Nvd", "cve_detail_id", id)
	if n, err := models.NVDSG(qm.Where("cve_detail_id=?", id)).One(); n != nil {
		return &NvdResolver{n}
	} else if err != nil {
		log15.Warn("Failed to select data", "method", "cve.Nvd", "err", err)
	}
	return nil

}

// Jvn : Jvn
func (r *CveResolver) Jvn() *JvnResolver {
	id := &r.c.ID
	log15.Info("Select data from jvn", "method", "cve.Jvn", "cve_detail_id", id)
	if j, err := models.JVNSG(qm.Where("cve_detail_id=?", id)).One(); j != nil {
		return &JvnResolver{j}
	} else if err != nil {
		log15.Warn("Failed to select data", "method", "cve.Jvn", "err", err)
	}
	return nil
}

// ID : Summary
func (r *NvdResolver) ID() int32 {
	return int32(r.n.ID)
}

// Summary : Summary
func (r *NvdResolver) Summary() *string {
	return &r.n.Summary.String
}

// Score :  Score
func (r *NvdResolver) Score() *float64 {
	return &r.n.Score.Float64
}

// AccessVector : AccessVector
func (r *NvdResolver) AccessVector() *string {
	return &r.n.AccessVector.String
}

// AccessComplexity : AccessComplexity
func (r *NvdResolver) AccessComplexity() *string {
	return &r.n.AccessComplexity.String
}

// Authentication : Authentication
func (r *NvdResolver) Authentication() *string {
	return &r.n.Authentication.String
}

// ConfidentialityImpact : ConfidentialityImpact
func (r *NvdResolver) ConfidentialityImpact() *string {
	return &r.n.ConfidentialityImpact.String
}

// IntegrityImpact : IntegrityImpact
func (r *NvdResolver) IntegrityImpact() *string {
	return &r.n.IntegrityImpact.String
}

// AvailabilityImpact : AvailabilityImpact
func (r *NvdResolver) AvailabilityImpact() *string {
	return &r.n.AvailabilityImpact.String
}

// Cpes  : Nvd
func (r *NvdResolver) Cpes() *[]*CpeResolver {

	nvdID := &r.n.ID
	log15.Info("Select data from cpe", "method", "cve.Cpe", "nvd_id", nvdID)
	if cpes, err := models.CpesG(qm.Where("nvd_id=?", nvdID)).All(); cpes != nil {
		log15.Info("select multi data", "cpes", cpes)
		var cpeResolvers []*CpeResolver
		for _, c := range cpes {
			cpeResolvers = append(cpeResolvers, &CpeResolver{c})
		}
		return &cpeResolvers
	} else if err != nil {
		log15.Warn("Failed to select data", "method", "cve.Cpe", "err", err)
	}
	return nil

}

// CweID : CweID
func (r *NvdResolver) CweID() *string {
	return &r.n.CweID.String
}

// References  : Nvd
func (r *NvdResolver) References() *[]*ReferenceResolver {

	nvdID := &r.n.ID
	log15.Info("Select data from cpe", "method", "cve.Reference", "nvd_id", nvdID)
	if cpes, err := models.ReferencesG(qm.Where("nvd_id=?", nvdID)).All(); cpes != nil {
		log15.Info("select multi data", "cpes", cpes)
		var cpeResolvers []*ReferenceResolver
		for _, c := range cpes {
			cpeResolvers = append(cpeResolvers, &ReferenceResolver{c})
		}
		return &cpeResolvers
	} else if err != nil {
		log15.Warn("Failed to select data", "method", "cve.Reference", "err", err)
	}
	return nil

}

// PublishedDate : PublishedDate
func (r *NvdResolver) PublishedDate() *string {
	t := r.n.PublishedDate.Time.String()
	return &t
}

// LastModifiedDate : LastModifiedDate
func (r *NvdResolver) LastModifiedDate() *string {
	t := r.n.LastModifiedDate.Time.String()
	return &t
}

// ID : Summary
func (r *JvnResolver) ID() int32 {
	return int32(r.j.ID)
}

// CveDetailID : Cvedetailid
func (r *JvnResolver) CveDetailID() *int32 {
	id := int32(r.j.CveDetailID.Int)
	return &id
}

// Title       : Title
func (r *JvnResolver) Title() *string {
	return &r.j.Title.String
}

// Summary     : Summary
func (r *JvnResolver) Summary() *string {
	return &r.j.Summary.String
}

// JVNLink : Jvnlink
func (r *JvnResolver) JVNLink() *string {
	return &r.j.JVNLink.String
}

// JvnID       : Jvnid
func (r *JvnResolver) JvnID() *string {
	return &r.j.JVNID.String
}

// Score       : Score
func (r *JvnResolver) Score() *float64 {
	return &r.j.Score.Float64
}

// Severity    : Severity
func (r *JvnResolver) Severity() *string {
	return &r.j.Severity.String
}

// Vector      : Vector
func (r *JvnResolver) Vector() *string {
	return &r.j.Vector.String
}

// Cpes  : Cpe
func (r *JvnResolver) Cpes() *[]*CpeResolver {

	jvnID := &r.j.ID
	log15.Info("Select data from cpe", "method", "cve.Cpe", "jvn_id", jvnID)
	if cpes, err := models.CpesG(qm.Where("jvn_id=?", jvnID)).All(); cpes != nil {
		log15.Info("select multi data", "cpes", cpes)
		var cpeResolvers []*CpeResolver
		for _, c := range cpes {
			cpeResolvers = append(cpeResolvers, &CpeResolver{c})
		}
		return &cpeResolvers
	} else if err != nil {
		log15.Warn("Failed to select data", "method", "cve.Cpe", "err", err)
	}
	return nil

}

// References  : Jvn
func (r *JvnResolver) References() *[]*ReferenceResolver {

	jvnID := &r.j.ID
	log15.Info("Select data from cpe", "method", "cve.Reference", "jvn_id", jvnID)
	if cpes, err := models.ReferencesG(qm.Where("jvn_id=?", jvnID)).All(); cpes != nil {
		log15.Info("select multi data", "cpes", cpes)
		var cpeResolvers []*ReferenceResolver
		for _, c := range cpes {
			cpeResolvers = append(cpeResolvers, &ReferenceResolver{c})
		}
		return &cpeResolvers
	} else if err != nil {
		log15.Warn("Failed to select data", "method", "cve.Reference", "err", err)
	}
	return nil

}

// PublishedDate : PublishedDate
func (r *JvnResolver) PublishedDate() *string {
	t := r.j.PublishedDate.Time.String()
	return &t
}

// LastModifiedDate : LastModifiedDate
func (r *JvnResolver) LastModifiedDate() *string {
	t := r.j.LastModifiedDate.Time.String()
	return &t
}

// JvnID : JVNID
func (r *CpeResolver) JvnID() *int32 {
	id := int32(r.c.JVNID.Int)
	return &id
}

// NvdID : NVDID
func (r *CpeResolver) NvdID() *int32 {
	id := int32(r.c.NVDID.Int)
	return &id
}

// CpeName : CpeName
func (r *CpeResolver) CpeName() *string {
	return &r.c.CpeName.String
}

// Part : Part
func (r *CpeResolver) Part() *string {
	return &r.c.Part.String
}

// Vendor : Vendor
func (r *CpeResolver) Vendor() *string {
	return &r.c.Vendor.String
}

// Product : Product
func (r *CpeResolver) Product() *string {
	return &r.c.Product.String
}

// Version       : Version
func (r *CpeResolver) Version() *string {
	return &r.c.Version.String
}

// VendorUpdate  : VendorUpdate
func (r *CpeResolver) VendorUpdate() *string {
	return &r.c.VendorUpdate.String
}

// Edition       : Edition
func (r *CpeResolver) Edition() *string {
	return &r.c.Edition.String
}

// Language      : Language
func (r *CpeResolver) Language() *string {
	return &r.c.Language.String
}

// JvnID : JVNID
func (r *ReferenceResolver) JvnID() *int32 {
	id := int32(r.r.JVNID.Int)
	return &id
}

// NvdID : NVDID
func (r *ReferenceResolver) NvdID() *int32 {
	id := int32(r.r.NVDID.Int)
	return &id
}

// Source : Source
func (r *ReferenceResolver) Source() *string {
	return &r.r.Source.String
}

// Link : Link
func (r *ReferenceResolver) Link() *string {
	return &r.r.Link.String
}

// TotalCount : totalCount
func (r *CveListResolver) TotalCount() int32 {
	return int32(len(r.cveIDs))
}

// Edges : edges
func (r *CveListResolver) Edges() *[]*CveListEdgeResolver {
	l := make([]*CveListEdgeResolver, r.to-r.from)
	for i := range l {
		l[i] = &CveListEdgeResolver{
			cursor: encodeCursor(r.from + i),
			cveID:  r.cveIDs[r.from+i],
		}
	}
	return &l
}

// Cursor : cursor
func (r *CveListEdgeResolver) Cursor() graphql.ID {
	return r.cursor
}

// Node : node
func (r *CveListEdgeResolver) Node() *CveResolver {
	return resolveCve(r.cveID)
}

// Cves : friends
func (r *CveListResolver) Cves() *[]*CveResolver {
	return resolveCves(r.cveIDs)
}

func resolveCve(cveID string) *CveResolver {
	if c, err := models.CveDetailsG(qm.Where("cve_id=?", cveID)).One(); c != nil {
		return &CveResolver{c}
	} else if err != nil {
		log15.Warn("Failed to select data", "method", "cve.resolveCve", "CveID", cveID)
	}
	return nil
}

func resolveCves(cveIDs []string) *[]*CveResolver {
	var cves []*CveResolver
	for _, id := range cveIDs {
		if c := resolveCve(id); c != nil {
			cves = append(cves, c)
		}
	}
	return &cves
}

// PageInfo : pageInfo
func (r *CveListResolver) PageInfo() *PageInfoResolver {
	return &PageInfoResolver{
		startCursor: encodeCursor(r.from),
		endCursor:   encodeCursor(r.to - 1),
		hasNextPage: r.to < len(r.cveIDs),
	}
}

func encodeCursor(i int) graphql.ID {
	return graphql.ID(base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("cursor%d", i+1))))
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

func getUniqueSlice(args []string) []string {
	//	results := make([]string, 0, len(args))
	var results []string
	encountered := map[string]bool{}
	for i := 0; i < len(args); i++ {
		if !encountered[args[i]] {
			encountered[args[i]] = true
			results = append(results, args[i])
		}
	}
	return results
}
