package DDBAgent

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PutItem(t *testing.T) {
	table_name := "my_test"

	my_test_item := MyTestItem{
		ID:   1,
		Name: "max",
	}

	output, err := NewDDBAgent(table_name).PutItem(my_test_item)
	if !assert.NoError(t, err) {
		t.Fatal()
	}
	b, _ := json.MarshalIndent(output, "", "	")
	fmt.Println(string(b))
}
