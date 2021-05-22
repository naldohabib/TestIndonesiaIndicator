package repository

import (
	"TestScrapeCRUD/dataCrud"
	"TestScrapeCRUD/models"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

// DataCrudRepoImpl ...
type DataCrudRepoImpl struct {
	db *gorm.DB
}

func (d DataCrudRepoImpl) GetDataByID(id int) (*models.YoutubeData, error) {
	dataUser := new(models.YoutubeData)

	if err := d.db.Table("tb_data").Where("id = ?", id).First(&dataUser).Error; err != nil {
		return nil, errors.New("ERROR: Error no data user with id you entered")
	}

	return dataUser, nil
}

func (d DataCrudRepoImpl) Delete(id int) error {
	data := models.YoutubeData{}

	err := d.db.Where("id=?", id).Delete(&data).Error
	if err != nil {
		return fmt.Errorf("[DataCrudRepoImpl.Delete] Error when query delete data with error: %w", err)
	}
	return nil
}

func (d DataCrudRepoImpl) Update(id int, data *models.YoutubeData) (*models.YoutubeData, error) {
	err := d.db.Model(&data).Where("id=?", id).Update(data).Error
	if err != nil {
		return nil, fmt.Errorf("DataCrudRepoImpl.Update Error when query update data with error: %w", err)
	}
	return data, nil
}

func (d DataCrudRepoImpl) Insert(data *models.YoutubeData) (*models.YoutubeData, error) {
	err := d.db.Save(&data).Error
		if err != nil {
			return nil, fmt.Errorf("[DataCrudRepoImpl.Insert] Error when query save data with: %w\n", err)
		}
		return data, nil
}

func (d DataCrudRepoImpl) GetAll() (*[]models.YoutubeData, error) {
	var data []models.YoutubeData
		err := d.db.Find(&data).Error
		if err != nil {
			return nil, fmt.Errorf("[DataCrudRepoImpl.GetAll] Error when query get all data with error: %w", err)
		}
		return &data, nil
}

// CreateDataCrudRepoImpl ...
func CreateDataCrudRepoImpl(db *gorm.DB) dataCrud.DataCrudRepo {
	return &DataCrudRepoImpl{db}
}
