package blog

import(
	"testing"
	"errors"
	"reflect"
)

func TestSaveArticle(t *testing.T){
	cases := []struct {
		name     string
		arg   Article
		wantErr bool
		mocks   map[string]interface{}
	}{
		{
			name : "Success : Saved Article",
			arg : Article{
				Author : "graw",
				Title : "fifth article",
				Content : "ready to publish my second article",
			},
			wantErr : false,
			mocks: map[string]interface{}{
				"StoreArticle": nil,
			},
		},
		{	
			name : "Fail : Error in Saving Article",
			arg : Article{
				Author : "graw",
				Title : "fifth article",
				Content : "ready to publish my seventh article",
			},
			wantErr : true,
			mocks: map[string]interface{}{
				"StoreArticle": errors.New("Failed to save article"),
			},
		},	
	}

	// backup real funcs
	realFuncs := make(map[string]interface{}, 0)
	realFuncs["StoreArticle"] = StoreArticle

	var unmock = func() {
		StoreArticle = realFuncs["StoreArticle"].(func(article Article) error)
	}


	// mock funcs
	var initMock = func(mocks map[string]interface{}) {
		StoreArticle = func(article Article) error {
			if err, ok := mocks["StoreArticle"].(error); ok{
				return err
			}
			return nil
		}
	}

	for i, tc := range cases {
		initMock(tc.mocks)

		err := SaveArticle(tc.arg)
		if (err != nil) != tc.wantErr {
			t.Errorf("[TestSaveArticle] for case %d : error = %v , wantErr %v\n", i+1, err, tc.wantErr)
			return
		}
	}

	// put back funcs to initial state
	unmock()
}

func TestFetchAllArticles(t *testing.T){
	cases := []struct {
		name     string
		want []Article
		wantErr bool
		mocks   map[string]interface{}
	}{
		{
			name : "Success : Retrieved all articles",
			want : []Article{
				{
					Author : "graw",
					Title : "first article",
					Content : "Published the first article",
				},
				{
					Author : "graw",
					Title : "second article",
					Content : "Published the second article",
				},
			},
			wantErr : false,
			mocks : map[string]interface{}{
				"RetrieveAllArticles" : []Article{
					{
						Author : "graw",
						Title : "first article",
						Content : "Published the first article",
					},
					{
						Author : "graw",
						Title : "second article",
						Content : "Published the second article",
					},
				},
			},
		},
		{
			name : "Fail : Failure in Retrieving all articles",
			want : []Article{},
			wantErr : true,
			mocks : map[string]interface{}{
				"RetrieveAllArticles" : errors.New("Failed to save article"),
			},
		}, 
	}

	// backup real funcs
	realFuncs := make(map[string]interface{}, 0)
	realFuncs["RetrieveAllArticles"] = RetrieveAllArticles

	var unmock = func() {
		RetrieveAllArticles = realFuncs["RetrieveAllArticles"].(func() ([]Article, error))
	}

	// mock funcs
	var initMock = func(mocks map[string]interface{}) {
		RetrieveAllArticles = func() ([]Article, error) {
			if articles, ok := mocks["RetrieveAllArticles"].([]Article); ok{
				return articles, nil
			} else if err, ok := mocks["RetrieveAllArticles"].(error); ok{
				return []Article{}, err
			}

			return []Article{}, nil
		}
	}

	for i, tc := range cases {
		initMock(tc.mocks)

		got, err := FetchAllArticles()
		if (err != nil) != tc.wantErr {
			t.Errorf("[TestSaveArticle] for case %d : error = %v , wantErr %v\n", i+1, err, tc.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("For case %d : FetchAllArticles() = %v, want %v\n", i+1, got, tc.want)
		}
	}

	// put back funcs to initial state
	unmock()
}


