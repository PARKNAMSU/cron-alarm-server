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
