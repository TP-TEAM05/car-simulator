package main

import (
	"fmt"
	"net"
)

type ConnectionsManager struct {
	ServerAddress *net.UDPAddr
	Connections   map[string]*Connection
}

func NewConnectionsManager(serverAddress *net.UDPAddr) *ConnectionsManager {
	return &ConnectionsManager{
		ServerAddress: serverAddress,
		Connections:   make(map[string]*Connection),
	}
}

func (manager *ConnectionsManager) GetOrCreateConnection(id string) *Connection {

	connection, ok := manager.Connections[id]
	if !ok {
		udpConnection, err := net.DialUDP("udp", nil, manager.ServerAddress)
		if err != nil {
			fmt.Println("Could not connect to " + manager.ServerAddress.String() + " - " + err.Error())
			return nil
		}
		fmt.Println("New Connection " + id + " connected to " + manager.ServerAddress.String())

		connection = &Connection{
			Id:            id,
			UDPConnection: udpConnection,
			OtherAddress:  manager.ServerAddress,
			NextSendIndex: 1,
		}

		manager.Connections[id] = connection
	}
	return connection
}

func (manager *ConnectionsManager) DeleteConnectionById(id string) {
	delete(manager.Connections, id)
	fmt.Println("Connection " + id + " disconnected from " + manager.ServerAddress.String())
}

func (manager *ConnectionsManager) DeleteConnection(connection *Connection) {
	manager.DeleteConnectionById(connection.Id)
}

func (manager *ConnectionsManager) DeleteAllConnectionsExcept(ids []string) {
	for id := range manager.Connections {
		if !contains(ids, id) {
			manager.DeleteConnectionById(id)
		}
	}
}
