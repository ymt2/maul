package maul

import (
	"fmt"
	"strings"
)

type (
	repository struct {
		owner string
		name  string
	}

	// Repositories imprements flag.Value interface.
	Repositories []repository
)

// String returns string representation of repositories.
func (rs *Repositories) String() string {
	if rs == nil {
		return ""
	}
	var res []string
	for _, r := range []repository(*rs) {
		res = append(res, fmt.Sprintf("%s/%s", r.owner, r.name))
	}
	return strings.Join(res, ", ")
}

// Set sets a each value presented.
func (r *Repositories) Set(v string) error {
	i := strings.LastIndex(v, "/")
	*r = append(*r, repository{owner: v[:i], name: v[i+1:]})
	return nil
}
