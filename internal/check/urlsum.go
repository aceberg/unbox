package check

import "net/url"

// SumURL summs two url.Values maps
func SumURL(q, q1 url.Values) url.Values {

	for k, v := range q1 {
		q[k] = append(q[k], v...)
	}

	return q
}
