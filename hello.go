package main

import "fmt"

func main() {

	var income, food, transport, housing float64

	fmt.Print("Введите ваш месячный доход: ")
	fmt.Scan(&income)

	fmt.Print("Введите расходы на еду: ")
	fmt.Scan(&food)

	fmt.Print("Введите расходы на транспорт: ")
	fmt.Scan(&transport)

	fmt.Print("Введите расходы на жилье: ")
	fmt.Scan(&housing)

	remainingIncome := income - (food + transport + housing)

	fmt.Println("Чистый доход: ", remainingIncome)

	if food/income >= 0.3 {
		fmt.Println("Внимание: расходы на еду составляют больше 30% от вашего дохода. Рекомендуется снизить траты в этой категории.")
	}
	if transport/income >= 0.3 {
		fmt.Println("Внимание: расходы на транспорт составляют больше 30% от вашего дохода. Рекомендуется снизить траты в этой категории.")
	}
	if housing/income >= 0.3 {
		fmt.Println("Внимание: расходы на жильё составляют больше 30% от вашего дохода. Рекомендуется снизить траты в этой категории.")
	}
}
