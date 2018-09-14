package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ultimate/src/database"
	"ultimate/src/domain"
	"ultimate/src/infrastructure"
)


func Inset(context *gin.Context)  {
	db := database.Connection()
	tx , err := db.Begin()
	infrastructure.Error(err)
	stmt,err := db.Prepare("INSERT INTO user(name,attr,`lock`,age) VALUES(?,?,?,?)")
	infrastructure.Error(err)
	result , err := stmt.Exec("john","ok","lock",12)
	tx.Commit()
	db.Close()
	id , err := result.LastInsertId()
	context.JSON(http.StatusOK,gin.H{
		"id":id,
	})
}

func Update(context *gin.Context)  {
	db := database.Connection()
	tx , err := db.Begin()
	infrastructure.Error(err)
	stmt , err := db.Prepare("UPDATE user SET name = ? WHERE id = ?")
	infrastructure.Error(err)
	result ,err := stmt.Exec("von",77)
	tx.Commit()
	db.Close()
	id , err := result.LastInsertId()
	infrastructure.Error(err)
	context.JSON(http.StatusOK,gin.H{
		"id":id,
	})
}

func Read(context *gin.Context)  {
	db := database.Connection()
	tx , err := db.Begin()
	infrastructure.Error(err)
	row , err := db.Query("SELECT * FROM user")
	tx.Commit()
	db.Close()
	var mod []domain.User
	for row.Next() {
		var user domain.User
		err := row.Scan(&user.ID,&user.Name,&user.Attr,&user.Lock,&user.Age)
		infrastructure.Error(err)
		mod = append(mod,user)
	}
	context.JSON(http.StatusOK,gin.H{
		"mod":mod,
	})

}

func Remove(context *gin.Context)  {
	db := database.Connection()
	tx , err := db.Begin()
	infrastructure.Error(err)
	stmt , err := db.Prepare("DELETE FROM user WHERE user.id = ?")
	infrastructure.Error(err)
	result , err := stmt.Exec(78)
	tx.Commit()
	db.Close()
	id , err := result.LastInsertId()
	infrastructure.Error(err)
	context.JSON(http.StatusOK,gin.H{
		"id":id,
	})
}