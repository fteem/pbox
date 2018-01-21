package pillbox

import (
	"time"
)

type Reminder struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	Morning   bool      `json:"morning"`
	Afternoon bool      `json:"afternoon"`
	Evening   bool      `json:"evening"`
	LastRan   time.Time `json:"last_ran"`
}

func (r Reminder) String() string {
	out := ""
	out += r.Body + " to be had at: "
	if r.Morning {
		out += "morning "
	}
	if r.Afternoon {
		out += "afternoon "
	}
	if r.Evening {
		out += "evening"
	}

	return out
}
