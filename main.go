package main

import (
	"fmt"
	"log"
	"net/http"
	"gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Home Page")
}

func wsEndPoint(w http.ResponseWriter, r *http.Request){
	// fmt.Fprintf(w,"WebSocket EndPoint")
	upgrader.CheckOrigin = func(r *http.Request) bool{ return true}

	// upgrade this connection to a WebSocket
	// connection

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil{
		log.Println(er)
	}

	// helperful log statement to show connections
	log.Println("Client Connected")

	reader(ws)
}



func setupRoutes(){
	http.HandleFunc("/",homePage)
	http.HandleFunc("/ws",wsEndPoint)
}


// define a reader which will listen for 
// new messages being sent to our WebSocket
// endpoint
func reader(conn *websocket.Conn){
	for{
		// read in a message
		messageType, p ,err := conn.ReadMessage()
		if err != nil{
			log.Println(err)
			return
		}

		// print out that message for clarity
		fmt.Println(string(p))
		if err := conn.WriteMessage(messageType,p); err != nil{
			log.Println(err)
			return
		}
	}
}



func main(){
	fmt.Println("Go WebSockets")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":8080",nil))
}