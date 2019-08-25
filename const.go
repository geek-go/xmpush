package xmpush

const (
	SandboxHost    = "https://sandbox.xmpush.xiaomi.com"
	ProductionHost = "https://api.xmpush.xiaomi.com"

	MessageRegIdURL   = "/v3/message/regid"
	MessageAliasURL   = "/v3/message/alias"
	MessageAllURL     = "/v3/message/all"
	MessageAccountURL = "/v2/message/user_account"
	MessageTopicURL   = "/v3/message/topic"
	MessageTopicOpURL = "/v3/message/multi_topic"

	MultiRegIdURL   = "/v2/multi_messages/regids"
	MultiAliasURL   = "/v2/multi_messages/aliases"
	MultiAccountURL = "/v2/multi_messages/user_accounts"

	StatsURL          = "/v1/stats/message/counters"
	MessageStatusURL  = "/v1/trace/message/status"
	MessagesStatusURL = "/v1/trace/messages/status"

	SubscribeURL   = "/v2/topic/subscribe"
	UnsubscribeURL = "/v2/topic/unsubscribe"

	SubscribeAliasURL   = "/v2/topic/subscribe/alias"
	UnsubscribeAliasURL = "/v2/topic/unsubscribe/alias"

	InvalidRegIdsURL = "https://feedback.xmpush.xiaomi.com/v1/feedback/fetch_invalid_regids"

	RegIdAliasURL = "/v1/alias/all"
	RegIdTopicURL = "/v1/topic/all"

	ScheduleJobExistURL  = "/v2/schedule_job/exist"
	ScheduleJobDeleteURL = "/v2/schedule_job/delete"
)
