package kotlin_exports

import (
    "go_client/outline"
	"log"
	"fmt"
)

type OutlineClient struct {
	*outline.OutlineClient
}

func (c *OutlineClient) Connect() error {
	log.Println("func (c *OutlineClient) Connect() error")
	return c.OutlineClient.Connect()
}

func (c *OutlineClient) Disconnect() error {
	return c.OutlineClient.Disconnect()
}

//func (c *OutlineClient) GetServerIP() net.IP {
//	return c.OutlineClient.GetServerIP()
//}

func (c *OutlineClient) Read() ([]byte, error) {
	log.Println("func (c *OutlineClient) Read()")
	return c.OutlineClient.Read()
}

func (c *OutlineClient) Write(buf []byte) (int, error) {
	log.Println("func (c *OutlineClient) Write(buf []byte)")
	return c.OutlineClient.Write(buf)
}

func NewOutlineClient(transportConfig string) *OutlineClient {
	log.Println("Before func NewOutlineClient(transportConfig string) *OutlineClient")
	cl := outline.NewClient(transportConfig)
	log.Println("After func NewOutlineClient(transportConfig string) *OutlineClient")
	return &OutlineClient{OutlineClient: cl}
}
