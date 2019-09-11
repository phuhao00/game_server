package test_20190911

import (
	"log"
	"net"
)

type client struct {

	Id string

	DialServerAddr string

	Protocol int

	IsAutoReconnect	bool

	cookie *Cookie

	handleConnected func()

	handleConnClosed func()

}
//
type  NewClientParam struct {

	Id string

	DialServerAddr string

	Protocol int

	IsAutoReconnect	bool

}
//
func NewClient(param *NewClientParam)(clientNew *client)  {

	clientNew=&client{
		Id:              param.Id,
		DialServerAddr:  param.DialServerAddr,
		Protocol:        param.Protocol,
		IsAutoReconnect: param.IsAutoReconnect,
	}

	return
	
}
//
func (self *client) start() {

	tcpAddr, err := net.ResolveTCPAddr("tcp4", self.DialServerAddr)

	if err != nil {

		log.Printf("Fatal error: %s", err.Error())
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)

	newCookieParam:=&NewCookieParam{
		conn,
			1,
	}

	self.cookie=NewCookie(newCookieParam)

	self.cookie.Run()

}
//
func (self *client)stop() {


}
//
func (self *Cookie)sendMsg(){

}


