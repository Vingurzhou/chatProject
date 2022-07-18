package ctrl

import (
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gopkg.in/fatih/set.v0"
)

var rwlocker sync.RWMutex
var clientMap map[int64]*Node = make(map[int64]*Node, 0)
var udpsendchan chan []byte = make(chan []byte, 1024)

type Node struct {
	Conn *websocket.Conn
	//并行转串行,
	DataQueue chan []byte
	GroupSets set.Interface
}

func Chat(c *gin.Context) {
	id := c.Query("id")
	token := c.Query("token")
	userId, _ := strconv.ParseInt(id, 10, 64)
	user := userService.Find(userId)
	isvalida := user.Token == token

	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return isvalida
		},
	}).Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte, 50),
		GroupSets: set.New(set.ThreadSafe),
	}
	comIds := contactService.SearchComunityIds(userId)
	for _, v := range comIds {
		node.GroupSets.Add(v)
	}
	//todo userid和node形成绑定关系
	rwlocker.Lock()
	clientMap[userId] = node
	rwlocker.Unlock()
	//todo 完成发送逻辑,con
	go sendproc(node)
	//todo 完成接收逻辑
	go recvproc(node)
	log.Printf("<-%d\n", userId)
	sendMsg(userId, []byte("hello,world!"))
}

//ws发送协程
func sendproc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

//ws接收协程
func recvproc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			log.Println(err.Error())
			return
		}
		//dispatch(data)
		//把消息广播到局域网
		broadMsg(data)
		log.Printf("[ws]<=%s\n", data)
	}
}

//todo 发送消息
func sendMsg(userId int64, msg []byte) {
	rwlocker.RLock()
	node, ok := clientMap[userId]
	rwlocker.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
func broadMsg(data []byte) {
	udpsendchan <- data
}
