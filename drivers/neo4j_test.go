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

	if err := manager.ExecuteWrite(context.Background(), "CREATE (person1:Person {name: \"Alice\", age: 30, title: \"Test1\"})", nil); err != nil {
		t.Error(err)
	}
	if err := manager.ExecuteWrite(context.Background(), "CREATE (person2:Person {name: \"Bob\", age: 25,title: \"Test2\"})", nil); err != nil {
		t.Error(err)
	}
	if err := manager.ExecuteWrite(context.Background(), "CREATE (city:City {name: \"New York\",title: \"Test3\"})", nil); err != nil {
		t.Error(err)
	}

	res, err := manager.ExecuteRead(context.Background(), "MATCH (n) RETURN n", nil)
	if err != nil {
		t.Error(err)
	}
	for i, v := range res.Records {
		fmt.Println(i, v.AsMap())
	}
	//if err := manager.ExecuteWrite(context.Background(), "MATCH (n) DETACH DELETE n", nil); err != nil {
	//	t.Error(err)
	//}
}
