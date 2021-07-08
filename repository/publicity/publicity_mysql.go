package publicity

import (
	"database/sql"
	"fmt"
	"github.com/louisbillaut/test/entity"
	publicity_type2 "github.com/louisbillaut/test/repository/publicity_type"
	_type2 "github.com/louisbillaut/test/repository/type"
	"github.com/louisbillaut/test/usecase/publicity_type"
	_type "github.com/louisbillaut/test/usecase/type"
	"time"
)

//PublicityMySQL mysql repo
type PublicityMySQL struct {
	db *sql.DB
}

//NewPublicityMySQL create new repository
func NewPublicityMySQL(db *sql.DB) *PublicityMySQL {
	return &PublicityMySQL{
		db: db,
	}
}

//Create publicity
func (p *PublicityMySQL) Create(pub *entity.Publicity) (entity.ID, error) {
	//insert publicity
	request, err := p.db.Prepare(`
		insert into publicity (id, ua, ip, os, created_at) 
		values(?,?,?,?,?)`)
	if err != nil {
		return pub.ID, err
	}
	_, err = request.Exec(
		pub.ID,
		pub.Ua,
		pub.Ip,
		pub.Os,
		pub.CreatedAt,
		)
	if err != nil {
		return pub.ID, err
	}
	err = request.Close()
	if err != nil {
		return pub.ID, err
	}

	typeRepo := _type2.NewTypeMySQL(p.db)
	typeService := _type.NewService(typeRepo)
	//insert publicity types
	getType, err := typeService.GetTypeByName(pub.Type)
	if err != nil {
		return pub.ID, err
	}
	if "" == getType.Type {
		return pub.ID, entity.InvalidType{}
	}
	request, err = p.db.Prepare(`
		insert into publicity_type (publicity_id, type_id) 
		values(?,?)`)
	if err != nil {
		return pub.ID, err
	}
	_, err = request.Exec(pub.ID, getType.ID)
	if err != nil {
		return pub.ID, err
	}
	err = request.Close()
	if err != nil {
		return pub.ID, err
	}
	return pub.ID, nil
}

//GetByTimeStamp get publicity by Timestamp
func (p *PublicityMySQL) GetByTimeStamp(from time.Time, to time.Time) ([]*entity.Publicity, error) {
	request, err := p.db.Prepare(`select id, ua, ip, os, created_at from publicity where created_at>=? and created_at<?`)
	if err != nil {
		return nil, err
	}
	var res []*entity.Publicity
	rows, err := request.Query(from, to)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var pub entity.Publicity
		var createdAt string
		err = rows.Scan(&pub.ID, &pub.Ua, &pub.Ip, &pub.Os, &createdAt)
		date, err := time.Parse("2006-01-02", createdAt)
		if err != nil {
			return nil, err
		}
		pub.CreatedAt = date
		res = append(res, &pub)
	}
	return res, nil
}

//GetByOS get publicity by OS
func (p *PublicityMySQL) GetByOS(os string) ([]*entity.Publicity, error) {
	request, err := p.db.Prepare(`select id, ua, ip, os, created_at from publicity where os=?`)
	if err != nil {
		return nil, err
	}
	var res []*entity.Publicity
	rows, err := request.Query(os)
	if err != nil {
		return nil, err
	}

	publicityTypeRepo := publicity_type2.NewPublicityTypeMySQL(p.db)
	publicityTypeService := publicity_type.NewService(publicityTypeRepo)
	typeRepo := _type2.NewTypeMySQL(p.db)
	typeService := _type.NewService(typeRepo)
	for rows.Next() {
		var pub entity.Publicity
		err = rows.Scan(&pub.ID, &pub.Ua, &pub.Ip, &pub.Os, &pub.CreatedAt)
		res = append(res, &pub)

		publicityType, err := publicityTypeService.GetPublicityTypeByPublicityID(pub.ID)
		if err != nil {
			return nil, err
		}
		typeGet, err := typeService.GetTypeByID(publicityType.TypeID)
		if err != nil {
			return nil, err
		}
		pub.Type = typeGet.Type
	}
	return res, nil
}

//GetByID get publicity by ID
func (p *PublicityMySQL) GetByID(id entity.ID) (*entity.Publicity, error) {
	request, err := p.db.Prepare(`select id, ua, ip, os, created_at from publicity where id=?`)
	if err != nil {
		return nil, err
	}
	var pub entity.Publicity
	rows, err := request.Query(id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&pub.ID, &pub.Ua, &pub.Ip, &pub.Os, &pub.CreatedAt)
	}
	if pub.ID.String() == "00000000-0000-0000-0000-000000000000" {
		return nil, entity.NotFound{Value: fmt.Sprintf("publicity (%s)", id.String())}
	}
	publicityTypeRepo := publicity_type2.NewPublicityTypeMySQL(p.db)
	publicityTypeService := publicity_type.NewService(publicityTypeRepo)
	publicityType, err := publicityTypeService.GetPublicityTypeByPublicityID(pub.ID)
	if err != nil {
		return nil, err
	}
	typeRepo := _type2.NewTypeMySQL(p.db)
	typeService := _type.NewService(typeRepo)
	typeGet, err := typeService.GetTypeByID(publicityType.TypeID)
	if err != nil {
		return nil, err
	}
	pub.Type = typeGet.Type
	return &pub, nil
}
