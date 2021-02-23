package utilities

import "github.com/segmentio/ksuid"

func GenerateRandomString() string {
	id := ksuid.New()
	return id.String()
}
