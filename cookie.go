package test_20190911

import "net"

//
type Cookie struct {

	conn net.Conn



}

type NewCookieParam struct {

	conn net.Conn

	protocol int


}

func NewCookie( param *NewCookieParam)(cookieNew *Cookie)  {

	cookieNew=&Cookie{

		conn: param.conn,

	}

	return
}

func (self *Cookie)Run()  {

	go self.ReadPacket()

	go self.WritePacket()


}
//
func (self *Cookie)ReadPacket()  {

	
}
//
func (self *Cookie)WritePacket()  {


}
