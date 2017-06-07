package vuls

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/inconshreveable/log15"
	cveModels "github.com/kotakanbe/go-cve-dictionary/models"
	graphql "github.com/neelance/graphql-go"
	"github.com/sadayuki-matsuno/graphql_sample/db"
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
		if c, err := db.GetCveByID(id); err != nil {
			log15.Warn("Failed to select data", "CveID", id, "err", err)
		} else {
			cveResolvers = append(cveResolvers, &CveResolver{c})
		}
	}
	log15.Info("Selected data from cves", "method", "cve.Cves", "cveResolvers", cveResolvers)
	return &cveResolvers
}

// CveList : cvesConnection
func (r *Resolver) CveList(args *CveListArgs) (*CveListResolver, error) {

	log15.Info("Select data from cvelist", "method", "cve.CveList", "args.CveIDs", args.CveIDs)
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

	log15.Info("return CveListResolver", "method", "cve.CveList", "uniqueCveIDs", uniqueCveIDs)
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
	c *cveModels.CveDetail
}

// NvdResolver :
type NvdResolver struct {
	n *cveModels.Nvd
}

// JvnResolver :
type JvnResolver struct {
	j *cveModels.Jvn
}

// CpeResolver : CpeResolver
type CpeResolver struct {
	c cveModels.Cpe
}

// ReferenceResolver : ReferenceResolver
type ReferenceResolver struct {
	r cveModels.Reference
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
	return r.c.CveID
}

// Nvd : Nvd
func (r *CveResolver) Nvd() *NvdResolver {
	return &NvdResolver{&r.c.Nvd}
}

// Jvn : Jvn
func (r *CveResolver) Jvn() *JvnResolver {
	return &JvnResolver{&r.c.Jvn}
}

// ID : ID
func (r *NvdResolver) ID() string {
	return "NVD_" + r.n.CveID
}

// Summary : Summary
func (r *NvdResolver) Summary() *string {
	return &r.n.Summary
}

// Score :  Score
func (r *NvdResolver) Score() *float64 {
	return &r.n.Score
}

// AccessVector : AccessVector
func (r *NvdResolver) AccessVector() *string {
	return &r.n.AccessVector
}

// AccessComplexity : AccessComplexity
func (r *NvdResolver) AccessComplexity() *string {
	return &r.n.AccessComplexity
}

// Authentication : Authentication
func (r *NvdResolver) Authentication() *string {
	return &r.n.Authentication
}

// ConfidentialityImpact : ConfidentialityImpact
func (r *NvdResolver) ConfidentialityImpact() *string {
	return &r.n.ConfidentialityImpact
}

// IntegrityImpact : IntegrityImpact
func (r *NvdResolver) IntegrityImpact() *string {
	return &r.n.IntegrityImpact
}

// AvailabilityImpact : AvailabilityImpact
func (r *NvdResolver) AvailabilityImpact() *string {
	return &r.n.AvailabilityImpact
}

// Cpes  : Nvd
func (r *NvdResolver) Cpes() *[]*CpeResolver {

	var cpeResolvers []*CpeResolver
	for _, cpe := range r.n.Cpes {
		log15.Info("check cpe", "method", "cve.Cpes", "cpe", cpe)
		cpeResolvers = append(cpeResolvers, &CpeResolver{cpe})
	}
	log15.Info("check cpe", "method", "cve.Cpes", "cpeResolvers", cpeResolvers)
	return &cpeResolvers
}

// CweID : CweID
func (r *NvdResolver) CweID() *string {
	return &r.n.CweID
}

// References  : Nvd
func (r *NvdResolver) References() *[]*ReferenceResolver {
	var referenceResolvers []*ReferenceResolver
	for _, r := range r.n.References {
		referenceResolvers = append(referenceResolvers, &ReferenceResolver{r})
	}
	return &referenceResolvers
}

// PublishedDate : PublishedDate
func (r *NvdResolver) PublishedDate() *string {
	t := r.n.PublishedDate.String()
	return &t
}

// LastModifiedDate : LastModifiedDate
func (r *NvdResolver) LastModifiedDate() *string {
	t := r.n.LastModifiedDate.String()
	return &t
}

// ID : ID
func (r *JvnResolver) ID() string {
	return "JVN_" + r.j.CveID
}

// CveDetailID : Cvedetailid
func (r *JvnResolver) CveDetailID() *int32 {
	id := int32(r.j.CveDetailID)
	return &id
}

