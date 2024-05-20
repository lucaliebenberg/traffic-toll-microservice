package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/lucaliebenberg/tolling/types"
)

const wsEndpoint = "ws://127.0.0.1:3000/ws"

const sendInterval = time.Second * 5

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
	// fmt.Println(conn)
	for {
		for i := 0; i < len(obuIDs); i++ {
			lat, long := genLatLong()
			data := types.OBUData{
				OBUID: obuIDs[i],
				Lat:   lat,
				Long:  long,
			}
			// fmt.Printf("%v\n", data)
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		// fmt.Println(genCoord())
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
