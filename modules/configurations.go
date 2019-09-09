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

type UpdateConfigurationBody struct {
	// StackID
	StackID uint64 `json:"stackID,string" default:""`
	// Key
	Key string `json:"key" default:""`
	// Value
	Value string `json:"value" default:""`
	// Remark
	Remark string `json:"remark" default:""`
}

func (req UpdateConfigurationBody) Validation(configID uint64) error {
	if configID == 0 && req.StackID == 0 && req.Key == "" {
		return errors.BadRequest.StatusError().WithMsg("configurationID stackID key 不能全为空").WithErrTalk()
	}
	if configID == 0 && ((req.StackID != 0 && req.Key == "") || (req.StackID == 0 && req.Key != "")) {
		return errors.BadRequest.StatusError().WithMsg("stackID key 不能单一为空").WithErrTalk()
	}
	return nil
}

func UpdateConfiguration(configID uint64, req UpdateConfigurationBody, db *sqlx.DB) error {
	err := req.Validation(configID)
	if err != nil {
		return err
	}
	model := &database.Configuration{
		Value:  req.Value,
		Remark: req.Remark,
	}
	if configID != 0 {
		model.ConfigurationID = configID
		err = model.UpdateByConfigurationIDWithStruct(db)
	} else {
		model.StackID = req.StackID
		model.Key = req.Key
		err = model.UpdateByStackIDAndKeyWithStruct(db)
	}
	return err
}
