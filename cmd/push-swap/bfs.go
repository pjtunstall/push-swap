package main

import (
	"push-swap/ps"
)

func bfs(s string, n int) ([]string, bool) {
	var result []string

	q := [][]string{{"sa"}, {"ra"}, {"rra"}}
	for len(q) > 0 {
		v := q[0]
		if len(q) == 1 {
			q = [][]string{}
		} else {
			q = q[1:]
		}
		if len(v) >= n {
			break
		}
		a, _ := ps.NewStack(s)
		b, _ := ps.NewStack("")
		ps.Run(&a, &b, v)
		_, sorted := ps.Check(a, b)
		if sorted {
			return v, true
		}

		u := make([]string, len(v))
		copy(u, v)
		switch v[len(v)-1] {
		case "sa":
			q = append(q, append(u, "ra"))
			q = append(q, append(u, "rra"))
		case "ra":
			q = append(q, append(u, "sa"))
			q = append(q, append(u, "ra"))
		case "rra":
			q = append(q, append(u, "sa"))
			q = append(q, append(u, "rra"))
		}
	}

	return result, false
}
