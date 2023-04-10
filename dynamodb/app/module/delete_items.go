package module

// func DeleteItems(items []DDBAgent.ITEM) (totalWCU float64) {
// 	if DDBAgent.DDBAgent == nil {
// 		DDBAgent.NewDDBAgent()
// 	}
// 	//
// 	for index, item := range items {
// 		//
// 		output, err := DDBAgent.DDBAgent.Delete(item)
// 		if err != nil {
// 			panic(err)
// 		}
// 		totalWCU = totalWCU + *output.ConsumedCapacity.CapacityUnits
// 		fmt.Printf("%v => did:%v, CapacityUnits:%v, WCUs:%v\n", index, item.DidModuleID, *output.ConsumedCapacity.CapacityUnits, totalWCU)
// 	}
// 	return
// }
