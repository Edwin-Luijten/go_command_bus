package commandbus

import (
	"sort"
	"testing"
)

func TestSortByPriority(t *testing.T) {
	items := make([]middleware, 0)

	items = append(items, middleware{function: func(command interface{}, next HandlerFunc) {}, priority: 0})
	items = append(items, middleware{function: func(command interface{}, next HandlerFunc) {}, priority: 1})
	items = append(items, middleware{function: func(command interface{}, next HandlerFunc) {}, priority: 2})

	sort.Sort(sortByPriority(items))

	if items[0].priority != 2 {
		t.Log("Unexpected priority")
		t.Fail()
	}
}
