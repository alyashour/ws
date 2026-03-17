package tasks

import (
	"fmt"
	"strconv"
	"strings"
)

type ID string

// generates the next raw id
// looks for the max and returns max + 1
// does not fill in holes
// reuses ID if max is deleted (lowk feel like that's ok)
func genNextID(existing []ID) ID {
	// find max
	max := int64(0)
	for _, id := range existing {
		n, _ := strconv.ParseInt(string(id), 36, 64)
		if n > max {
			max = n
		}
	}

	// format max in base 36
	return ID(fmt.Sprintf("%04s", strconv.FormatInt(max+1, 36)))
}

func FormatID(id ID, prefix string) string {
	s := strings.TrimLeft(string(id), "0")
	if s == "" {
		s = "0"
	}
	return fmt.Sprintf("%s", s)
}

func normalizeID(input string) ID {
	// strip prefix if present (t, p, j etc.)
	input = strings.TrimLeftFunc(input, func(r rune) bool {
		return !('0' <= r && r <= '9') && !('a' <= r && r <= 'z')
	})
	// parse as base36 and reformat to 4 chars
	n, _ := strconv.ParseInt(input, 36, 64)
	return ID(fmt.Sprintf("%04s", strconv.FormatInt(n, 36)))
}
