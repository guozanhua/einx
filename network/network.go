package network

import (
	"github.com/Cyinx/einx/agent"
	"github.com/Cyinx/einx/component"
	"github.com/Cyinx/einx/module"
)

type Agent = agent.Agent
type AgentID = agent.AgentID
type ServerType uint32

type ModuleEventer = module.ModuleEventer

type ComponentID = component.ComponentID
type ComponentType = component.ComponentType
type Component = component.Component

func GenComponentID() ComponentID {
	return component.GenComponentID()
}

const (
	ServerType_TCP = iota
	ServerType_UDP
	ClientType_TCP
	ClientType_UDP
)

type ITcpServerCom interface {
	GetID() ComponentID
	GetType() ComponentType
}

type Linker interface {
	Ping()
	Pong()
}

type ITcpClientCom interface {
	GetID() ComponentID
	GetType() ComponentType
	Connect(addr string)
}

type ConnType uint16

const (
	ConnType_TCP = iota
	ConnType_UDP
)

type WriteWrapper struct {
	msg_type byte
	msg_id   ProtoTypeID
	buffer   []byte
}