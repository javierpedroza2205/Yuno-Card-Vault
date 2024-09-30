package utils


type ErrorReason struct {
	Type        string `json:"type"`
	StatusCode  int    `json:"status_code"`
	Message     string `json:"message"`
	UserMessage string `json:"user_message"`
	Code        string `json:"code"`
}

type Error struct {
	Status string      `json:"status"`
	Error  ErrorReason `json:"error"`
}