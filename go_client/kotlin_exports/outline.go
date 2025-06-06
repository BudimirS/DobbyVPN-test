package kotlin_exports

import (
    "go_client/outline"
	"log"
)

type OutlineClient struct {
	*outline.OutlineClient
}

func (c *OutlineClient) Connect() error {
	log.Println("start Connect()")
	err := c.OutlineClient.Connect()
	log.Println("end Connect()")
	return err
}

func (c *OutlineClient) Disconnect() error {
	log.Println("start Disconnect()")
	err := c.OutlineClient.Disconnect()
	log.Println("end Disconnect()")
	return err
}

//func (c *OutlineClient) GetServerIP() net.IP {
//	return c.OutlineClient.GetServerIP()
//}

func (c *OutlineClient) Read() ([]byte, error) {
	log.Println("start Read()")
	buf, err := c.OutlineClient.Read()
	log.Println("end Read()")
	return buf, err
}

func (c *OutlineClient) Write(buf []byte) (int, error) {
	log.Println("start Write(buf []byte)")
	n, err := c.OutlineClient.Write(buf)
	log.Println("end Write(buf []byte)")
	return n, err
}

func NewOutlineClient(transportConfig string) *OutlineClient {
	log.Println("start func NewOutlineClient(transportConfig string)")
	cl := outline.NewClient(transportConfig)
	res := &OutlineClient{OutlineClient: cl}
	log.Println("end func NewOutlineClient(transportConfig string)")
	return res
}
