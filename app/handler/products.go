package handler

import (
	"encoding/json"
	"net/http"
	"github.com/jinzhu/gorm"
	"github.com/gorilla/mux"
    "../model"
)



func GetAllProducts(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	products := []model.Product{}
	db.Find(&products)
	respondJSON(w, http.StatusOK, products)
}

// insert(Post) a product into the database
func CreateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	product := model.Product{}

	decoder := json.NewDecoder(r.Body)
	
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, product)
}

//Get detail of individual product
func GetProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productName := vars["name"]
	product := getProductOr404(db, productName, w, r)
	if product == nil {
		return
	}

	respondJSON(w, http.StatusOK, product)
}


//delete a single product
func DeleteProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	product := getProductOr404(db, name, w, r)
	if product == nil {
		return
	}
	if err := db.Delete(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusNoContent, nil)
}


func UpdateProduct(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]
	product := getProductOr404(db, name, w, r)
	if product == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	if err := db.Save(&product).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, product)
}


//Get a single product by name.
func getProductOr404(db *gorm.DB, name string, w http.ResponseWriter, r *http.Request) *model.Product {
	product := model.Product{}
	if err := db.First(&product, model.Product{Name: name}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &product
}


