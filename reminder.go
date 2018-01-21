package pillbox

type Reminder struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	Morning   bool   `json:"morning"`
	Afternoon bool   `json:"afternoon"`
	Evening   bool   `json:"evening"`
}
