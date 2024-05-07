package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

const wsEndpoint = "127.0.0.1:3000"

const sendInterval = time.Second

type OBUData struct {
	OBUID int     `json:"OBUID"`
	Lat   float64 `json:"Lat"`
	Long  float64 `json:"Long"`
}

// func sendOBUData(data OBUData) error {}

func genLatLong() (float64, float64) {
	return genCoord(), genCoord()
}

func genCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()
	return n + f
}

func main() {
	obuIDs := generateOBUIDs(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)
	for {
		for i := 0; i < len(obuIDs); i++ {
			lat, long := genLatLong()
			data := OBUData{
				OBUID: obuIDs[i],
				Lat:   lat,
				Long:  long,
			}
			fmt.Printf("%v\n", data)
		}
		fmt.Println(genCoord())
		time.Sleep(sendInterval)
	}
}

func generateOBUIDs(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
