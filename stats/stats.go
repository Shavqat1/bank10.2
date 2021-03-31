package stats

import "github.com/Shavqat1/bank10.2/pkg/bank/types"

//TotalInCategory находит сумму покупок в определённой категории.
func TotalInCategory(payments []types.Payment,category types.Category)types.Money{
	sum:=int64(0)
	for _, payment:=range payments {
	if (category==payment.Category){
		sum+=int64(payment.Amount)
	}

	}
	return types.Money(sum) 
}