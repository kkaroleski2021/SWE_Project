package product

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProdInterface interface {
	InitialMigration(DNS string)
	AddProduct(w http.ResponseWriter, r *http.Request)
	UploadImg(w http.ResponseWriter, r *http.Request, prodID int)
	//UploadImg(image io.Reader, token string)
}

// get db obj that was created in user package
var DB *gorm.DB
var imgDB *gorm.DB
var err error

func InitialMigration(DNS string) {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	} else {
		DB.AutoMigrate(&Product{})
	}
	imgDB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	} else {
		imgDB.AutoMigrate(&Upfile{})
	}
}

func AddProduct(w http.ResponseWriter, r *http.Request) {
	//var id = GetPathID(r)
	w.Header().Set("Content-Type", "application/json")
	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	product.UserId = r.Context().Value("request_id").(int)
	// figure out tags
	result := DB.Create(&product)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	UploadImg(w, r, int(product.ID))
	json.NewEncoder(w).Encode(product)
}

func UploadImg(w http.ResponseWriter, r *http.Request, prodID int) {
	w.Header().Set("Content-Type", "application/json")
	//max upload of 10 MB files
	r.ParseMultipartForm(200000)
	formdata := r.MultipartForm
	fileHeaders := formdata.File["files"]

	for i := range fileHeaders {
		file, err := fileHeaders[i].Open()
		if err != nil {
			json.NewEncoder(w).Encode("Error Retrieving File")
			fmt.Println(err)
			return
		}
		defer file.Close()

		upfile := Upfile{
			ProdID: prodID,
			Fname:  fileHeaders[i].Filename,
			Fsize:  strconv.FormatInt(int64(fileHeaders[i].Size), 10),
		}
		upfile.Ftype = fileHeaders[i].Header.Get("Content-type")

		//Create file
		tempFile, err := ioutil.TempFile("frontend/src/assets/uploads", "upload-*.jpg")
		if err != nil {
			fmt.Println(err)
		}
		defer tempFile.Close()
		upfile.Path = tempFile.Name()

		//read all contents of uploaded file into byte array
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
		}
		// write this byte array to our temporary file
		tempFile.Write(fileBytes)

		result := imgDB.Create(&upfile)
		if result.Error != nil {
			fmt.Println(result.Error)
		}
		//http.Redirect(w, r, "/", 301)
	}

	json.NewEncoder(w).Encode("File Uploaded Successfully")
}
