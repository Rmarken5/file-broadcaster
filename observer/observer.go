package observer

import (
	"fmt"
	"net"
	"strings"
)
//go:generate mockgen -destination=../mocks/mock_conn.go --package=mocks net Conn

type ConnectionData struct {
	Address string
	Conn    net.Conn
}

func (c *ConnectionData) OnUpdate(files []string) error {
	fileString := strings.Join(files, ",")

	fmt.Println("writing files: ", fileString)

	if _, err := c.Conn.Write([]byte(fileString)); err != nil {
		fmt.Printf("Unable to write %s to %s", fileString, c.Address)
		return fmt.Errorf("error %v: ", err)
	}
	fmt.Println("File String written.")
	return nil
}

func (c *ConnectionData) GetIdentifier() string {
	return c.Address
}
