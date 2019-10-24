package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"strconv"
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

func (t *House) GetHouse(args *Args, house *HouseDAO) error {
	log.Printf(strconv.Itoa(args.Id))
	for i:=0; i < len(houses); i++ {
		if houses[i].Id == args.Id {
			*house = houses[i]
		}
	}
	log.Printf(house.Name)
	return nil
}

// func main() {
// 	l, e := net.Listen("tcp", ":1234")

// 	house := new(House)
// 	err := rpc.Register(house)
// 	rpc.HandleHTTP()

// 	if e != nil {
// 		log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
// 	}
// 	log.Println("Serving RPC handler")
// 	err = http.Serve(l, nil)
// 	if err != nil {
// 		log.Fatalf("Error serving: %s", err)
// 	}
// }

func main() {
	server := rpc.NewServer()
	house := new(House)
	server.Register(house)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		server.ServeHTTP(w, r)
	})

	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("Couldn't start listening on port 1234. Error %s", e)
	}

	log.Println("Serving RPC handler")
	for {
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }
        go server.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}

// https://stackoverflow.com/questions/36610140/call-golang-call-jsonrpc-with-curl