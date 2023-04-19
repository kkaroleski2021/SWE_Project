package product

import (
	"context"
	"encoding/json"
	"fmt"

	//"go/service"
	"strconv"
	"sync"
	"time"

	//"io/ioutil"
	//"strconv"

	"io"
	"net/http"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var rand uint32
var randmu sync.Mutex

type ProdInterface interface {
	InitialMigration(DNS string)
	AddProduct(w http.ResponseWriter, r *http.Request)
	UploadImg(w http.ResponseWriter, r *http.Request)
	//UploadImg(image io.Reader, token string)
}

// get db obj that was created in user package
var DB *gorm.DB
var imgDB *gorm.DB
var orderedProductsDB *gorm.DB
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

func OrderedProductsMigration(DNS string) {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	} else {
		DB.AutoMigrate(&OrderedProduct{})
	}
	orderedProductsDB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	} else {
		imgDB.AutoMigrate(&OrderedProduct{})
	}
}

/* first implementation of Add Product
	images and text for product are in one form-data type ..
	you have to select the entry by exact name "name","price", "tags" , "file" etc

second implementation getting json info and getting the form data for images is diff..
 could be used if we have a separate path for uploading images or something
*/

/*
func AddProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")

	r.ParseMultipartForm(200000)
	formdata := r.MultipartForm
	name := formdata.File["name"]
	for i := range name {
		data, err := name[i].Open()
		if err != nil {
			json.NewEncoder(w).Encode("Error Retrieving Name")
			fmt.Println(err)
			return
		}
		defer data.Close()


	}

	mr, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var product Product
	var files []Upfile
	var upfile Upfile
	for {
		part, err := mr.NextPart()

		// This is OK, no more parts
		if err == io.EOF {
			break
		}

		// Some error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		println(part.FormName())
		println(part.Header.Get("name"))

		// JSON 'text' part
		if part.FormName() == "name" {
			jsonDecoder := json.NewDecoder(part)
			err = jsonDecoder.Decode(&product)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}

		// img 'file' part
		if part.FormName() == "file" {
			upfile.Fname = part.FileName()

			fmt.Println("URL:", part.FileName())
			outfile, err := os.Create("./frontend/src/assets/uploads" + part.FileName())
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			upfile.Path = outfile.Name()
			defer outfile.Close()

			_, err = io.Copy(outfile, part)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			files = append(files, upfile)
		}

	}
	result := DB.Create(&product)
	if result.Error != nil {
		fmt.Println(result.Error)
	} else {
		json.NewEncoder(w).Encode(product)
	}

	for i := 0; i < len(files); i++ {
		files[i].ProdID = int(product.ID)
		result := imgDB.Create(&files[i])
		if result.Error != nil {
			fmt.Println(result.Error)
		} else {
			json.NewEncoder(w).Encode(files[i])
		}
	}
} */

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
	json.NewEncoder(w).Encode(product)
	//UploadImg(w, r, int(product.ID))

	//write the id to the request to pass along to next function
	ctx := context.WithValue(r.Context(), "request_id", product.ID)
	http.Redirect(w, r.WithContext((ctx)), "/newlisting/addimages", http.StatusSeeOther)
}

// For orders
func AddOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Order Order
	//var order Order
	json.NewDecoder(r.Body).Decode(&Order)
	//err := decoder.Decode(&Order)
	result := orderedProductsDB.Create(&Order)
	if result.Error != nil {
		fmt.Println(err.Error())
	}
	json.NewEncoder(w).Encode(Order)
	fmt.Println("Order successfully created")
}

//adding order to the database

func UploadImg(w http.ResponseWriter, r *http.Request) {
	mr, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var files []Upfile
	var upfile Upfile
	for {
		part, err := mr.NextPart()
		// This is OK, no more parts
		if err == io.EOF {
			break
		}
		// Some error
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// img 'file' part
		if part.FormName() == "file" {
			upfile.Fname = randomName(part.FileName())
			outfile, err := os.Create("./frontend/src/assets/uploads/" + upfile.Fname)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			upfile.Path = outfile.Name()
			defer outfile.Close()

			_, err = io.Copy(outfile, part)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			files = append(files, upfile)
		}

	}
	prodID := r.Context().Value("request_id").(int)
	for i := 0; i < len(files); i++ {
		files[i].ProdID = int(prodID)
		result := imgDB.Create(&files[i])
		if result.Error != nil {
			fmt.Println(result.Error)
		} else {
			json.NewEncoder(w).Encode(files[i])
		}
	}
}

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func nextSuffix() string {
	randmu.Lock()
	r := rand
	if r == 0 {
		r = reseed()
	}
	r = r*1664525 + 1013904223 // constants from Numerical Recipes
	rand = r
	randmu.Unlock()
	return strconv.Itoa(int(1e9 + r%1e9))[1:]
}

func randomName(name string) string {
	dir := "./frontend/src/assets/uploads/"
	nconflict := 0
	origName := name
	for i := 0; i < 10000; i++ {
		newName := nextSuffix() + origName
		fullPath := dir + newName
		f, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if os.IsExist(err) {
			if nconflict++; nconflict > 10 {
				randmu.Lock()
				rand = reseed()
				randmu.Unlock()
			}
			continue
		}
		f.Close()
		return newName
	}
	return randomName(name)
}

//doesnt work ?
/* func UploadImg(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "multipart/form-data")
	prodID := r.Context().Value("request_id").(int)
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
		println(fileHeaders[i].Filename)
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
		} else {
			json.NewEncoder(w).Encode(tempFile)
		}
		//http.Redirect(w, r, "/", 301)
	}

	json.NewEncoder(w).Encode("File Uploaded Successfully")
} */
