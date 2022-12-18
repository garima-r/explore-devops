package blog

import(
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"fmt"
)

const (
	DB_HOST     = "localhost"
	DB_PORT     = 5432
    DB_USER     = "postgres"
    DB_PASSWORD = "123456789"
    DB_NAME     = "blog_demo"
)

var db *sql.DB

// InitDB  sets up DB
func InitDB() {
	var err error

    dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
    db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Printf("[initDB] error initializing postgres db - %s\n", err.Error())
		log.Fatal("Failure while connecting to slave", err)
	}
}

//StoreArticle func stores article in db
var StoreArticle = func(article Article) error {
	storeArticleQuery := `INSERT into articles (author, title, content) VALUES ($1, $2, $3)` 

	_, err := db.Exec(storeArticleQuery, article.Author, article.Title, article.Content)
	if err != nil {
		return err
	}

	return nil
}

//RetrieveAllArticles func retrieves articles from db
var RetrieveAllArticles = func() ([]Article, error) {
	rows, err := db.Query("select author, title, content from articles")
	if err != nil{
		return []Article{}, err
	}

	var articles []Article
	for rows.Next(){
		var author, title, content string

		err = rows.Scan(&author, &title, &content)
		if err != nil {
			return articles, err
		}

		articles = append(articles, Article{Author: author, Title : title, Content: content})
	}

	return articles, nil
}
