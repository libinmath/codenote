package main

import (
	"fmt"
	"gostock"
)

func main() {
	// 获取股票价格
	price, err := gostock.GetPrice("GOOG")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 使用股票价格历史数据预测未来价格
	futurePrice, err := gostock.PredictPrice(price)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("预测GOOG股价为%f
", futurePrice)
}
