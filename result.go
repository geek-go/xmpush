package xmpush

//小米推送接口返回的结果，主体结构是一致的
type Result struct {
	Code        int64       `json:"code"`                  //0表示成功，非0表示失败
	Result      string      `json:"result"`                //"ok" 表示成功,"error" 表示失败
	Description string      `json:"description,omitempty"` //对发送消息失败原因的解释
	Info        string      `json:"info,omitempty"`        //详细信息
	Reason      string      `json:"reason,omitempty"`      //失败原因
	Data        *ResultData `json:"data,omitempty"`        //本身就是一个json字符串
}

type ResultData struct {
	BadRegids string `json:"bad_regids"` //推送失败的cids
	Id        string `json:"id"`         //消息的Id
}
