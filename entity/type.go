package entity

type Type struct {
	ID        ID
	Type string
}

func NewType(newType string) (*Type, error) {
	if newType == "" {
		return nil, InvalidArgument{ErrorString: "you must give a name to the new type"}
	}
	t := &Type{
		ID:        NewID(),
		Type: newType,
	}
	return t, nil
}
