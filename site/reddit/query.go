package reddit

import (
	"fmt"
	"net/url"
	"strings"

	graw "github.com/turnage/graw/reddit"
)

type QueryParam struct {
	v interface{}
}

func (q QueryParam) String() string {
	switch v := q.v.(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	default:
		return ""
	}
}

type QueryParams interface {
	GetParams(name string) []QueryParam
}

type queryParams map[string][]QueryParam

func (q queryParams) GetParams(name string) []QueryParam {
	rv, exists := q[name]
	if !exists {
		return nil
	}
	return rv
}

func (q queryParams) add(name string, qp QueryParam) {
	q[name] = append(q[name], qp)
}

// PostQueryClause defines a predicate on a post.
type PostQueryClause interface {
	// CheckPost checkes if the post match the query.
	CheckPost(QueryParams, *graw.Post) bool
}

type PostQueryClauseFunc func(QueryParams, *graw.Post) bool

func (f PostQueryClauseFunc) CheckPost(q QueryParams, p *graw.Post) bool {
	return f(q, p)
}

type QueryRegistry struct {
	queryerByNames map[string]PostQueryClause
	freezed        bool
}

func newQueryRegistry() *QueryRegistry {
	return &QueryRegistry{
		queryerByNames: map[string]PostQueryClause{},
		freezed:        false,
	}
}

var defaultRegister = newQueryRegistry()

func (r *QueryRegistry) ensureUnfreezed() {
	if r.freezed {
		panic("query registry is already freezed")
	}
}

func (r *QueryRegistry) freeze() {
	r.ensureUnfreezed()
	r.freezed = true
}

func (r *QueryRegistry) register(name string, q PostQueryClause) {
	r.ensureUnfreezed()

	r.queryerByNames[name] = q
}

func (r *QueryRegistry) Get(name string) (PostQueryClause, bool) {
	q, exists := r.queryerByNames[name]
	return q, exists
}

func GetQuery(name string) (PostQueryClause, bool) {
	return defaultRegister.Get(name)
}

const queryUrlValueKeyPrefix = "q:"

func ParseQueryFromURLValues(vs url.Values) (QueryParams, []PostQueryClause) {
	qp := &queryParams{}
	var qcs []PostQueryClause
	qcLoaded := map[string]struct{}{}
	for k, vv := range vs {
		if !strings.HasPrefix(queryUrlValueKeyPrefix, k) || k == queryUrlValueKeyPrefix {
			continue
		}
		name := k[len(queryUrlValueKeyPrefix):]

		qc, exists := GetQuery(name)
		if !exists {
			continue
		}
		if _, loaded := qcLoaded[name]; !loaded {
			qcs = append(qcs, qc)
			qcLoaded[name] = struct{}{}
		}

		for v := range vv {
			qp.add(name, QueryParam{v: v})
		}
	}
	return qp, nil
}

func init() {
	defaultRegister.freeze()
}
