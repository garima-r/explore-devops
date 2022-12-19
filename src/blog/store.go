package blog

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[InitDB] Error loading .env file")
	}

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
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
	if err != nil {
		return []Article{}, err
	}

	var articles []Article
	for rows.Next() {
		var author, title, content string

		err = rows.Scan(&author, &title, &content)
		if err != nil {
			return articles, err
		}

		articles = append(articles, Article{Author: author, Title: title, Content: content})
	}

	return articles, nil
}
