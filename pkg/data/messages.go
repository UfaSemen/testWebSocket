package data

//request from client to server with two numbers
type Request struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}

//response from server to client with sum and product
type Response struct {
	Sum     int `json:"sum"`
	Product int `json:"product"`
}

//client output files structure
type Result struct {
	Req  Request  `json:"req"`
	Resp Response `json:"resp"`
}
