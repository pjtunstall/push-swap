package ps

import (
	"errors"
)

func Run(a, b *Stack, instructions []string) error {
	for _, instruction := range instructions {
		switch instruction {
		case "sa":
			Sx(a)
		case "sb":
			Sx(b)
		case "ss":
			Ss(a, b)
		case "pa":
			Px(a, b)
		case "pb":
			Px(b, a)
		case "ra":
			Rx(a)
		case "rb":
			Rx(b)
		case "rr":
			Rr(a, b)
		case "rra":
			Rrx(a)
		case "rrb":
			Rrx(b)
		case "rrr":
			Rrr(a, b)
		default:
			return errors.New("invalid instruction")
		}
	}
	return nil
}
