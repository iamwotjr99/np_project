package models

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Id        int    `json:"id"`
	AccountID string `json:"accountID"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

func (a Article) Get() (article Article, err error) {
	// db, err := sql.Open("mysql", "article_admin:np1234@tcp(mysql-articles)/toget_study")
	db, err := sql.Open("mysql", "root:asd98048@tcp(localhost:3306)/toget_study")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT * FROM article WHERE id=?", a.Id)
	err = row.Scan(&article.Id, &article.AccountID, &article.Content, &article.CreatedAt)
	if err != nil {
		return
	}
	return
}

func (a Article) GetAll() (articles []Article, err error) {
	db, err := sql.Open("mysql", "root:asd98048@tcp(localhost:3306)/toget_study")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM article")
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.Id, &article.AccountID, &article.Content, &article.CreatedAt)
		if err != nil {
			panic(err.Error())
		}
		articles = append(articles, article)
	}
	defer rows.Close()

	return articles, err
}

func Add(article Article) {
	// db, err := sql.Open("mysql", "article_admin:np1234@tcp(mysql-articles)/toget_study")
	db, err := sql.Open("mysql", "root:asd98048@tcp(localhost:3306)/toget_study")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO article(author, title, content) VALUES (?, ?, ?)",
		article.AccountID, article.Content)
	if err != nil {
		panic(err.Error())
	}

	// rs, err := stmt.Exec(a.Author, a.Title, a.Content)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// id, err := rs.LastInsertId()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// Id = int(id)
	defer insert.Close()
}
