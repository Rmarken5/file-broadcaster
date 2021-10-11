package main

import (
	"github.com/Rmarken5/file-broadcaster/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)


func Test_AcceptClients(t *testing.T) {

	ctr := gomock.NewController(t)
	netConn := mocks.NewMockConn(ctr)
	netConn.EXPECT().RemoteAddr().AnyTimes().Return("8000")
	netListener := mocks.NewMockListener(ctr)
	netListener.EXPECT().Close().Return(nil)
	netListener.EXPECT().Accept().AnyTimes().Return(netConn,nil)
	fileListener := mocks.NewMockIFileListener(ctr)
	subject := mocks.NewMockSubject(ctr)
	server := server{
		FileListener: fileListener,
		FileSubject: subject,
	}

	server.acceptClients(netListener)
}
