package database

import (
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

func (db *appdbimpl) GetCommentList(idPhoto int) ([]string, error) {
	var CommentList []string

	CommentRows, err := db.c.Query(`
	SELECT Username,Text
	FROM  User,Comment
	WHERE Comment.IdUser=User.Username and  IdPhoto=`, idPhoto)

	defer CommentRows.Close()

	for CommentRows.Next() {
		var Username string
		var Text string
		CommentRows.Scan(&Username)
		CommentRows.Scan(&Text)
		CommentList = append(CommentList, Username, Text)
	}

	// format the list
	return CommentList, err
}
