package sdk

type Reply struct {
	Code string
	Url  string
}

const (
	HaveNoRright      = "403"
	PostSuccess       = "200"
	DeleteSuccsess    = "201"
	ServerHandleError = "400"
)
