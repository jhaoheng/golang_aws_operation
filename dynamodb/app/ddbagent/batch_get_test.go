package DDBAgent

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BatchGetItem(t *testing.T) {
	items := []GetItem{
		0: {"id", "1"},
		1: {"id", "2"},
	}

	output, err := NewDDBAgent("my_table_20230410213618").BatchGetItem(items)
	if !assert.NoError(t, err) {
		t.Fatal()
	}

	fmt.Println(
		func() string {
			b, _ := json.MarshalIndent(output, "", "	")
			return string(b)
		}(),
	)
}
