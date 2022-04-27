package model

type Response struct {
	ResponseCode string `json:"response_code" example:"200"`
	ResultCode string `json:"result_code" example:"AS1001"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}
