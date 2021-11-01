package module

import (
	"dynamodb/app/ddbservice"
	"testing"
)

func Test_DeleteItems(t *testing.T) {
	var items = []ddbservice.ITEM{
		0: {
			DidModuleID: "d7e5ce2b-0017-40a0-bba9-d07a7679485c",
			CreatedTs:   "1635685852.001",
		},
	}
	DeleteItems(items)
}
