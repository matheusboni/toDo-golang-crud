package repository

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
	"toDo-golang-crud/model"
)

type DBConnection struct {
	DB * gorm.DB
}

func (conn *DBConnection) Save(td *model.ToDo) (*model.ToDo, error) {

	(*td).Id = uuid.NewString()
	(*td).CreatedAt = time.Now()
	(*td).UpdatedAt = time.Now()

	err := conn.DB.Debug().Model(&model.ToDo{}).Create(&td).Error

	if err != nil {
		return nil, err
	}

	return td, nil
}

func (conn *DBConnection) GetById(id string) (*model.ToDo, error, bool) {

	var td *model.ToDo
	err := conn.DB.Debug().Model(model.ToDo{}).Where("id = ?", id).Take(&td).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New(fmt.Sprintf("ToDo with id: %s was not found", id)), true
	}

	if err != nil {
		return nil, err, false
	}

	return td, err, false
}

func (conn *DBConnection) GetALl() (*[]model.ToDo, error, bool) {

	var todos []model.ToDo
	err := conn.DB.Debug().Model(&model.ToDo{}).Limit(100).Find(&todos).Error

	if err != nil {
		return nil, err, false
	} else if len(todos) == 0 {
		return nil, errors.New("no toDo found"), true
	}

	return &todos, err, false
}

func (conn *DBConnection) Put(td *model.ToDo) (*model.ToDo, error) {

	conn.DB = conn.DB.Debug().Model(&model.ToDo{}).Where("id = ?", td.Id).Take(&model.ToDo{}).UpdateColumns(
		map[string]interface{}{
			"title":  td.Title,
			"description":  td.Description,
			"status":     td.Status,
			"updated_at": time.Now(),
		},
	)

	if conn.DB.Error != nil {
		return nil, conn.DB.Error
	}

	err := conn.DB.Debug().Model(&model.ToDo{}).Where("id = ?", td.Id).Take(&td).Error

	if err != nil {
		return nil, err
	}
	return td, nil
}

func (conn *DBConnection) Delete(id string) error {
	conn.DB = conn.DB.Debug().Model(&model.ToDo{}).Where("id = ?", id).Take(&model.ToDo{}).Delete(&model.ToDo{})

	if conn.DB.Error != nil {
		return conn.DB.Error
	}

	return nil
}
