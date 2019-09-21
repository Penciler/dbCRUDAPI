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
  //ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
  ID        string   `gorm:"primary_key;AUTO_INCREMENT"`
  Name     string `json:"name"`
  Email    string    `json:"email"`
 }

 userRes struct {
  //ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
  ID        string   `gorm:"primary_key;AUTO_INCREMENT"`
  Name     string `json:"name"`
  Email    string    `json:"email"`
 }

 model interface {
 	create(c *gin.Context) (returnUser userModel, err error)
 	read(id string) (returnUser userModel, err error)
 	update(id string, c *gin.Context) (returnUser userModel, err error)
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
	  v1.GET("/:id", getSingleUser(user))
	  v1.PATCH("/:id", updateUser(user))
	  /*
	  v1.GET("/", fetchAllUser)
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

func (user userModel) read(id string) (returnUser userModel, err error){
      if err := db.First(&user, id); err.Error != nil {
      	  return user, err.Error
      }
      return user, nil
}

func (user userModel) update(id string, c *gin.Context) (returnUser userModel, err error){
	  var temp userModel
      if err1 := db.First(&user, id); err1.Error != nil {
      	  return user, err1.Error
      }
      if err2 := c.BindJSON(&temp); err2 != nil {
          return user, err2
      }
      user.ID = id
      if err3 := db.Model(&user).Update("Name", temp.Name, "Email", temp.Email); err3.Error != nil {
      	  return user, err3.Error
      }
      return user, nil
}

// create user
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

// get single user data
func getSingleUser(user model) gin.HandlerFunc {
	return func(c *gin.Context){
		id := c.Param("id")
		returnUser, err := user.read(id) 
		if err != nil {
 			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "User not read!"})
 		}
 		userData := userRes{ID: returnUser.ID, Name: returnUser.Name, Email: returnUser.Email}

 		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": userData})
 	}

}

// update user data
func updateUser(user model) gin.HandlerFunc {
	return func(c *gin.Context){
		id := c.Param("id")
		returnUser, err := user.update(id, c) 
		if err != nil {
 			c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "User not update!"})
 		}

 		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User updated successfully!", "resourceId": returnUser.ID})
 	}

}