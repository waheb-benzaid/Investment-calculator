package main

import "fmt"

func main()  {

	revenue := 0.0
	expenses := 0.0
	taxReate := 0.0

	fmt.Println("profit calculator app")
	
	fmt.Println("please enter your revenue!")
	fmt.Scan(&revenue)

	fmt.Println("please enter your expenses!")
	fmt.Scan(&expenses)

	fmt.Println("please enter your tax rate!")
	fmt.Scan(&taxReate)

	ebt := revenue - expenses
	profit := ebt * taxReate
	ratio := ebt/profit
	
	fmt.Printf("Ebt is : %f , profit is %f , ratio is : %f\n", ebt, profit, ratio)
}