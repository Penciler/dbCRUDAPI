package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/jinzhu/gorm"
    _"github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type (
 // userModel describes user data
 userModel struct {
  gorm.Model
  Name     string `json:"name"`
  Email    string    `json:"email"`
 }
)

func main(){
	router := gin.Default()
	v1 := router.Group("api/vi/users")
	{
	  v1.POST("/", createUser)
	  /*
	  v1.GET("/", fetchAllUser)
	  v1.GET("/:id", fetchSingleUser)
	  v1.PUT("/:id", updateUser)
	  v1.DELETE("/:id", deleteUser)
	  */
	}

	router.Run()
}

func init(){
	var err error
	db, err = gorm.Open("mysql","root:password@tcp(127.0.0.1:8081)/mysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
		println(err)
	}
	db.AutoMigrate(&userModel{})
}

// createTodo add a new todo
func createUser(c *gin.Context) {
 user := userModel{Name: c.PostForm("name"), Email: c.PostForm("email")}
 db.Save(&user)
 c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!", "resourceId": user.ID})
}