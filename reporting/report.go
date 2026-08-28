package reporting

import (
	"fmt"
	"frontend_go/domain"
	"sort"
	"strings"
	"time"
)

type Summary struct {
	Total    int
	ByStatus map[string]int
	Latest   time.Time
}

func Summarize(records []domain.Record) Summary {
	out := Summary{ByStatus: map[string]int{}}
	for _, r := range records {
		out.Total++
		out.ByStatus[r.Status]++
		if r.UpdatedAt.After(out.Latest) {
			out.Latest = r.UpdatedAt
		}
	}
	return out
}
func RenderTree(root *domain.TreeNode) string {
	lines := []string{}
	var walk func(*domain.TreeNode, int)
	walk = func(n *domain.TreeNode, d int) {
		for _, c := range n.Children {
			lines = append(lines, strings.Repeat("  ", d)+fmt.Sprintf("%s [%s]", c.Participant.Name, c.Participant.State))
			walk(c, d+1)
		}
	}
	walk(root, 0)
	sort.Strings(lines)
	return strings.Join(lines, "\n")
}
