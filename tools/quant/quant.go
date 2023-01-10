package main 
 
import ( 
	"fmt" 
	"math" 
) 
 
func main() { 
	// 设定投资金额、投资时间、年化收益率
	investment := 1000.00
	years := 10.00
	interestRate := 0.07
 
	// 计算投资收益
	investmentReturn := investment * math.Pow((1 + interestRate), years)
 
	// 输出结果
	fmt.Printf("投资%v元，投资%v年，年化收益率%.2f%%，最终投资收益：%.2f元
", investment, years, interestRate * 100, investmentReturn) 
}
// func main() {
// 	// Get the initial stock price.
// 	initialPrice := float64(50)

// 	// Get the current stock price.
// 	currentPrice := float64(60)

// 	// Calculate the return on the stock.
// 	returnOnStock := (currentPrice - initialPrice) / initialPrice

// 	// Print the return on the stock.
// 	fmt.Printf("The stock return is %.2f%%
// ", math.Abs(returnOnStock)*100)
// }
