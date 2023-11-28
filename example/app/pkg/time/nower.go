package time

import "time"

type Nower interface {
	Now() time.Time
}
