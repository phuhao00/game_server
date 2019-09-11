package test_20190911

import (
	"bytes"
	"encoding/binary"
	"errors"
	"zinx/itface"
	"zinx/utils"
)

type Packet struct {

	msgLength uint32

	version uint16

	msgId uint32

	errorCode uint32

	msgData []byte

	hasReqId bool

	ReqId uint32

}
//
func NewPacket(msgId uint32,errCode uint32,msg []byte) (packet *Packet){

	packet=&Packet{

		msgId:msgId,
		errorCode:errCode,
		msgData:msg,
		hasReqId:false,
	}

	data,err:=packet.Pack()

	if err==nil {

		packet.msgData=data
	}

	return
}
//封包方法(压缩数据)
func(self *Packet) Pack()([]byte, error) {
	//创建一个存放bytes字节的缓冲
	dataBuff := bytes.NewBuffer([]byte{})

	//写dataLen
	if err := binary.Write(dataBuff, binary.LittleEndian, self.msgLength); err != nil {
		return nil, err
	}

	//写msgID
	if err := binary.Write(dataBuff, binary.LittleEndian, self.msgId); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(dataBuff, binary.LittleEndian, self.msgData); err != nil {
		return nil ,err
	}

	return dataBuff.Bytes(), nil
}

//拆包方法(解压数据)
func(self *Packet) Unpack(binaryData []byte)(itface.IMessage, error) {
	//创建一个从输入二进制数据的ioReader
	dataBuff := bytes.NewReader(binaryData)

	//读dataLen
	if err := binary.Read(dataBuff, binary.LittleEndian, &self.msgLength); err != nil {
		return nil, err
	}

	//读msgID
	if err := binary.Read(dataBuff, binary.LittleEndian, &self.msgId); err != nil {
		return nil, err
	}

	//判断dataLen的长度是否超出我们允许的最大包长度
	if (utils.GlobalObject.MaxPacketSize > 0 && self.msgLength > utils.GlobalObject.MaxPacketSize) {
		return nil, errors.New("Too large msg data recieved")
	}

	//这里只需要把head的数据拆包出来就可以了，然后再通过head的长度，再从conn读取一次数据
	return msg, nil
}



