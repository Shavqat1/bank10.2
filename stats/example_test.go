package stats

import (
	"fmt"

	"github.com/Shavqat1/bank10.2/pkg/bank/types"
)
func ExampleTotalInCategory() {
payments :=[]types.Payment{
    {
     Amount: 10_000,
	 Category: "Restaurant",
	},
	{
		Amount: 12_000,
		Category: "Restaurant",
	},
	}
	Sum:=TotalInCategory(payments,"Restaurant")
	fmt.Println(Sum)
	//Output:
	//22000

}