package log_repository

type InsertLogUserAuthCodeInput struct {
	UserId         int
	ReceiveAccount string
	AuthType       string
	Code           string
	Action         string
	Status         int
	IpAddr         string
}

type InsertLogUserApiKeyInput struct {
	UserId   int
	IpAddr   string
	Hostname string
	ApiKey   string
	Action   string // issue, delete, periodExtend
}
