package reddit

import "time"

const (
	botUserAgent = "you-dont-have-to"
)

var (
	botScriptRateLimit = time.Duration(3) * time.Second
)
