package main

import (
	"errors"

	"push-swap/ps"
)

func run(a, b *ps.Stack, instructions []string) error {
	for _, instruction := range instructions {
		switch instruction {
		case "sa":
			ps.Sx(a)
		case "sb":
			ps.Sx(b)
		case "ss":
			ps.Ss(a, b)
		case "pa":
			ps.Px(a, b)
		case "pb":
			ps.Px(b, a)
		case "ra":
			ps.Rx(a)
		case "rb":
			ps.Rx(b)
		case "rr":
			ps.Rr(a, b)
		case "rra":
			ps.Rrx(a)
		case "rrb":
			ps.Rrx(b)
		case "rrr":
			ps.Rrr(a, b)
		default:
			return errors.New("invalid instruction")
		}
	}
	return nil
}
