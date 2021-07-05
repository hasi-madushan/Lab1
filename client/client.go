package main

import (
	"Lab1/common"
	"fmt"
	"net/rpc"
)

func main() {

	// get RPC client by dialing at `rpc.DefaultRPCPath` endpoint
	client, _ := rpc.DialHTTP("tcp", "localhost:9000") // or `localhost:9000`

	/*--------------*/

	// create john variable of type `common.Student`
	var vegNames []string
	var unitPrice, availableQty float32
	var vegs []common.Vegetable

	fmt.Println("Select a Command")
	fmt.Println("1 : Get All Vegetable Names ")
	fmt.Println("2 : Get UnitPrice ")
	fmt.Println("3 : Get Available Qty ")
	fmt.Println("4 : Add a new Vegetable")
	fmt.Println("5 : Update a Vegetable ")

	//reader := bufio.NewReader(os.Stdin)
	var command int8
	_, err := fmt.Scanf("%d", &command)

	if err != nil {
		fmt.Println(err)
		return
	}

	switch command {
	case 1:
		/*--------------Get All Vegetables ---------*/
		if err := client.Call("Vegetable.GetAllVegetableNames", 0, &vegNames); err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("All Vegetable Names = %v \n", vegNames)
		}
		break
	case 2:
		/*--------------Get UnitPrice ---------*/

		if err := client.Call("Vegetable.GetUnitPrice", 1, &unitPrice); err != nil {
			fmt.Println("Error: Vegetable.getUnitPrice()", err)
		} else {
			fmt.Printf("Vegetable.id = %v | UnitPrice  = %v \n", 1, unitPrice)
		}
		break
	case 3:
		/*--------------Get Available Qty ---------*/

		if err := client.Call("Vegetable.GetAvailableQty", 1, &availableQty); err != nil {
			fmt.Println("Error: Vegetable.GetAvailableQty()", err)
		} else {
			fmt.Printf("Vegetable.id = %v | Available Qty  = %v \n", 1, availableQty)
		}
		break
	case 4:
		/*--------------Add a new Vegetable ---------*/
		newVeg := common.Vegetable{
			Id:           1,
			Name:         "Tomato",
			UnitPrice:    30.20,
			AvailableQty: 5,
		}

		if err := client.Call("Vegetable.AddNewVegetable", newVeg, &vegs); err != nil {
			fmt.Println("Error: Vegetable.Add()", err)
		} else {
			fmt.Printf("Added Succesfully \n")
		}
		break
	case 5:
		/*--------------Update a Vegetable ---------*/

		updateVeg := common.Vegetable{
			Id:           1,
			AvailableQty: 25.4,
		}

		if err := client.Call("Vegetable.UpdateVegetable", updateVeg, &updateVeg); err != nil {
			fmt.Println("Error: Vegetable.UpdateVegetable()", err)
		} else {
			fmt.Printf("Succesfully Updated Vegetable %v \n ", updateVeg)
		}
		break
	default:
		fmt.Println("Invalid Command")

	}

}

func getAllVegetableNames() {

}
