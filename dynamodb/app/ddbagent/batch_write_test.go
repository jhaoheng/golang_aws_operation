package DDBAgent

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BatchWriteItem(t *testing.T) {
	my_test_items := []MyTestItem{
		0: {ID: 1, Name: "max"},
		1: {ID: 2, Name: "sunny"},
	}
	b, _ := json.Marshal(my_test_items)

	items := []map[string]interface{}{}
	json.Unmarshal(b, &items)
	// fmt.Println(items)

	output, err := NewDDBAgent("my_table_20230410213618").BatchWriteItem(items)
	if !assert.NoError(t, err) {
		t.Fatal()
	}

	b, _ = json.MarshalIndent(output, "", "	")
	fmt.Println(string(b))
}
