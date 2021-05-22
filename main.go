package main

import (
	"TestScrapeCRUD/driver"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"log"
	"net/http"
	"os"

dataCrudHandler "TestScrapeCRUD/dataCrud/handler"
dataCrudService "TestScrapeCRUD/dataCrud/service"
dataCrudRepo "TestScrapeCRUD/dataCrud/repository"
)

func init() {
	gotenv.Load()
}

func main() {
	port := os.Getenv("PORT")
	db := driver.Connect()
	defer db.Close()
	driver.InitTable(db)

	router := mux.NewRouter().StrictSlash(true)

	dataCrudRepo := dataCrudRepo.CreateDataCrudRepoImpl(db)
	dataCrudService := dataCrudService.CreateDataCrudServiceImpl(dataCrudRepo)
	dataCrudHandler.CreateDataCrudHandler(router, dataCrudService)

	fmt.Println("Starting web server at port : ", port)
	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		log.Fatal()
	}
}
