package dataCrud

import (
	"TestScrapeCRUD/models"
)

type DataCrudRepo interface {
	Insert(data *models.YoutubeData) (*models.YoutubeData, error)
	GetAll() (*[]models.YoutubeData, error)
	GetDataByID(id int) (*models.YoutubeData, error)
	//GetAllBookDetail(id int) (*[]model.BookDetail, error)
	Delete(id int) error
	Update(id int, data *models.YoutubeData) (*models.YoutubeData, error)
}