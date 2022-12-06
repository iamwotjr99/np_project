package models

import (
	"database/sql"
	"fmt"
	"log"
	"mime/multipart"

	_ "github.com/go-sql-driver/mysql"
)

type Article struct {
	Author    string                  `form:"author" binding:"required"`
	AccountID string                  `form:"accountID" binding:"required"`
	Content   string                  `form:"content" binding:"required"`
	LikeCnt   string                  `form:"likecnt" binding:"required"`
	Images    []*multipart.FileHeader `form:"file"`
}

type Res_Article struct {
	PostID    string  `json:"postID"`
	Author    string  `json:"author"`
	AccountID string  `json:"accountID"`
	Content   string  `json:"content"`
	CreatedAt string  `json:"createdAt"`
	LikeCnt   string  `json:"liekcnt"`
	Images    []Image `json:"imageList"`
}

type Image struct {
	Filename string `json:"filename"`
	Filepath string `json:"filepath"`
}

type ImageList struct {
	Images []Image
}

func (il *ImageList) AddItem(image Image) []Image {
	il.Images = append(il.Images, image)
	return il.Images
}

func Get(id int) (article Res_Article, err error) {
	db, err := sql.Open("mysql", "root:asd98048@tcp(localhost:3306)/healer_com")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	query, err := db.Query("SELECT * FROM article WHERE id = ?", id)
	if err != nil {
		log.Fatalln(err)
	}

	var resultArticle Res_Article
	images := []Image{}

	for query.Next() {
		err = query.Scan(&resultArticle.PostID, &resultArticle.Author, &resultArticle.AccountID,
			&resultArticle.Content, &resultArticle.CreatedAt, &resultArticle.LikeCnt)
		if err != nil {
			panic(err.Error())
		}

		queryImages, err := db.Query("SELECT * FROM article_images WHERE article_id = ?", id)
		if err != nil {
			panic(err.Error())
		}

		for queryImages.Next() {
			var filename sql.NullString
			var filepath sql.NullString
			var fileID sql.NullInt64
			var id sql.NullInt64

			err = queryImages.Scan(&id, &filename, &filepath, &fileID)
			if err != nil {
				panic(err.Error())
			}
			defer queryImages.Close()

			image := Image{Filename: filename.String, Filepath: filepath.String}

			images = append(images, image)
		}
	}
	defer query.Close()

	resultArticle.Images = images

	return resultArticle, err
}

func GetAll() (articles []Res_Article, err error) {
	db, err := sql.Open("mysql", "root:asd98048@tcp(localhost:3306)/healer_com")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM article ORDER BY id DESC")
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var article Res_Article
		images := []Image{}

		err = rows.Scan(&article.PostID, &article.Author, &article.AccountID, &article.Content, &article.CreatedAt,
			&article.LikeCnt)
		if err != nil {
			panic(err.Error())
		}

		rows_images, err := db.Query("SELECT * FROM article_images WHERE article_id = ?", article.PostID)
		if err != nil {
			panic(err.Error())
		}

		for rows_images.Next() {
			var filename sql.NullString
			var filepath sql.NullString
			var fileID sql.NullInt64
			var id sql.NullInt64

			err = rows_images.Scan(&id, &filename, &filepath, &fileID)
			if err != nil {
				panic(err.Error())
			}

			image := Image{Filename: filename.String, Filepath: filepath.String}

			images = append(images, image)
		}
		defer rows_images.Close()

		article.Images = images

		articles = append(articles, article)

	}
	defer rows.Close()

	return articles, err
}

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
	if err != nil {
		panic(err.Error())
	}

	for _, image := range imageList.Images {
		result_image, err := db.Exec("INSERT INTO article_images(file_name, file_url, article_id) VALUES(?, ?, ?)",
			image.Filename, image.Filepath, id)
		if err != nil {
			panic(err.Error())
		}

		n, err := result_image.RowsAffected()
		if err != nil {
			panic(err.Error())
		}
		if n == 1 {
			fmt.Println("1 row inserted.")
		}
	}
}

func Update(article Article, imageList ImageList, id int) {
	db, err := sql.Open("mysql", "root:asd98048@tcp(localhost:3306)/healer_com")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	query, err := db.Exec("UPDATE article SET author=?, accountID=?, content=?, likeCnt=? WHERE id=?",
		article.Author, article.AccountID, article.Content, article.LikeCnt, id)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(query)

	for _, image := range imageList.Images {
		queryImages, err := db.Exec("UPDATE article_images SET file_name=?, file_url=? WHERE article_id=?",
			image.Filename, image.Filepath, id)
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(queryImages.RowsAffected())
	}
}

func Delete(id int) (err error) {
	db, err := sql.Open("mysql", "root:asd98048@tcp(localhost:3306)/healer_com")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	delImages, err := db.Exec("DELETE from article_images WHERE article_id=?", id)
	if err != nil {
		panic(err.Error())
	}

	delImages_id, err := delImages.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Delete Images ID: ", delImages_id)

	delArticle, err := db.Exec("DELETE from article WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	delArticle_id, err := delArticle.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Delete Article ID: ", delArticle_id)

	return err
}
