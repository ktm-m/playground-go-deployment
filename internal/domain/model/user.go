package model

type GetUserByIDReq struct {
	ID string `json:"id"`
}

type GetUserByIDRespContent struct {
	ID          string `json:"id"`
	FirstnameTH string `json:"firstnameTH"`
	LastnameTH  string `json:"lastnameTH"`
	Age         int    `json:"age"`
}