// Title       : Title
func (r *JvnResolver) Title() *string {
	return &r.j.Title
}

// Summary     : Summary
func (r *JvnResolver) Summary() *string {
	return &r.j.Summary
}

// JVNID : JvnID
func (r *JvnResolver) JVNID() *string {
	return &r.j.JvnID
}

// JVNLink : Jvnlink
func (r *JvnResolver) JVNLink() *string {
	return &r.j.JvnLink
}

// Score       : Score
func (r *JvnResolver) Score() *float64 {
	return &r.j.Score
}

// Severity    : Severity
func (r *JvnResolver) Severity() *string {
	return &r.j.Severity
}

// Vector      : Vector
func (r *JvnResolver) Vector() *string {
	return &r.j.Vector
}

// Cpes  : Cpe
func (r *JvnResolver) Cpes() *[]*CpeResolver {

	//	jvnID := &r.j.ID
	//	log15.Info("Select data from cpe", "method", "cve.Cpe", "jvn_id", jvnID)
	//	if cpes, err := models.CpesG(qm.Where("jvn_id=?", jvnID)).All(); cpes != nil {
	//		log15.Info("select multi data", "cpes", cpes)
	//		var cpeResolvers []*CpeResolver
	//		for _, c := range cpes {
	//			cpeResolvers = append(cpeResolvers, &CpeResolver{c})
	//		}
	//		return &cpeResolvers
	//	} else if err != nil {
	//		log15.Warn("Failed to select data", "method", "cve.Cpe", "err", err)
	//	}
	//	return nil
	var cpeResolvers []*CpeResolver
	for _, cpe := range r.j.Cpes {
		cpeResolvers = append(cpeResolvers, &CpeResolver{cpe})
	}
	return &cpeResolvers

}

// References  : Jvn
func (r *JvnResolver) References() *[]*ReferenceResolver {

	//	jvnID := &r.j.ID
	//	log15.Info("Select data from cpe", "method", "cve.Reference", "jvn_id", jvnID)
	//	if cpes, err := models.ReferencesG(qm.Where("jvn_id=?", jvnID)).All(); cpes != nil {
	//		log15.Info("select multi data", "cpes", cpes)
	//		var cpeResolvers []*ReferenceResolver
	//		for _, c := range cpes {
	//			cpeResolvers = append(cpeResolvers, &ReferenceResolver{c})
	//		}
	//		return &cpeResolvers
	//	} else if err != nil {
	//		log15.Warn("Failed to select data", "method", "cve.Reference", "err", err)
	//	}
	//	return nil
	var referenceResolvers []*ReferenceResolver
	for _, r := range r.j.References {
		referenceResolvers = append(referenceResolvers, &ReferenceResolver{r})
	}
	return &referenceResolvers

}

// PublishedDate : PublishedDate
func (r *JvnResolver) PublishedDate() *string {
	t := r.j.PublishedDate.String()
	return &t
}

// LastModifiedDate : LastModifiedDate
func (r *JvnResolver) LastModifiedDate() *string {
	t := r.j.LastModifiedDate.String()
	return &t
}

// CpeName : CpeName
func (r *CpeResolver) CpeName() *string {
	return &r.c.CpeName
}

// Part : Part
func (r *CpeResolver) Part() *string {
	return &r.c.Part
}

// Vendor : Vendor
func (r *CpeResolver) Vendor() *string {
	return &r.c.Vendor
}

// Product : Product
func (r *CpeResolver) Product() *string {
	return &r.c.Product
}

// Version       : Version
func (r *CpeResolver) Version() *string {
	return &r.c.Version
}

// VendorUpdate  : VendorUpdate
func (r *CpeResolver) VendorUpdate() *string {
	return &r.c.VendorUpdate
}

// Edition       : Edition
func (r *CpeResolver) Edition() *string {
	return &r.c.Edition
}

// Language      : Language
func (r *CpeResolver) Language() *string {
	return &r.c.Language
}

// Source : Source
func (r *ReferenceResolver) Source() *string {
	return &r.r.Source
}

// Link : Link
func (r *ReferenceResolver) Link() *string {
	return &r.r.Link
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
	var err error
	var c *cveModels.CveDetail
	if c, err = db.GetCveByID(cveID); err != nil {
		log15.Warn("Failed to select data", "method", "cve.resolveCve", "CveID", cveID)
	} else if c != nil {
		return &CveResolver{c}
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
