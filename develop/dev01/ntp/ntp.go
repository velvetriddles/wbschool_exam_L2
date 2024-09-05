package ntp

import (
	"time"

	"github.com/beevik/ntp"
)

const ADDRESS = "0.beevik-ntp.pool.ntp.org"

func GetTime(address string) (time.Time, error) {
	if address == "" {
		address = ADDRESS
	}
	curTime, err := ntp.Time(address)
	if err != nil {
		return time.Time{}, err
	}
	return curTime, nil
}
