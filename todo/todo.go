package todo

import (
	"github.com/antlinker/go-mqtt/client"
	"github.com/antlinker/sdk/asapi"
	"github.com/antlinker/sdk/utils"
)

var (
	gHandle *Handle
)

// SetAuthorizeHandle 设置授权处理
func SetAuthorizeHandle(auh *asapi.AuthorizeHandle) {
	if gHandle == nil {
		gHandle = new(Handle)
	}
	gHandle.auh = auh
}

// MQTTConfig mqtt配置参数
type MQTTConfig struct {
	BrokerAddress     string
	ClientID          string
	Timeout           uint16
	CleanSession      bool
	ReconnectInterval int
	Username          string
	PasswordHandler   func() string
	EnableTLS         bool
}

// SetMQTTClient 设置MQTT客户端
func SetMQTTClient(mqcfg *MQTTConfig) {
	if gHandle == nil {
		gHandle = new(Handle)
	}

	cfg := &utils.MQTTConfig{
		BrokerAddress:     mqcfg.BrokerAddress,
		ClientID:          mqcfg.ClientID,
		Timeout:           mqcfg.Timeout,
		CleanSession:      mqcfg.CleanSession,
		ReconnectInterval: mqcfg.ReconnectInterval,
		Username:          mqcfg.Username,
		PasswordHandler:   mqcfg.PasswordHandler,
		EnableTLS:         mqcfg.EnableTLS,
	}

	cli, err := utils.NewMQTTClient(cfg)
	if err != nil {
		panic(err)
	}

	gHandle.mqcli = cli
}

// SetMQTTClienter 设置MQTT客户端
func SetMQTTClienter(cli client.MqttClienter) {
	gHandle.mqcli = cli
}

// Add 增加待办事项
func Add(req *AddRequest) (err error) {
	err = gHandle.Add(req)
	return
}

// Done 完成待办事项
func Done(req *DoneRequest) (err error) {
	err = gHandle.Done(req)
	return
}

// Del 删除待办事项
func Del(req *DelRequest) (err error) {
	err = gHandle.Del(req)
	return
}
