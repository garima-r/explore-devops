package blog

var SaveArticle = func(article Article) error{
	err := StoreArticle(article)
	if err != nil {
		return err
	}

	return nil
}


func FetchAllArticles() ([]Article, error){
	articleData, err := RetrieveAllArticles()
	if err != nil {
		return []Article{}, err
	}

	return articleData, nil
}