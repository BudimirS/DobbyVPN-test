//go:build android || ios

package outline

import (
	"fmt"
	"go_client/common"
	"go_client/outline/internal"
	"log"
	"net"
	"unsafe"
	//_ "go_client/logger"
)

const Name = "outline"

type OutlineClient struct {
	device *internal.OutlineDevice
	config string
}

func NewClient(transportConfig string) *OutlineClient {
	c := &OutlineClient{config: transportConfig}
	log.Println("outline client created")
	common.Client.SetVpnClient(Name, c)
	return c
}

func (c *OutlineClient) Connect() error {
	od, err := internal.NewOutlineDevice(c.config)
	if err != nil {
		log.Printf("failed to create outline device: %v\n", err)
		return err
	}

	log.Println("outline device created")
	log.Println("outline client connected")

	c.device = od
	common.Client.MarkActive(Name)
	return nil
}

func (c *OutlineClient) Disconnect() error {
	err := c.device.Close()
	if err != nil {
		log.Printf("failed to close outline device: %v\n", err)
		return err
	}
	log.Println("outline client disconnected")
	common.Client.MarkInactive(Name)
	return nil
}

func (c *OutlineClient) Refresh() error {
	return c.device.Refresh()
}

func (c *OutlineClient) GetServerIP() net.IP {
	return c.device.GetServerIP()
}

func (c *OutlineClient) Read() ([]byte, error) {
	// Allocate a buffer with extra space to ensure 8-byte alignment.
	// The `bulkBarrierPreWrite: unaligned arguments` error suggests that the Go runtime's
	// garbage collector write barrier might be encountering an unaligned address
	// when processing the `buf` slice, especially when interacting with native code.
	// By ensuring 8-byte alignment, we mitigate potential issues with memory access
	// patterns expected by the underlying system or CGO calls.
	const bufferSize = 65536
	const alignment = 8
	alignedBuf := make([]byte, bufferSize+alignment-1)

	// Calculate the offset to achieve 8-byte alignment
	offset := 0
	if rem := uintptr(len(alignedBuf)) % alignment; rem != 0 {
		offset = alignment - int(rem)
	}
	buf := alignedBuf[offset : offset+bufferSize]

	log.Println(fmt.Sprintf("outline client: read data; before; size: %d (%d)", len(buf), uintptr(unsafe.Pointer(&buf[0]))%alignment))
	n, err := c.device.Read(buf)
	log.Println(fmt.Sprintf("outline client: read data; after; size: %d (%d)", n, uintptr(unsafe.Pointer(&buf[0]))%alignment))
	if err != nil {
		log.Printf("failed to read data: %v\n", err)
		return nil, fmt.Errorf("failed to read data: %w", err)
	}
	return buf, nil
}

func (c *OutlineClient) Write(buf []byte) (int, error) {
	log.Println(fmt.Sprintf("outline client: write data; before; size: %d (%d)", len(buf), len(buf)%8))
	n, err := c.device.Write(buf)
	log.Println(fmt.Sprintf("outline client: write data; after; size: %d (%d)", n, n%8))
	if err != nil {
		log.Printf("failed to write data: %v\n", err)
		return 0, fmt.Errorf("failed to write data: %w", err)
	}
	return n, nil
}
