package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/FulgurCode/school-erp-api/helpers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Message struct {
	TargetID string                 `json:"targetId"`
	Name     string                 `json:"name"`
	Message  map[string]interface{} `json:"message"`
}

var connections = map[string]*websocket.Conn{}
var connectionsM sync.Mutex

func HandleWebSocket(c *gin.Context) {
	var upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// Allow all connections
			return true
		},
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// Handle error
		return
	}

	var id = helpers.GenerateUniqueID()
	connectionsM.Lock()
	connections[id] = conn
	connectionsM.Unlock()

	conn.WriteJSON(map[string]interface{}{"name": "id", "data": id})

	fmt.Println(connections)


	for {
		// Read message from client
		_, msg, err := conn.ReadMessage()
		if err != nil {
			// Handle error or connection close
			break
		}

		var data Message
		json.Unmarshal(msg, &data)
		fmt.Println(data)

		targetID := data.TargetID

		connectionsM.Lock()
		targetConn, exists := connections[targetID]
		connectionsM.Unlock()

		if exists {
			err := targetConn.WriteJSON(map[string]interface{}{"name": data.Name, "data": data.Message["image"]})
			if err != nil {
				break
			}

		}

	}

	conn.Close()
	// Remove the connection from the active connections map
	connectionsM.Lock()
	delete(connections, id)
	connectionsM.Unlock()

	fmt.Println(connections)
}
