package handler

import (
	"TestScrapeCRUD/dataCrud"
	"TestScrapeCRUD/models"
	"TestScrapeCRUD/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// DataCrudHandler ...
type DataCrudHandler struct {
	dataCrud dataCrud.DataCrudService
}

func (h DataCrudHandler) insertData(writer http.ResponseWriter, request *http.Request) {
	book, err := h.getFormData(writer, request)
	if err != nil {
		utils.HandleError(writer, http.StatusNotFound, err.Error())
		return
	}

	data, err := h.dataCrud.Insert(book)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, err.Error())
		fmt.Printf("[DataCrudHandler.Insert]Error when request data to service with error : %w\n", err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h *DataCrudHandler) getFormData(writer http.ResponseWriter, request *http.Request) (*models.YoutubeData, error) {
	channelId := request.FormValue("channel_id")
	title := request.FormValue("title")
	channelName := request.FormValue("channel_name")
	publishedAt := request.FormValue("published_at")


	data := models.YoutubeData{
		ChannelId: channelId,
		Title:        title,
		ChannelName: channelName,
		PublishedAt: publishedAt,
	}

	return &data, nil
}

func (h DataCrudHandler) getAll(writer http.ResponseWriter, request *http.Request) {
	data, err := h.dataCrud.GetAll()
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Ooops something error")

		fmt.Printf("[DataCrudHandler.getAll] Error when request data to service with error: %v\n", err)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, data)
}

func (h DataCrudHandler) delete(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])
	if err != nil {
		utils.HandleError(writer, http.StatusBadRequest, "ID not valid")
		fmt.Printf("[DataCrudHandler.delete]Error when convert pathvar with error : %v\n", err)
	}

	err = h.dataCrud.Delete(id)
	if err != nil {
		utils.HandleError(writer, http.StatusNoContent, "Oppss, something error")
		fmt.Printf("[DataCrudHandler.delete]Error when request data to service with error : %v\n", err)
		return
	}
	utils.HandleSuccess(writer, http.StatusOK, nil)
}

func (h DataCrudHandler) updateData(writer http.ResponseWriter, request *http.Request) {
	pathVar := mux.Vars(request)
	id, err := strconv.Atoi(pathVar["id"])

	_, err = h.dataCrud.GetDataByID(id)
	if err != nil {
		fmt.Printf("[DataCrudHandler.update] Error when check id to usecase with error: %v\n", err)
		utils.HandleError(writer, http.StatusBadRequest, "ID DOES NOT EXIST")
		return
	}

	var data = models.YoutubeData{}

	err = json.NewDecoder(request.Body).Decode(&data)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, "Oopss, something error")
		fmt.Printf("[DataCrudHandler.getData] Error when decode data with error: %v\n", err)
		return
	}

	dataUpdate, err := h.dataCrud.Update(id, &data)
	if err != nil {
		utils.HandleError(writer, http.StatusInternalServerError, err.Error())
		fmt.Printf("[UserHandler.update] Error when send data to usecase with error : %v", err)
		return
	}

	utils.HandleSuccess(writer, http.StatusOK, dataUpdate)
}

// CreateDataCrudHandler ...
func CreateDataCrudHandler(resp *mux.Router, dataCrud dataCrud.DataCrudService) {
	bookHandler := DataCrudHandler{dataCrud}

	resp.HandleFunc("/dataCrudAdd", bookHandler.insertData).Methods(http.MethodPost)

	resp.HandleFunc("/listData", bookHandler.getAll).Methods(http.MethodGet)
	resp.HandleFunc("/data/{id}", bookHandler.delete).Methods(http.MethodDelete)
	resp.HandleFunc("/dataUpdate/{id}", bookHandler.updateData).Methods(http.MethodPut)

}
