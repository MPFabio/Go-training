package gigasecond

import "time"


func AddGigasecond(t time.Time) time.Time {
	gigasecond := time.Second * 1_000_000_000
	return t.Add(gigasecond)
}


