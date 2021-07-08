package entity

type PublicityType struct {
	PublicityID        ID
	TypeID ID
}

func NewPublicityType(publicityID ID, typeID ID) *PublicityType {
	return &PublicityType{
		PublicityID: publicityID,
		TypeID:      typeID,
	}
}