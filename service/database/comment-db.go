package database

import (
	"fmt"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/components"
)

func (db *appdbimpl) InsertComment(comment components.Comment) error {

	_, err := db.c.Exec(`

	INSERT OR IGNORE INTO Comment(IdUser,IdPhoto,DataTime,Text) 
		VALUES(?,?,?,?)`,
		comment.User.Id, comment.IdPhoto, comment.UploadDataTime, comment.Text)

	if err != nil {
		return err
	}
	return nil
}

func (db *appdbimpl) DeleteComment(commentId int) error {

	res, err := db.c.Exec(`
		DELETE FROM Comment
		WHERE IdComment=? `, commentId)

	if err != nil {
		return err
	}

	aff, err := res.RowsAffected()

	// if there are no affected rows
	// then the photo was not commented
	if aff == 0 {
		return ErrCommNotFound
	}
	return nil

}

func (db *appdbimpl) GetComment(commentId int) (components.Comment, error) {
	var comment components.Comment

	err := db.c.QueryRow(`
	SELECT IdComment,IdUser,IdPhoto,DataTime,Text
	FROM Comment
	WHERE IdComment=?`, commentId).Scan(&comment.IdComment, &comment.User.Id, &comment.IdPhoto, &comment.UploadDataTime, &comment.Text)
	if err != nil {
		return comment, ErrCommNotFound
	}
	return comment, nil
}

func (db *appdbimpl) GetCommentList(idPhoto int) ([]components.Comment, error) {
	var CommentList []components.Comment
	
	CommentRows, err := db.c.Query(`
	SELECT User.Username,User.IdUser,Comment.IdPhoto,Comment.IdComment,Comment.DataTime,Comment.Text
	FROM  User,Comment
	WHERE Comment.IdUser=User.IdUser and  Comment.IdPhoto=?`, idPhoto)

	fmt.Println("query eseguita si trappa")

	if err != nil {
		return CommentList, err
	}
	defer CommentRows.Close()

	for CommentRows.Next() {
		var comment components.Comment
		CommentRows.Scan(&comment.User.Username,&comment.User.Id,&comment.IdPhoto,&comment.IdComment,&comment.UploadDataTime,&comment.Text)
		CommentList = append(CommentList, comment)
	}

	fmt.Println("comment LIST", CommentList)
	// format the list
	return CommentList, nil
}
