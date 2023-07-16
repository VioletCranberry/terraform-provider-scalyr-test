package client

import (
	"strconv"
	"time"
)

type ApiTime time.Time

func (tt *ApiTime) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*tt = ApiTime(time.Unix(0, millis*int64(time.Millisecond)))
	return nil
}

func (tt ApiTime) Time() time.Time {
	return time.Time(tt).UTC()
}

func (tt ApiTime) String() string {
	return tt.Time().String()
}
