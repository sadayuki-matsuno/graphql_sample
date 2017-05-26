package vuls

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
	cve "github.com/kotakanbe/go-cve-dictionary/models"
	graphql "github.com/neelance/graphql-go"
)

var cveData = make(map[string]*cve.CveDetail)

func init() {
	file, e := ioutil.ReadFile("./vuls/cve.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	cveDetail := cve.CveDetail{}
	json.Unmarshal(file, &cveDetail)
	//fmt.Printf("Results: %v\n", cveDetail)

	cveData[cveDetail.CveID] = &cveDetail
	fmt.Println("---init()---")
}

// Cves : Cves
func (r *Resolver) Cves(args *struct{ CveIDs []string }) *[]*CveResolver {

	var cveResolvers []*CveResolver
	spew.Dump(args.CveIDs)
	uniqueCveIDs := getUniqueSlice(args.CveIDs)
	spew.Dump(uniqueCveIDs)
	for _, id := range uniqueCveIDs {
		if c := cveData[id]; c != nil {
			cveResolvers = append(cveResolvers, &CveResolver{c})
		}
	}
	return &cveResolvers
}

// Cve : Cve
func (r *Resolver) Cve(args *struct{ CveID string }) *CveResolver {
	if c := cveData[args.CveID]; c != nil {
		return &CveResolver{c}
	}
	return nil
}

// CveList : cvesConnection
func (r *Resolver) CveList(args *CveListArgs) (*CveListResolver, error) {

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

// CveResolver :
type CveResolver struct {
	c *cve.CveDetail
}

// NvdResolver :
type NvdResolver struct {
	n *cve.Nvd
}

// JvnResolver :
type JvnResolver struct {
	j *cve.Jvn
}

// CveListResolver : CveListResolver
type CveListResolver struct {
	cveIDs []string
	from   int
	to     int
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
	//return r.c.CveID
	cveID := "2022"
	return cveID
}

// Nvd : Nvd
func (r *CveResolver) Nvd() *NvdResolver {
	//	return r.s.Name
	n := &r.c.Nvd
	return &NvdResolver{n}
}

// Jvn : Jvn
func (r *CveResolver) Jvn() *JvnResolver {
	//	return r.s.Name
	j := &r.c.Jvn
	return &JvnResolver{j}
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

// CweID : CweID
func (r *NvdResolver) CweID() *string {
	return &r.n.CweID
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

// CveDetailID : Cvedetailid
func (r *JvnResolver) CveDetailID() *string {
	c := strconv.Itoa(int(r.j.CveDetailID))
	return &c
}

// Title       : Title
func (r *JvnResolver) Title() *string {
	return &r.j.Title
}

// Summary     : Summary
func (r *JvnResolver) Summary() *string {
	return &r.j.Summary
}

// JvnLink : Jvnlink
func (r *JvnResolver) JvnLink() *string {
	return &r.j.JvnLink
}

// JvnID       : Jvnid
func (r *JvnResolver) JvnID() *string {
	return &r.j.JvnID
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
	spew.Dump(r.cveID)
	return resolveCve(r.cveID)
}

// Cves : friends
func (r *CveListResolver) Cves() *[]*CveResolver {
	return resolveCves(r.cveIDs)
}

func resolveCve(cveID string) *CveResolver {
	if c := cveData[cveID]; c != nil {
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
