package packet

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/wuqifei/chat/logs"
	"github.com/wuqifei/server_lib/libio"
)

type Packet struct {

	// 最大的接收字节数
	MaxRecvBufferSize uint16

	// 最大的发送字节数
	MaxSendBufferSize uint16
}

// NewPacket 新建一个packet
func NewPacket(maxRecvSize, maxSendSize uint16) *Packet {

	p := new(Packet)
	p.MaxRecvBufferSize = maxRecvSize
	p.MaxSendBufferSize = maxSendSize
	return p
}

func RollPacket(route uint16, p proto.Message) []byte {

	b, e := proto.Marshal(p)
	if e != nil {
		logs.Error("Roolpachet error [%v]", e)
		return nil
	}
	lb := make([]byte, 2)
	libio.PutUint16BE(lb, route)
	lb = append(lb, b...)
	return lb
}

func (p *Packet) Read(r *libio.Reader) ([]byte, error) {
	n := p.readHeader(r)

	if r.Error() != nil {
		return nil, r.Error()
	}
	//这里如果，字节数目过长
	if n > p.MaxRecvBufferSize {
		return nil, fmt.Errorf("recv too long")
	}

	if n == 0 {
		return nil, nil
	}

	b := r.ReadBytes(int(n))
	if r.Error() != nil {
		return nil, r.Error()
	}

	return b, nil
}

func (p *Packet) Write(w *libio.Writer, b []byte) error {
	length := len(b)
	if length > int(p.MaxSendBufferSize) {
		return fmt.Errorf("send too long")
	}
	logs.Debug("write msg :[%+v]", b)
	pt := make([]byte, 2)
	libio.PutUint16BE(pt, uint16(length))
	pt = append(pt, b...)

	w.Write(pt)
	return w.Error()
}

// 读取头，然后得到整个的长度
func (p *Packet) readHeader(r *libio.Reader) uint16 {
	// 读取两个字节长度
	return r.ReadUint16BE()
}
