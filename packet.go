package test_20190911

type Packet struct {

	msgLength int

	version uint16

	msgId uint32

	errorCode uint32

	msgData []byte

	hasReqId bool

	ReqId uint32

}
//
func NewPacket(msgId uint32,errCode uint32) (packet *Packet){

	packet=&Packet{

		msgId:msgId,
		errorCode:errCode,
		msgData:make([]byte,0),
		hasReqId:false,
	}

	return
}

