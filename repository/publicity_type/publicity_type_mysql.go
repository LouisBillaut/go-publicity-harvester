package publicity_type

import (
	"database/sql"
	"fmt"
	"github.com/louisbillaut/test/entity"
)

//PublicityTypeMySQL mysql repo
type PublicityTypeMySQL struct {
	db *sql.DB
}

//NewPublicityTypeMySQL create new repository
func NewPublicityTypeMySQL(db *sql.DB) *PublicityTypeMySQL {
	return &PublicityTypeMySQL{
		db: db,
	}
}

func (p *PublicityTypeMySQL) GetByPublicityID(id entity.ID) (*entity.PublicityType, error) {
	request, err := p.db.Prepare(`select publicity_id, type_id from publicity_type where publicity_id = ?`)
	if err != nil {
		return nil, err
	}
	var resultType entity.PublicityType
	rows, err := request.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&resultType.PublicityID, &resultType.TypeID)
		if err != nil {
			return nil, err
		}
	}
	if resultType.PublicityID.String() == "00000000-0000-0000-0000-000000000000" || resultType.TypeID.String() == "00000000-0000-0000-0000-000000000000" {
		return nil, entity.NotFound{Value: fmt.Sprintf("publicity_type (%s)", id.String())}
	}
	return &resultType, nil
}
