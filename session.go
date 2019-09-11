package test_20190911

import (
	"github.com/golang-collections/go-datastructures/queue"
	"net"
	"sync"
)

//
type Session struct {

	conn net.Conn

	user *User

	state int

	writeQ	*queue.Queue

	handleConnected func()

	handleCloseConn func()

	isConnected bool

	protocolType int

	nextHeartBeatTime int64

	heartMutex *sync.Mutex

	readDataSizeAll int64

	writeDataSizeAll int64

	WriteBufferLen int

	ReadBufferLen  int

	StreamMode     bool

	Timeout        int

}
//
func (self *Session)Run(){


}
//
type User struct {


}

func (self *User)Run(){

}
