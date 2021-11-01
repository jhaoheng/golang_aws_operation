package module

import (
	"dynamodb/app/ddbservice"
	"fmt"
)

func DeleteItems(items []ddbservice.ITEM) (totalWCU float64) {
	if ddbservice.DDBService == nil {
		ddbservice.NewDDB()
	}
	//
	for index, item := range items {
		//
		output, err := ddbservice.DDBService.Delete(item)
		if err != nil {
			panic(err)
		}
		totalWCU = totalWCU + *output.ConsumedCapacity.CapacityUnits
		fmt.Printf("%v => did:%v, CapacityUnits:%v, WCUs:%v\n", index, item.DidModuleID, *output.ConsumedCapacity.CapacityUnits, totalWCU)
	}
	return
}
