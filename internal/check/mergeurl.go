package check

import "net/url"

// MergeURLValues merges two url.Values maps
func MergeURLValues(q, q1 url.Values) url.Values {

	for k, v := range q1 {
		q[k] = append(q[k], v...)
	}

	return q
}
