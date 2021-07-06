package main

import (
	"Lab1/common"
	"fmt"
	"net/rpc"
)

func main() {

	client, _ := rpc.DialHTTP("tcp", "localhost:9000") // or `localhost:9000`

	var vegetableNamesList []string
	var priceForVegetable float64
	var quantityForVegetable float64
	var isAddVegetableSuccess bool
	var isUpdatePriceSuccess bool
	var isUpdateQuantitySuccess bool
	var vegetableList []common.Vegetable
	var loopRun bool

	loopRun = true

	for loopRun {
		fmt.Println("=======================================================")
		fmt.Println("Please select operation in below list. Type the Number")
		fmt.Println("1 : Display All the Vegetable Names")
		fmt.Println("2 : Display Unit Price for Vegetable Name")
		fmt.Println("3 : Display Available Quantity for Vegetable Name")
		fmt.Println("4 : Add a new Vegetable in to Vegetable List")
		fmt.Println("5 : Update Unit Price for Given Vegetable Name")
		fmt.Println("6 : Update Available Quantity for Given Vegetable Name")
		fmt.Println("7 : Display All the Vegetables With Details")
		fmt.Println("10 : Exit")
		fmt.Print("Please Type Input : ")
		var operation int8
		_, err := fmt.Scanf("%d", &operation)
		if err != nil {
			fmt.Println(err)
			return
		}

		switch operation {
		case 1:
			err := client.Call("Vegetable.GetVegetablesNameList", 0, &vegetableNamesList)
			fmt.Println("=======================================================")

			if err == nil {
				fmt.Printf("				Vegetable Name List")
				for i := 0; i < len(vegetableNamesList); i++ {
					fmt.Printf("%s", vegetableNamesList[i])
				}
			} else {
				panic(err)
			}
			fmt.Println("=======================================================")
			break

		case 2:
			fmt.Println("=======================================================")
			fmt.Print("Please Type Vegetable Name : ")
			var vegetableName string
			_, errScan := fmt.Scanf("%s", &vegetableName)
			if errScan != nil {
				panic(errScan)
			}

			err := client.Call("Vegetable.GetUnitPriceForVegetableName", vegetableName, &priceForVegetable)
			if err == nil {
				fmt.Printf("Price 1kg of %s is %f \n", vegetableName, priceForVegetable)
			} else {
				fmt.Println("Error: ", err)
			}
			fmt.Println("=======================================================")
			break

		case 3:
			fmt.Println("=======================================================")
			fmt.Print("Please Type Vegetable Name : ")
			var vegetableName string
			_, errScan := fmt.Scanf("%s", &vegetableName)
			if errScan != nil {
				panic(errScan)
			}

			err := client.Call("Vegetable.GetQuantityForVegetableName", vegetableName, &quantityForVegetable)
			if err == nil {
				fmt.Printf("Available Quantity of %s is %f \n", vegetableName, quantityForVegetable)
			} else {
				fmt.Println("Error: ", err)
			}
			fmt.Println("=======================================================")
			break
		case 4:
			fmt.Println("=======================================================")
			fmt.Print("Please Type Vegetable Name : ")
			var vegetableName string
			_, errName := fmt.Scanf("%s", &vegetableName)
			if errName != nil {
				panic(errName)
			}
			fmt.Printf("Please Type Unit Price of %s : ", vegetableName)
			var vegetablePrice float64
			_, errPrice := fmt.Scanf("%f", &vegetablePrice)
			if errPrice != nil {
				panic(errPrice)
			}
			fmt.Printf("Please Type Available Quantity of %s : ", vegetableName)
			var vegetableQuantity float64
			_, errQuantity := fmt.Scanf("%f", &vegetableQuantity)
			if errQuantity != nil {
				panic(errQuantity)
			}

			newVegetable := common.Vegetable{
				Name:     vegetableName,
				Price:    vegetablePrice,
				Quantity: vegetableQuantity,
			}

			err := client.Call("Vegetable.AddNewVegetableDetail", newVegetable, &isAddVegetableSuccess)
			if err == nil {
				fmt.Println("Added New Vegetable Successfully")
			} else {
				fmt.Println("Error: ", err)
			}
			fmt.Println("=======================================================")
			break

		case 5:
			fmt.Println("=======================================================")
			fmt.Print("Please Type Vegetable Name : ")
			var vegetableName string
			_, errName := fmt.Scanf("%s", &vegetableName)
			if errName != nil {
				panic(errName)
			}
			fmt.Printf("Please Type updated Unit Price of %s : ", vegetableName)
			var vegetablePrice float64
			_, errPrice := fmt.Scanf("%f", &vegetablePrice)
			if errPrice != nil {
				panic(errPrice)
			}
			updateVegetable := common.Vegetable{
				Name:     vegetableName,
				Price:    priceForVegetable,
				Quantity: 0,
			}

			err := client.Call("Vegetable.UpdateVegetable", updateVegetable, &isUpdatePriceSuccess)
			if err == nil {
				fmt.Printf("Succesfully Updated Price of  %s \n ", vegetableName)
			} else {
				fmt.Println("Error: ", err)
			}

			break
		default:
			fmt.Println("Invalid Operation. please Type correct Number")

		}

	}

}

func getAllVegetableNames() {

}