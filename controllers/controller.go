package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/naveenkakumanu/todoapp/db"
)

func Create(c *gin.Context) {
	var task db.Task
	c.BindJSON(&task)
	conn := db.DB()
	conn.Create(&task)
	c.JSON(200, gin.H{"message": "Created Task Successfully"})
}

func Read(c *gin.Context) {
	var task db.Task
	id := c.Param("id")
	conn := db.DB()
	conn.Find(&task, "id = ?", id)
	c.JSON(200, task)
}

func Update(c *gin.Context) {
	var task db.Task
	id := c.Param("id")
	c.BindJSON(&task)
	conn := db.DB()
	conn.Model(&task).Where("id", id).Update("name", task.Name)
	c.JSON(200, task)
}

func Delete(c *gin.Context) {
	var task db.Task
	id := c.Param("id")
	conn := db.DB()
	conn.Find(&task, "id = ?", id)
	c.JSON(200, task)
}

func CreateUser(c *gin.Context) {
	var user db.User
	c.BindJSON(&user)
	conn := db.DB()
	conn.Create(&user)
	c.JSON(200, gin.H{"message": "Created user Successfully"})
}

func ReadUser(c *gin.Context) {
	var user db.User
	id := c.Param("id")
	conn := db.DB()
	conn.Find(&user, "id = ?", id)
	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	var user db.User
	id := c.Param("id")
	c.BindJSON(&user)
	conn := db.DB()
	conn.Model(&user).Where("id", id).Update("name", user.Name)
	c.JSON(200, user)

}

// DeleteUser
func DeleteUser(c *gin.Context) {
	var user db.User
	id := c.Param("id")
	conn := db.DB()
	conn.Delete(&user, "id = ?", id)
	c.JSON(200, user)

}
