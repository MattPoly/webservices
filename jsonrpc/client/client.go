package main

import (
	"log"
	"net/rpc/jsonrpc"
)

type HouseDAO struct {
	Id         int
	Name       string
	Region     string
	CoatOfArms string
	Words      string
}

var houses = []HouseDAO{
	HouseDAO{
		Id:         1,
		Name:       "House Algood",
		Region:     "The Westerlands",
		CoatOfArms: "A golden wreath, on a blue field with a gold border(Azure, a garland of laurel within a bordure or)",
		Words:      "",
	},
	HouseDAO{
		Id:         2,
		Name:       "House Allyrion of Godsgrace",
		Region:     "Dorne",
		CoatOfArms: "Gyronny Gules and Sable, a hand couped Or",
		Words:      "No Foe May Pass",
	},
	HouseDAO{
		Id:         3,
		Name:       "House Amber",
		Region:     "The North",
		CoatOfArms: "",
		Words:      "",
	},
}

type Args struct {
	Id int
}

type House int

func main() {
	client, err := jsonrpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	
	args := Args{1}
	var reply HouseDAO
	err = client.Call("House.GetHouse", args, &reply)
	if err != nil {
		log.Fatal("house error:", err)
	}
	log.Printf("House: %d", args.Id)
}