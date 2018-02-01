package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/JustonDavies/url_shortener/models"
	base62 "github.com/JustonDavies/url_shortener/utils"
)

type DatabaseClient struct {
	database *sql.DB
}

type Record struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

func (client *DatabaseClient) GetOriginalURL(writer http.ResponseWriter, request *http.Request) {
	var url string
	var parameters = mux.Vars(request)

	// Get ID from base62 string
	var id = base62.ToBase10(parameters["encoded_string"])
	var exception = client.database.QueryRow(
		"SELECT url FROM web_url WHERE id = $1",id).Scan(&url)

	// Handle response details
	if exception != nil {
		writer.Write([]byte(exception.Error()))
	} else {
		writer.WriteHeader(http.StatusOK)
		writer.Header().Set(
			"Content-Type",
			"application/json",
			)

		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		writer.Write(response)
	}
}

func (client *DatabaseClient) GenerateShortURL(writer http.ResponseWriter, request *http.Request) {
	var id int
	var record Record

	var postBody, _ = ioutil.ReadAll(request.Body)
	json.Unmarshal(postBody, &record)

	var exception = client.database.QueryRow(
		"INSERT INTO web_url(url) VALUES($1) RETURNING id",
		record.URL,
		).Scan(&id)

	var responseMap = map[string]interface{}{"encoded_string": base62.ToBase62(id)}

	if exception != nil {
		writer.Write([]byte(exception.Error()))
	} else {
		writer.Header().Set(
			"Content-Type",
			"application/json",
				)

		response, _ := json.Marshal(responseMap)
		writer.Write(response)
	}
}

func main() {
	var database, exception = models.InitializeDatabase()
	if exception != nil {
		log.Println(exception)
	}

	var client = &DatabaseClient{database: database}
	//if exception != nil {
	//	panic(exception)
	//}

	defer database.Close()

	// Create a new router
	var router = mux.NewRouter()

	// Attach an elegant path with handler
	router.HandleFunc(
		"/v1/short/{encoded_string:[a-zA-Z0-9]*}",
		client.GetOriginalURL,
		).Methods("GET")

	router.HandleFunc(
		"/v1/short",
		client.GenerateShortURL,
		).Methods("POST")

		// Construct Server
	var server = &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}


//func junk() {
//	log.Println(128/72)
//	log.Println(int(math.Floor(float64(128/72))))
//
//	db, err := models.InitializeDatabase()
//	if err != nil {
//		log.Println(db)
//	}
//
//	x := 100
//	base62String := base62.ToBase62(x)
//	log.Println(base62String)
//	normalNumber := base62.ToBase10(base62String)
//	log.Println(normalNumber)
//}