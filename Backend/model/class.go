package model

import (
	"encoding/json"
	"strconv"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
}

func (t *Time) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0)
	return nil
}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
func (t Time) Unix() int64 {
	return time.Time(t).Unix()
}

// Time returns the JSON time as a time.Time instance in UTC
func (t Time) Time() time.Time {
	return time.Time(t).UTC()
}

// String returns t as a formatted string
func (t Time) String() string {
	return t.Time().String()
}

type Class struct {
	ClassName   string `json:"className"`
	TotalVotes  int    `json:"total_votes" default:"0"`
	LastUpdated Time   `json:"lastUpdated"`
}

type ClassWithoutTotalVotes struct {
	Class
}

func (c ClassWithoutTotalVotes) MarshalJSON() ([]byte, error) {
	type Alias struct {
		ClassName   string `json:"className"`
		LastUpdated Time   `json:"lastUpdated"`
	}
	return json.Marshal(Alias{ClassName: c.Class.ClassName, LastUpdated: c.Class.LastUpdated})
}
