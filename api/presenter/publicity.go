package presenter

import "github.com/louisbillaut/test/entity"

type Publicity struct {
	Type string `json:"type"`
	Ua string `json:"ua"`
	Ip string `json:"ip"`
	Os string `json:"os"`
}

func ResponsePublicityCreated(pub *entity.Publicity) *Publicity {
	return &Publicity{
		Type: pub.Type,
		Ua: pub.Ua,
		Ip: pub.Ip,
		Os: pub.Os,
	}
}
