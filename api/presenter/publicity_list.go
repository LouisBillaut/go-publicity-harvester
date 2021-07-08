package presenter

import "github.com/louisbillaut/test/entity"

type PublicityList struct {
	PublicityList []Publicity `json:"publicityList"`
}

func ResponsePublicityList(pubs []*entity.Publicity) *PublicityList {
	var res []Publicity
	for _, p := range pubs {
		pub := Publicity{
			Type: p.Type,
			Ua:    p.Ua,
			Ip:    p.Ip,
			Os:    p.Os,
		}
		res = append(res, pub)
	}
	return &PublicityList{PublicityList: res}
}

