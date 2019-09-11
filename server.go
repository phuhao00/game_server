package test_20190911

import (
	"net"
	"sync"
)

//
type ServerInterface interface {

	start()

	stop()

	resolvePacket()

	handleConnected()

	handleConnClosed()

}
//
type WorldServer struct {

	ServerBase


}

//
type ServerBase struct {

	sessions sync.Map

	id string

	addr string

	listener net.Listener

	protocol int

}
//
func (self *ServerBase)start()  {

}
//
func (self *ServerBase)stop()  {

}
//
func (self *ServerBase)resolvePacket()  {

}
//
func (self *ServerBase)handleConnected()  {

}
//
func (self *ServerBase)handleConnClosed()  {

}
