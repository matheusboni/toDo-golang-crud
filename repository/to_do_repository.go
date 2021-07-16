package repository

import (
	"github.com/google/uuid"
	"toDo-golang-crud/model"
)

var _db = map[string]model.ToDo{}

func Save(toDo *model.ToDo) *model.ToDo {
	id := uuid.NewString()
	(*toDo).Id = id

	_db[id] = *toDo

	return toDo
}

func Get(id string) *model.ToDo {
	val := _db[id]
	return &val
}

func GetALl() []model.ToDo {
	toDos := make([]model.ToDo, 0, len(_db))

	for  _, value := range _db {
		toDos = append(toDos, value)
	}

	return toDos
}

func Put(toDo *model.ToDo) *model.ToDo {
	_db[(*toDo).Id] = *toDo

	return toDo
}

func Delete(id string) {
	delete(_db, id)
}

