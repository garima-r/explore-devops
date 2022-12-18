package blog

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestStoreArticle(t *testing.T) {
	var mock sqlmock.Sqlmock
	var err error

	db, mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		data Article
	}

	cases := []struct {
		name     string
		args     args
		wantErr  bool
		mockFunc func(args)
	}{
		{
			name: "Success : Store Article in DB",
			args: args{
				data: Article{
					Author:  "graw",
					Title:   "seventh article",
					Content: "ready to publish my seventh article",
				},
			},
			wantErr: false,
			mockFunc: func(a args) {
				//mock.ExpectExec("INSERT into articles (.+) VALUES (.+)").WillReturnResult(sqlmock.NewResult(1,1))
				mock.ExpectExec("INSERT into articles").WithArgs(a.data.Author, a.data.Title, a.data.Content).WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "Fail : Store Article in DB",
			args: args{
				data: Article{
					Author:  "graw",
					Title:   "seventh article",
					Content: "ready to publish my eight article",
				},
			},
			wantErr: true,
			mockFunc: func(a args) {
				//mock.ExpectExec("INSERT into articles (.+) VALUES (.+)").WillReturnResult(sqlmock.NewResult(1,1))
				mock.ExpectExec("INSERT into articles").WithArgs(a.data.Author, a.data.Title, a.data.Content).WillReturnError(fmt.Errorf("error duplicate row"))
			},
		},
	}

	for i, tc := range cases {
		tc.mockFunc(tc.args)
		err := StoreArticle(tc.args.data)
		if (err != nil) != tc.wantErr {
			t.Errorf("[TestStoreArticle] for case %d : error = %v , wantErr %v\n", i+1, err, tc.wantErr)
			return
		}
	}
}

func TestRetrieveAllArticles(t *testing.T) {
	var mock sqlmock.Sqlmock
	var err error

	db, mock, err = sqlmock.New()
	if err != nil {
		t.Errorf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	type args struct {
		data Article
	}

	cases := []struct {
		name     string
		want     []Article
		wantErr  bool
		mockFunc func()
	}{
		{
			name: "Success : Retrieve Article from DB",
			want: []Article{
				{
					Author:  "graw",
					Title:   "seventh article",
					Content: "ready to publish my seventh article",
				},
				{
					Author:  "graw",
					Title:   "eight article",
					Content: "ready to publish my eight article",
				},
			},
			wantErr: false,
			mockFunc: func() {
				rows := sqlmock.NewRows([]string{"author", "title", "content"}).
					AddRow("graw", "seventh article", "ready to publish my seventh article").
					AddRow("graw", "eight article", "ready to publish my eight article")
				mock.ExpectQuery(`select (.+) from articles`).WillReturnRows(rows)
			},
		},
		{
			name:    "Fail : Retrieve Article from DB",
			want:    []Article{},
			wantErr: true,
			mockFunc: func() {
				mock.ExpectQuery(`select (.+) from articles`).WillReturnError(errors.New("some error"))
			},
		},
	}

	for i, tc := range cases {
		tc.mockFunc()
		got, err := RetrieveAllArticles()
		if (err != nil) != tc.wantErr {
			t.Errorf("RetrieveAllArticles() for case %d : error = %v, wantErr %v\n", i+1, err, tc.wantErr)
			return
		}
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("For case %d : RetrieveAllArticles() = %v, want %v\n", i+1, got, tc.want)
		}
	}

}
