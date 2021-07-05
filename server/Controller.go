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

func Read() []Vegetable {
	// Open jsonFile
	jsonFile, err := os.Open("vegetables.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	// defer the closing of jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize Vegetable array
	var vegetables []Vegetable

	json.Unmarshal(byteValue, &vegetables)

	return vegetables

}

func Write(data []Vegetable) {
	jsonFile, _ := os.Open("vegetables.json")
	byte, _ := json.Marshal(data)
	ioutil.WriteFile("vegetables.json", byte, 0644)
	defer jsonFile.Close()
}

func (v *Vegetable) GetAllVegetableNames(payload int, reply *[]string) error {
	allVegs := Read()
	var allVegNames []string
	for i := 0; i < len(allVegs); i++ {
		allVegNames = append(allVegNames, allVegs[i].Name)
	}

	*reply = allVegNames
	return nil

}
func (v *Vegetable) AddNewVegetable(veg Vegetable, reply *Vegetable) error {
	mutex.Lock()
	defer mutex.Unlock()
	allVegs := Read()

	for i := 0; i < len(allVegs); i++ {
		if veg.Id == allVegs[i].Id {
			return fmt.Errorf("Already exist vegetable.id = %v ", veg.Id)
		}
	}
	allVegs = append(allVegs, veg)
	Write(allVegs)
	*reply = veg
	return nil

}
func (v *Vegetable) GetUnitPrice(id int, reply *float32) error {
	fmt.Println("hasotha")
	allVegs := Read()

	for i := 0; i < len(allVegs); i++ {
		if id == allVegs[i].Id {
			*reply = allVegs[i].UnitPrice
			return nil
		}
	}
	return fmt.Errorf("Not found vegetable.id =%v ", id)

}
func (v *Vegetable) GetAvailableQty(id int, reply *float32) error {
	allVegs := Read()

	for i := 0; i < len(allVegs); i++ {
		if id == allVegs[i].Id {
			*reply = allVegs[i].AvailableQty
			return nil
		}
	}
	return fmt.Errorf("Not found vegetable.id =%v ", id)

}

func (v *Vegetable) UpdateVegetable(veg Vegetable, reply *Vegetable) error {
	mutex.Lock()
	defer mutex.Unlock()
	allVegs := Read()

	for i := 0; i < len(allVegs); i++ {
		if veg.Id == allVegs[i].Id {
			if veg.AvailableQty != 0 {
				allVegs[i].AvailableQty = veg.AvailableQty
			}
			if veg.UnitPrice != 0 {
				allVegs[i].UnitPrice = veg.UnitPrice
			}
			Write(allVegs)
			*reply = veg
			return nil
		}
	}
	return fmt.Errorf("Not found vegetable.id =%v ", veg.Id)

}

func NewVegetableFactory() *Vegetable {
	return &Vegetable{}
}
