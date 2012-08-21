package neo4go

import (
	"testing"
)

var URL = "http://localhost:7474"

func TestNewNeo4j(t *testing.T) {
	_, err := NewNeo4j(URL)
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Neo4j connection established.")
	}
}

func TestGetNode(t *testing.T) {
	neo, _ := NewNeo4j(URL)
	// Node 0 should always exist
	node, err := neo.GetNode(100)

	if err != nil {
		t.Fatal(err)
	}

	if node.Id != 1 {
		t.Error("Wrong node returned in GetNode")
	}
}
