package models

import (
	"database/sql"
	"fmt"
	"mime/multipart"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Author    string                  `form:"author" binding:"required"`
	AccountID string                  `form:"accountID" binding:"required"`
	Content   string                  `form:"content" binding:"required"`
	LikeCnt   string                  `form:"likecnt" binding:"required"`
	Images    []*multipart.FileHeader `form:"file" binding:"required"`
}

type Image struct {
	Filename string
	Filepath string
}

type ImageList struct {
	Images []Image
}

func (il *ImageList) AddItem(image Image) []Image {
	il.Images = append(il.Images, image)
	return il.Images
}

// func (a Article) Get() (article Article, err error) {
// 	// db, err := sql.Open("mysql", "article_admin:np1234@tcp(mysql-articles)/toget_study")
// 	db, err := sql.Open("mysql", "root:asd98048@tcp(localhost:3306)/toget_study")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer db.Close()

// 	row := db.QueryRow("SELECT * FROM article WHERE id=?", a.Id)
// 	err = row.Scan(&article.Id, &article.AccountID, &article.Content, &article.CreatedAt)
// 	if err != nil {
// 		return
// 	}
// 	return
// }

// func (a Article) GetAll() (articles []Article, err error) {
// 	db, err := sql.Open("mysql", "root:asd98048@tcp(localhost:3306)/toget_study")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT * FROM article")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	for rows.Next() {
// 		var article Article
// 		err = rows.Scan(&article.Id, &article.AccountID, &article.Content, &article.CreatedAt)
// 		if err != nil {
// 			panic(err.Error())
// 		}
// 		articles = append(articles, article)
// 	}
// 	defer rows.Close()

// 	return articles, err
// }

func Add(article Article, imageList ImageList) {
	// db, err := sql.Open("mysql", "article_admin:np1234@tcp(mysql-articles)/toget_study")
	db, err := sql.Open("mysql", "root:asd98048@tcp(localhost:3306)/healer_com")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Exec("INSERT INTO article(author, accountID, content, likeCnt) VALUES (?, ?, ?, ?)",
		article.Author, article.AccountID, article.Content, article.LikeCnt)
	if err != nil {
		panic(err.Error())
	}

	id, err := result.LastInsertId()

	for _, image := range imageList.Images {
		result_image, err := db.Exec("INSERT INTO article_images(file_name, file_url, article_id) VALUES(?, ?, ?)",
			image.Filename, image.Filepath, id)
		if err != nil {
			panic(err.Error())
		}

		n, err := result_image.RowsAffected()
		if n == 1 {
			fmt.Println("1 row inserted.")
		}
	}
}
