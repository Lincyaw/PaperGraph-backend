package drivers

import (
	"context"
	"fmt"
	"testing"
)

func TestNewNeo4jSessionManager(t *testing.T) {
	manager, err := NewNeo4jSessionManager("bolt://10.10.10.196:7687", "", "")
	if err != nil {
		t.Error(err)
	}

	if err := manager.ExecuteWrite(context.Background(), "CREATE (person1:Person {name: \"Alice\", age: 30})", nil); err != nil {
		t.Error(err)
	}
	if err := manager.ExecuteWrite(context.Background(), "CREATE (person2:Person {name: \"Bob\", age: 25})", nil); err != nil {
		t.Error(err)
	}
	if err := manager.ExecuteWrite(context.Background(), "CREATE (city:City {name: \"New York\"})", nil); err != nil {
		t.Error(err)
	}

	res, err := manager.ExecuteRead(context.Background(), "MATCH (n) RETURN n", nil)
	if err != nil {
		t.Error(err)
	}
	single, err := res.Single(context.Background())
	if err != nil {
		t.Error(err)
	}
	fmt.Println(single)
}
