package chat

import (
	"encoding/json"
	"fmt"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/liangdas/armyant/task"
	"github.com/liangdas/armyant/work"
	"math/rand"
	"strconv"
)

type Work struct {
	work.MqttWork
	manager *Manager
}

func NewWork(manager *Manager) task.Work {
	this := new(Work)
	this.manager = manager

	opts := this.GetDefaultOptions("ws://127.0.0.1:3653")
	opts.SetClientID(strconv.FormatInt(rand.Int63(), 10))
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		fmt.Println("ConnectionLost", err.Error())
	})

	opts.SetOnConnectHandler(func(client MQTT.Client) {
		fmt.Println("OnConnectHandler")
	})

	err := this.Connect(opts)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 主动监听 并且 定义 主体内容   只有客户端 主动监听 设置接口名称时，是topic，如果是服务器端则是id
	this.On("/chat/face", func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("%80s%v\n", "", string(msg.Payload()))
	})
	return this
}

func (w *Work) RunWorker(t task.Task) {
	var userID string
	fmt.Println("请输入要发送的用户ID")
	fmt.Scanln(&userID)
	fmt.Println("已建立通信，可以发送消息了")
	for {
		var data string
		fmt.Scanln(&data)
		str := fmt.Sprintf("{\"userID\":\"%v\", \"data\":\"%v\"}", userID, data)
		w.RequestNR("gate/HD_chat/face", []byte(str))
	}
}

func (w *Work) UnmarshalResult(payload []byte) map[string]interface{} {
	rmsg := map[string]interface{}{}
	json.Unmarshal(payload, &rmsg)
	return rmsg["Result"].(map[string]interface{})
}

func (w *Work) Init(t task.Task) {

}

func (w *Work) Close(t task.Task) {
	w.GetClient().Disconnect(0)
}
