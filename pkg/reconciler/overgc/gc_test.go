package overgc

import (
	"fmt"
	"testing"
)

func nonactiveRevisions2(revs []bool) []bool {
	swap := len(revs)
	for i := 0; i < swap; {
		if revs[i] {
			swap--
			revs[i] = revs[swap]
		} else {
			i++
		}
	}
	return revs[:swap]
}

func TestA(t *testing.T) {
	a := []bool{
		false,
		true,
		true,
		false,
		false,
		false,
		false,
		true,
		true,
		true,
		true,
		false,
		false,
	}

	fmt.Println(nonactiveRevisions2(a))

}
