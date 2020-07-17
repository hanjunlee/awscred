package core

import "time"

type (
	// Cache -
	Cache struct {
		AccessKeyID     string
		SecretAccessKey string
		SessionToken    string
		Expiration      time.Time
	}

	// Config -
	Config struct {
		On             string
		SerialNumber   string
		DurationSecond int64

		Cache Cache
	}
)