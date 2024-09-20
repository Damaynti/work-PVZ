package postgres

import (
	"context"

	"example.com/mymodule/internal/model"
	"example.com/mymodule/internal/pkg/db"
)

type PVZRepo struct {
	db *db.Database
}

func NewPVZ(database *db.Database) *PVZRepo {
	return &PVZRepo{db: database}
}

func (r *PVZRepo) Add(ctx context.Context,pvz *model.PVZInput) (int64, error) {
	
	var id int64

	err:=r.db.ExecQueryRow(ctx,`INSERT INTO pvz(title,address,contactInformation) VALUES ($1,$2,$3) RETURNING id;`,pvz.Title,pvz.Address,pvz.ContactInformation).Scan(&id)
	return id,err
}


func (r *PVZRepo) GetAllPVZ(ctx context.Context) ([]model.PVZ, error) {
    var pvzList []model.PVZ
	raws, err := r.db.Query(ctx, "SELECT id, title, address, contactinformation, isDel FROM pvz")
	if err != nil {
        return nil, err
    }
    defer raws.Close()
	for raws.Next() {
		var pvz model.PVZRead
        if err := raws.Scan(&pvz.ID, &pvz.Title, &pvz.Address,&pvz.ContactInformation,&pvz.IsDel); err != nil {
            return nil, err
        }
		if pvz.IsDel==false{
			pvzList = append(pvzList, model.PVZ{Title: pvz.Title,Address: pvz.Address,ContactInformation: pvz.ContactInformation,ID: pvz.ID})
		}
    }
    if err := raws.Err(); err != nil {
		return nil, err
    }
    return pvzList, nil
}

func (r *PVZRepo) DeletePVZ(ctx context.Context, id int64) error {
	_,err:=r.db.Exec(ctx,`UPDATE pvz SET isDel = $1 WHERE id = $2;`,true,id)
	return err
}
