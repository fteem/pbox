package pillbox

import "time"

type Checkpoint struct {
	Hour   int
	Minute int
}

func (c Checkpoint) Equal(t time.Time) bool {
	return c.Hour == t.Hour() && c.Minute == t.Minute()
}
