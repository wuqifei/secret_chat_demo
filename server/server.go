package server

import (
	"fmt"

	"github.com/wuqifei/chat/logs"
	"github.com/wuqifei/chat/packet"
	"github.com/wuqifei/server_lib/libnet2"
)

var srv *Server

type Server struct {
	s       libnet2.LibserverInterface
	address string
}

func New(address string) {
	srv = new(Server)
	srv.address = address

	srv.new()
}

func (s *Server) new() {
	netOption := libnet2.DefaultOption()
	netOption.Address = s.address
	sessionOption := libnet2.DefaultSessionOption()
	server, err := libnet2.NewWithOption(netOption, sessionOption)
	if err != nil {
		panic(err)
	}
	s.s = server
	server.Run()
	libnet2.ServerPacket = packet.NewPacket(60000, 60000)
	libnet2.ServerSessionBlock = s.recvSession
	libnet2.ServerErrorBlock = s.errBlock
	libnet2.SessionCloseBlock = s.sessClosed
	libnet2.SessionErrorBlock = s.sessErrBlock
	libnet2.SessionRecvBlock = s.sessRecvMsg
}

func (s *Server) recvSession(sess libnet2.Session2Interface) {
	logs.Debug("[%d]recvSession", sess.GetUniqueID())
	// 生成一个64位的随机码
	ConfirmS2C(sess.GetUniqueID(), sess)
}

func (s *Server) errBlock(err error) {
	fmt.Printf("error of block [%v]\n", err)
}

func (s *Server) sessClosed(sess libnet2.Session2Interface) {
	fmt.Printf("session [%d]closed \n", sess.GetUniqueID())
}

func (s *Server) sessErrBlock(sess libnet2.Session2Interface, err error) {

	fmt.Printf("session [%d] error of block [%v]\n", sess.GetUniqueID(), err)
}

func (s *Server) sessRecvMsg(sess libnet2.Session2Interface, val []byte) {
	Handle(sess, val)
}
