package model

type Response struct {
	Code string
	Status string
	Msg string
	Reason string
	Data interface{}
}
