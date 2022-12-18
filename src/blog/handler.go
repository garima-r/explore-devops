// http related func : eg: Serve HTTP; validate request here

package blog

import(
	"log"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

// Hnadler to process the request for publishing artcile
func PublishArticleHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		log.Printf("[PublishArticleHandler] method not allowd")
		return
	}
	
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		log.Printf("[PublishArticleHandler] error getting request params - %s\n", err.Error())
		return  // try whether return is required or not
	}
	
	var newArticle Article
	err = json.Unmarshal(reqBody, &newArticle)
	if err != nil {
		http.Error(w, "Error unmarshalling request body", http.StatusBadRequest)
		log.Printf("[PublishArticleHandler] error unmarshalling request params - %s\n", err.Error())
		return
	}

	//request parameters validation
	if newArticle.Title == "" || newArticle.Content == "" {
		http.Error(w, "Invalid data in request body", http.StatusBadRequest)
		log.Printf("[PublishArticleHandler] error getting email_id - %s\n", err.Error())
		return
	}

	// call to logic
	err = SaveArticle(newArticle)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("[PublishArticleHandler] error creating article - %s\n", err.Error())
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Hnadler to process the request for displaying artciles
func DisplayArticlesHandler(w http.ResponseWriter, r *http.Request){

	if r.Method != "GET" {
        http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		log.Printf("[DisplayArticlesHandler] method not allowd")
        return
    }

	result, err := FetchAllArticles()
	if err != nil{
		http.Error(w, "Error fetching articles", http.StatusInternalServerError)
		log.Printf("[DisplayArticlesHandler] error fetching articles - %s\n", err.Error())
		return
	}

	jsonResponse, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error converting result to json", http.StatusInternalServerError)
		log.Printf("[DisplayArticlesHandler] converting result to json - %s\n", err.Error())
		return
	}

	w.Write(jsonResponse)
}









