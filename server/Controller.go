package main

import (
	"Lab1/common"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

var mutex sync.Mutex

type Vegetable common.Vegetable

func ReadAllVegetables() []Vegetable {
	vegetablesFile, err := os.Open("vegetables.json")
	if err != nil {
		panic(err)
	}
	defer vegetablesFile.Close()
	vegetablesByteDetails, _ := ioutil.ReadAll(vegetablesFile)

	var vegetablesList []Vegetable
	json.Unmarshal(vegetablesByteDetails, &vegetablesList)

	return vegetablesList
}

func WriteAllVegetables(vegetablesList []Vegetable) {
	vegetablesFile, _ := os.Open("vegetables.json")
	vegetablesByteDetails, _ := json.Marshal(vegetablesList)
	ioutil.WriteFile("vegetables.json", vegetablesByteDetails, 0644)
	defer vegetablesFile.Close()
}

func (v *Vegetable) GetAllVegetablesDetailsList(id int, reply *[]Vegetable) error {
	*reply = ReadAllVegetables()
	return nil
}

func (v *Vegetable) GetVegetableDetails(name string, reply *Vegetable) error {
	vegetablesList := ReadAllVegetables()
	for i := 0; i < len(vegetablesList); i++ {
		if name == vegetablesList[i].Name {
			*reply = vegetablesList[i]
			return nil
		}
	}
	return fmt.Errorf("%s is not in the Vegetable list", name)
}

func (v *Vegetable) GetVegetablesNameList(id int, reply *[]string) error {
	vegetablesList := ReadAllVegetables()
	var VegetableNamesList []string
	for i := 0; i < len(vegetablesList); i++ {
		VegetableNamesList = append(VegetableNamesList, vegetablesList[i].Name)
	}
	*reply = VegetableNamesList

	return nil
}

func (v *Vegetable) GetUnitPriceForVegetableName(name string, reply *float64) error {
	vegetablesList := ReadAllVegetables()
	for i := 0; i < len(vegetablesList); i++ {
		if name == vegetablesList[i].Name {
			*reply = vegetablesList[i].Price
			return nil
		}
	}

	return fmt.Errorf("%s is not in the Vegetable list", name)

}

func (v *Vegetable) GetQuantityForVegetableName(name string, reply *float64) error {
	vegetablesList := ReadAllVegetables()
	for i := 0; i < len(vegetablesList); i++ {
		if name == vegetablesList[i].Name {
			*reply = vegetablesList[i].Quantity
			return nil
		}
	}

	return fmt.Errorf("%s is not in the Vegetable list", name)

}

func (v *Vegetable) AddNewVegetableDetail(vegetable Vegetable, reply *bool) error {
	mutex.Lock()
	defer mutex.Unlock()

	vegetablesList := ReadAllVegetables()
	for i := 0; i < len(vegetablesList); i++ {
		if vegetable.Name == vegetablesList[i].Name {
			return fmt.Errorf("%s is already in the vegetable list", vegetable.Name)
		}
	}

	vegetablesList = append(vegetablesList, vegetable)
	WriteAllVegetables(vegetablesList)

	*reply = true
	return nil
}

func (v *Vegetable) UpdatePriceOfVegetableByName(vegetable Vegetable, reply *Vegetable) error {
	mutex.Lock()
	defer mutex.Unlock()

	vegetablesList := ReadAllVegetables()
	for i := 0; i < len(vegetablesList); i++ {
		if vegetable.Name == vegetablesList[i].Name {
			if vegetable.Price < 0 {
				return fmt.Errorf("%s's price is wrong.Please add correct Price", vegetable.Name)
			}
			vegetablesList[i].Price = vegetable.Price
			WriteAllVegetables(vegetablesList)
			*reply = vegetable
			return nil
		}
	}

	return fmt.Errorf("%s is not in the Vegetable list", vegetable.Name)
}
func (v *Vegetable) UpdateQuantityOfVegetableByName(vegetable Vegetable, reply *Vegetable) error {
	mutex.Lock()
	defer mutex.Unlock()

	vegetablesList := ReadAllVegetables()
	for i := 0; i < len(vegetablesList); i++ {
		if vegetable.Name == vegetablesList[i].Name {
			if vegetable.Quantity < 0 {
				return fmt.Errorf("%s's Quantity is wrong.Please add correct Price", vegetable.Name)
			}
			vegetablesList[i].Quantity = vegetable.Quantity
			WriteAllVegetables(vegetablesList)
			*reply = vegetable
			return nil
		}
	}

	return fmt.Errorf("%s is not in the Vegetable list", vegetable.Name)
}

func NewVegetableController() *Vegetable {
	return &Vegetable{}
}
