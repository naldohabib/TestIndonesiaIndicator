package service

import (
	"TestScrapeCRUD/dataCrud"
	"TestScrapeCRUD/models"
	"errors"
)

// DataCrudServiceImpl ...
type DataCrudServiceImpl struct {
	dataCrud dataCrud.DataCrudRepo
}

func (d DataCrudServiceImpl) GetDataByID(id int) (*models.YoutubeData, error) {
	return d.dataCrud.GetDataByID(id)
}

func (d DataCrudServiceImpl) Delete(id int) error {
	err := d.dataCrud.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (d DataCrudServiceImpl) Update(id int, data *models.YoutubeData) (*models.YoutubeData, error) {
	_, err := d.dataCrud.GetDataByID(id)
	if err != nil {
		return nil, errors.New("ID does not exist")
	}

	user, err := d.Update(id, data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (d DataCrudServiceImpl) Insert(data *models.YoutubeData) (*models.YoutubeData, error) {
	if err := data.Validate(); err != nil {
		return nil, err
	}

	data, err := d.dataCrud.Insert(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d DataCrudServiceImpl) GetAll() (*[]models.YoutubeData, error) {
	data, err := d.dataCrud.GetAll()
	if err != nil {
		return nil, err
	}

	return data, nil
}

// CreateDataCrudServiceImpl ...
func CreateDataCrudServiceImpl(dataCrud dataCrud.DataCrudRepo) dataCrud.DataCrudService {
	return &DataCrudServiceImpl{dataCrud}
}
