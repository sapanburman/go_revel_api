package utils

type SuccessRes struct {
	Status string `json:"status"`
	Code   int		`json:"code"`
	Data 	interface{}	`json:"data"`
	Message	string		`json:"message"`
}

type ErrorRes struct {
	Status string	`json:"status"`
	Code 	int	`json:"code"`
	Message	string	`json:"message"`
}