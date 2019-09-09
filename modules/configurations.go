package modules

import (
	"github.com/johnnyeven/libtools/clients/client_id"
	"github.com/johnnyeven/libtools/helper"
	"github.com/johnnyeven/libtools/sqlx"
	"github.com/johnnyeven/libtools/sqlx/builder"
	"github.com/johnnyeven/service-configurations/constants/errors"
	"github.com/johnnyeven/service-configurations/database"
)

type CreateConfigurationBody struct {
	// StackID
	StackID uint64 `json:"stackID,string"`
	// Key
	Key string `json:"key"`
	// Value
	Value string `json:"value"`
	// Remark
	Remark string `json:"remark" default:""`
}

func CreateConfiguration(req CreateConfigurationBody, db *sqlx.DB, clientID client_id.ClientIDInterface) error {
	id, err := helper.NewUniqueID(clientID)
	if err != nil {
		return err
	}
	model := &database.Configuration{
		ConfigurationID: id,
		StackID:         req.StackID,
		Key:             req.Key,
		Value:           req.Value,
	}

	return model.Create(db)
}

func FetchConfiguration(stackID uint64, size, offset int32, db *sqlx.DB) (result database.ConfigurationList, count int32, err error) {
	model := &database.Configuration{}
	if stackID == 0 {
		err = errors.BadRequest
		return
	}

	condition := builder.And(model.T().F("StackID").Eq(stackID))
	result, count, err = model.FetchList(db, size, offset, condition)
	return
}
