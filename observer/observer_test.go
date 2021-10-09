package observer

import (
	"fmt"
	"github.com/Rmarken5/file-broadcaster/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)
//go:generate mockgen -destination=../mocks/mock_conn.go --package=mocks net Conn
func TestConnectionData_OnUpdate_Nil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConn(ctrl)
	m.EXPECT().Write(gomock.Any()).Return(11, nil)

	conData := ConnectionData{
		"addr",
		m,
	}

	err := conData.OnUpdate([]string{"file-1.txt"})
	if err != nil {
		t.FailNow()
	}
}

func TestConnectionData_OnUpdate_NotNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mocks.NewMockConn(ctrl)
	m.EXPECT().Write(gomock.Any()).Return(0, fmt.Errorf("SomeError"))

	conData := ConnectionData{
		"addr",
		m,
	}

	err := conData.OnUpdate([]string{"file-1.txt"})
	if err == nil {
		t.FailNow()
	}
}

func TestConnectionData_GetIdentifier(t *testing.T) {
	connData := ConnectionData{
		"hello",
		nil,
	}

	addr := connData.GetIdentifier()
	assert.EqualValues(t, "hello", addr)
}
