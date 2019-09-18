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
  ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
  Name     string `json:"name"`
  Email    string    `json:"email"`
 }

 model interface {
 	create(c *gin.Context) (returnUser userModel, err error)
 }

)

func main(){
	var user userModel
	router := setRoute(user)

	router.Run()
}

func setRoute(user model) (router *gin.Engine) {
	router = gin.Default()
	v1 := router.Group("api/vi/users")
	{
	  v1.POST("/", createUser(user))
	  /*
	  v1.GET("/", fetchAllUser)
	  v1.GET("/:id", fetchSingleUser)
	  v1.PUT("/:id", updateUser)
	  v1.DELETE("/:id", deleteUser)
	  */
	}
	return router
}

func initDB(dbType string, args string) (err error) {
	//db, err = gorm.Open("mysql","root:password@tcp(127.0.0.1:8081)/mysql?charset=utf8&parseTime=True&loc=Local")
	db, err = gorm.Open(dbType,args)
	if err != nil {
		panic("failed to connect database")
		println(err)
		return err
	}
	db.AutoMigrate(&userModel{})
	return nil
}

func (user userModel) create(c *gin.Context) (returnUser userModel, err error){
      if err1 := c.BindJSON(&user); err1 != nil {
          return user, err1
      }
      if err2 := db.Save(&user); err2.Error != nil {
      	  return user, err2.Error
      }
      return user, nil
}

// createTodo add a new todo
func createUser(user model) gin.HandlerFunc {
	return func(c *gin.Context){
		returnUser, err := user.create(c) 
		if err != nil {
 			c.JSON(http.StatusCreated, gin.H{"status": http.StatusInternalServerError, "message": "User not created!"})
 		}
 		//c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!", "resourceId": user.ID})
 		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created successfully!", "resourceId": returnUser.ID})
 	}

}