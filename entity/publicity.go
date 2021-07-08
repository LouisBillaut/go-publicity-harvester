package entity

import (
	useragent "github.com/mileusna/useragent"
	"time"
)

type Publicity struct {
	ID        ID
	Type     string
	Ua        string
	Ip        string
	CreatedAt time.Time
	Os        string
}

func NewPublicity(pubType string, ua string, ip string) *Publicity {
	e := &Publicity{
		ID:        NewID(),
		Ua:        ua,
		Ip:        ip,
		CreatedAt: time.Now().UTC(),
		Os:        useragent.Parse(ua).OS,
		Type:     pubType,
	}
	return e
}
