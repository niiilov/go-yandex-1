package models

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result"`
}

type Errors struct {
	Error string `json:"error"`
}
