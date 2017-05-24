package vuls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	cve "github.com/kotakanbe/go-cve-dictionary/models"
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
	fmt.Printf("Results: %v\n", cveDetail)

	cveData[cveDetail.CveID] = &cveDetail
}

// Cve : Cve
func (r *Resolver) Cve(args *struct{ CveID string }) *CveResolver {
	if c := cveData[args.CveID]; c != nil {
		return &CveResolver{c}
	}
	//	return &CveResolver{
	//		&cve{
	//			CveID: "2022",
	//			Name:  "CVE-0001",
	//		},
	//	}
	return nil
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

// CveID : id
func (r *CveResolver) CveID() string {
	//return r.c.CveID
	return "2022"
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
