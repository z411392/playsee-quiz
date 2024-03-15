package queries

import (
	"fmt"
	"strings"

	LinkedList "playsee.co/interview/utils/dataStructure/LinkedList"
)

func ReadAsLinkedList(values ...interface{}) string {
	linkedList := LinkedList.FromVariadic(values...)
	node := linkedList.Head
	if node == nil {
		return ""
	}
	builder := &strings.Builder{}
	index := 0
	linkedList.Iterate(func(current *LinkedList.Node, next *LinkedList.Node, previous *LinkedList.Node) {
		if previous == nil {
			builder.WriteString(fmt.Sprintf("head -> %v\n", current.Value))
		}
		if next == nil {
			builder.WriteString(fmt.Sprintf("tail -> %v\n", current.Value))
			return
		}
		if previous == nil {
			return
		}
		index += 1
		builder.WriteString(fmt.Sprintf("node%d -> %v\n", index, current.Value))
	})
	return fmt.Sprintf("\n%s", builder.String())
}
