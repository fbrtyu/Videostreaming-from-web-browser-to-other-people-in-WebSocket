package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Пропускаем любой запрос
	},
}

type Server struct {
	clients       map[*websocket.Conn]bool
	handleMessage func(message []byte) // хандлер новых сообщений
}

func StartServer(handleMessage func(message []byte)) *Server {
	server := Server{
		make(map[*websocket.Conn]bool),
		handleMessage,
	}

	http.HandleFunc("/ws", server.echo)
	go http.ListenAndServe(":8080", nil) // Уводим http сервер в горутину

	return &server
}

func (server *Server) echo(w http.ResponseWriter, r *http.Request) {
	connection, _ := upgrader.Upgrade(w, r, nil)
	defer connection.Close()

	server.clients[connection] = true        // Сохраняем соединение, используя его как ключ
	defer delete(server.clients, connection) // Удаляем соединение

	for {
		mt, message, err := connection.ReadMessage()

		if err != nil || mt == websocket.CloseMessage {
			break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
		}

		go server.handleMessage(message)
	}
}

func (server *Server) WriteMessage(message []byte) {
	for conn := range server.clients {
		conn.WriteMessage(websocket.TextMessage, message)
	}
}
