package main

import (
	"push-swap/ps"
)

func bfs(c ps.Stack, n int) []string {
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
		a, _ := ps.NewStack(c.GetNumsString())
		b, _ := ps.NewStack("")
		inv := inverse(v)
		ps.Run(&a, &b, inv)
		_, sorted := ps.Check(a, b)
		if sorted {
			return v
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

	return result
}

func inverse(instructions []string) []string {
	var result []string
	for i := len(instructions) - 1; i >= 0; i-- {
		switch instructions[i] {
		case "sa":
			result = append(result, "sa")
		case "sb":
			result = append(result, "sb")
		case "ss":
			result = append(result, "ss")
		case "pa":
			result = append(result, "pb")
		case "pb":
			result = append(result, "pa")
		case "ra":
			result = append(result, "rra")
		case "rb":
			result = append(result, "rrb")
		case "rr":
			result = append(result, "rrr")
		case "rra":
			result = append(result, "ra")
		case "rrb":
			result = append(result, "rb")
		case "rrr":
			result = append(result, "rr")
		}
	}
	return result
}
