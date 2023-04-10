package DDBAgent

import (
	"os"
	"testing"
)

type MyTestItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestMain(m *testing.M) {
	m.Run()
	os.Exit(0)
}
