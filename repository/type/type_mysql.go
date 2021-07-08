package _type

import (
	"database/sql"
	"fmt"
	"github.com/louisbillaut/test/entity"
)

//TypeMySQL mysql repo
type TypeMySQL struct {
	db *sql.DB
}

//NewTypeMySQL create new repository
func NewTypeMySQL(db *sql.DB) *TypeMySQL {
	return &TypeMySQL{
		db: db,
	}
}

//Create type
func (t *TypeMySQL) Create(typ *entity.Type) (entity.ID, error) {
	_, err := t.GetByName(typ.Type)
	if err == nil {
		return typ.ID, entity.AlreadyExist{Value: typ.Type}
	}
	//insert type
	request, err := t.db.Prepare(`
		insert into type (id, type) 
		values(?,?)`)
	if err != nil {
		return typ.ID, err
	}
	_, err = request.Exec(
		typ.ID,
		typ.Type,
	)
	if err != nil {
		return typ.ID, err
	}
	err = request.Close()
	if err != nil {
		return typ.ID, err
	}
	return typ.ID, nil
}

func (t *TypeMySQL) GetByName(name string) (*entity.Type, error) {
	request, err := t.db.Prepare(`select id, type from type where type = ?`)
	if err != nil {
		return nil, err
	}
	var resultType entity.Type
	rows, err := request.Query(name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&resultType.ID, &resultType.Type)
		if err != nil {
			return nil, err
		}
	}
	if resultType.ID.String() == "00000000-0000-0000-0000-000000000000"{
		return nil, entity.NotFound{Value: fmt.Sprintf("type (%s)", name)}
	}
	return &resultType, nil
}

func (t *TypeMySQL) GetByID(id entity.ID) (*entity.Type, error) {
	request, err := t.db.Prepare(`select id, type from type where id = ?`)
	if err != nil {
		return nil, err
	}
	var resultType entity.Type
	rows, err := request.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&resultType.ID, &resultType.Type)
		if err != nil {
			return nil, err
		}
	}
	if resultType.ID.String() == "00000000-0000-0000-0000-000000000000"{
		return nil, entity.NotFound{Value: fmt.Sprintf("type (%s)", id.String())}
	}
	return &resultType, nil
}

