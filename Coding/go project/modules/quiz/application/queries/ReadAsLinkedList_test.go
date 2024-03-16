package queries

import (
	"testing"
)

const message = `
expected: %s
got: %s
`

func TestWhenGivenEmptyArguments(test *testing.T) {
	expected := ""
	got, _ := ReadAsLinkedList()
	if got != expected {
		test.Errorf(message, expected, got)
	}
}

func TestWhenGivenOneElement(test *testing.T) {
	expected := `
head -> a
tail -> a
`
	got, _ := ReadAsLinkedList("a")
	if got != expected {
		test.Errorf(message, expected, got)
	}
}

func TestWhenGivenTwoElements(test *testing.T) {
	expected := `
head -> a
tail -> b
`
	got, _ := ReadAsLinkedList("a", "b")
	if got != expected {
		test.Errorf(message, expected, got)
	}
}

func TestWhenGivenThreeElements(test *testing.T) {
	expected := `
head -> a
node1 -> b
tail -> c
`
	got, _ := ReadAsLinkedList("a", "b", "c")
	if got != expected {
		test.Errorf(message, expected, got)
	}
}

func TestWhenGivenFourElements(test *testing.T) {
	expected := `
head -> a
node1 -> b
node2 -> c
tail -> d
`
	got, _ := ReadAsLinkedList("a", "b", "c", "d")
	if got != expected {
		test.Errorf(message, expected, got)
	}
}

func TestWhenGivenFiveElements(test *testing.T) {
	expected := `
head -> a
node1 -> b
node2 -> c
node3 -> d
tail -> e
`
	got, _ := ReadAsLinkedList("a", "b", "c", "d", "e")
	if got != expected {
		test.Errorf(message, expected, got)
	}
}
