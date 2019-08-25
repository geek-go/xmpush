package xmpush

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

var Cfg = &XmpushConfig{
	AppSecret: "",
	Package:   "",
}

type XmpushConfig struct {
	AppSecret string `toml:"app_secret"`
	Package   string `toml:"package"`
}

//消息payload，根据业务自定义
type Payload struct {
	PushTitle    string `json:"push_title"`
	PushBody     string `json:"push_body"`
	IsShowNotify string `json:"is_show_notify"`
	Ext          string `json:"ext"`
}

//测试单推
func TestXmPush_SendByCid(t *testing.T) {
	iXmPush, err := NewXmPush(Cfg)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	cid := "xxx"
	payLoad := Payload{"这是测试title", "这是测试内容", "1", ""}
	err = iXmPush.SendByCid(cid, &payLoad)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}

//测试群推
func TestXmPush_SendByCids(t *testing.T) {
	iXmPush, err := NewXmPush(Cfg)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	cids := []string{"xxx"}
	payLoad := Payload{"这是测试title", "这是测试内容", "1", ""}
	err = iXmPush.SendByCids(cids, &payLoad)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}

//测试全推
func TestXmPush_SendAll(t *testing.T) {
	iXmPush, err := NewXmPush(Cfg)
	if err != nil {
		t.Error(err)
		os.Exit(1)
	}

	payLoad := Payload{"这是测试title", "这是测试内容", "1", ""}
	err = iXmPush.SendAll(&payLoad)
	if err != nil {
		t.Error(err)
	} else {
		t.Log("ok")
	}
}

type XMPush struct {
	Config *XmpushConfig
}

//获取实例
func NewXmPush(config *XmpushConfig) (*XMPush, error) {

	if config.Package == "" || config.AppSecret == "" {
		return nil, errors.New("请检查配置")
	}

	xm := &XMPush{
		Config: config,
	}

	return xm, nil
}

//根据用户cid推送
func (m *XMPush) SendByCid(cid string, payload *Payload) error {
	return m.SendByCids([]string{cid}, payload)
}

//根据用户cids批量推送
func (m *XMPush) SendByCids(cids []string, payload *Payload) error {

	payload_str, _ := json.Marshal(payload)

	//是否透传
	passThrough := 1
	if payload.IsShowNotify == "1" {
		passThrough = 0 //通知栏消息
	}

	message := &Message{
		Payload:               string(payload_str),
		Title:                 payload.PushTitle,
		Description:           payload.PushBody,
		PassThrough:           int32(passThrough),
		NotifyType:            1,
		RestrictedPackageName: m.Config.Package,
		Extra: map[string]string{
			"notify_foreground": "1",
		},
	}

	message.RegistrationId = strings.Join(cids, ",")

	result, err := SendMessageByRegIds(m.Config.AppSecret, message)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

//根据别名推送批量推送
func (m *XMPush) SendByAliases(aliases []string, payload *Payload) error {

	payload_str, _ := json.Marshal(payload)

	//是否透传
	passThrough := 1
	if payload.IsShowNotify == "1" {
		passThrough = 0 //通知栏消息
	}

	message := &Message{
		Payload:               string(payload_str),
		Title:                 payload.PushTitle,
		Description:           payload.PushBody,
		PassThrough:           int32(passThrough),
		NotifyType:            1,
		RestrictedPackageName: m.Config.Package,
		Extra: map[string]string{
			"notify_foreground": "1",
		},
	}

	message.Alias = strings.Join(aliases, ",")

	result, err := SendMessageByRegAliasIds(m.Config.AppSecret, message)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}

//推送给所有人
func (m *XMPush) SendAll(payload *Payload) error {

	payload_str, _ := json.Marshal(payload)

	//是否透传
	passThrough := 1
	if payload.IsShowNotify == "1" {
		passThrough = 0 //通知栏消息
	}

	message := &Message{
		Payload:               string(payload_str),
		Title:                 payload.PushTitle,
		Description:           payload.PushBody,
		PassThrough:           int32(passThrough),
		NotifyType:            1,
		RestrictedPackageName: m.Config.Package,
		Extra: map[string]string{
			"notify_foreground": "1",
		},
	}

	result, err := SendMessageAll(m.Config.AppSecret, message)
	if err != nil {
		return err
	}

	fmt.Println(result)

	return nil
}
