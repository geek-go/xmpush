package xmpush

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

//通过RegId群推
func SendMessageByRegIds(appSecret string, message *Message) (*Result, error) {

	if message.RegistrationId == "" {
		return nil, errors.New("registration_id is required")
	}

	form, _ := messageToForm(message)
	form.Add("registration_id", message.RegistrationId) //追加registration_id

	res, err := DoPost(MessageRegIdURL, appSecret, form)
	if err != nil {
		return nil, err
	}

	var result = &Result{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return result, err
	}

	return result, nil
}

//通过Alias群推
func SendMessageByRegAliasIds(appSecret string, message *Message) (*Result, error) {

	if message.Alias == "" {
		return nil, errors.New("alias is required")
	}

	form, _ := messageToForm(message)
	form.Add("alias", message.Alias) //追加alias

	res, err := DoPost(MessageAliasURL, appSecret, form)
	if err != nil {
		return nil, err
	}

	var result = &Result{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return result, err
	}

	return result, nil
}

//给所有人发送消息
func SendMessageAll(appSecret string, message *Message) (*Result, error) {

	form, _ := messageToForm(message)

	res, err := DoPost(MessageAllURL, appSecret, form)
	var result = &Result{}
	err = json.Unmarshal(res, result)
	if err != nil {
		return result, err
	}

	return result, nil
}

//消息转表单，小米推送接口使用form表单提交
func messageToForm(message *Message) (*url.Values, error) {

	form := &url.Values{}

	form.Add("restricted_package_name", message.RestrictedPackageName)
	form.Add("payload", message.Payload)
	form.Add("title", message.Title)
	form.Add("description", message.Description)
	form.Add("notify_type", fmt.Sprintf("%d", message.NotifyType))
	form.Add("pass_through", fmt.Sprintf("%d", message.PassThrough))

	if message.NotifyID > 0 {
		form.Add("notify_id", fmt.Sprintf("%d", message.NotifyID))
	}

	if message.TimeToLive > 0 {
		form.Add("time_to_live", fmt.Sprintf("%d", message.TimeToLive))
	}

	if message.TimeToSend > 0 {
		form.Add("time_to_send", fmt.Sprintf("%d", message.TimeToSend))
	}

	if message.Extra != nil {
		for k, v := range message.Extra {
			form.Add(fmt.Sprintf("extra.%s", k), v)
		}
	}

	return form, nil
}
