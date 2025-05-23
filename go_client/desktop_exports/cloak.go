package main

import (
	"C"
	"go_client/cloak"
)

//export StartCloakClient
func StartCloakClient(localHost, localPort, config string, udp bool) {
	cloak.StartCloakClient(localHost, localPort, config, udp)
}

//export StopCloakClient
func StopCloakClient() {
	cloak.StopCloakClient()
}
