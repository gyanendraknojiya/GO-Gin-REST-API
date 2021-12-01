package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
}

var users = []user{
	{
		ID:       "1",
		Name:     "Leanne Graham",
		UserName: "Bret",
		Email:    "Sincere@april.biz",
		Phone:    "1-770-736-8031 x56442",
		Website:  "hildegard.org",
	},
	{
		ID:       "2",
		Name:     "Ervin Howell",
		UserName: "Antonette",
		Email:    "Shanna@melissa.tv",
		Phone:    "010-692-6593 x09125",
		Website:  "anastasia.net",
	},
	{
		ID:       "3",
		Name:     "Clementine Bauch",
		UserName: "Samantha",
		Email:    "Nathan@yesenia.net",
		Phone:    "1-463-123-4447",
		Website:  "ramiro.info",
	},
	{
		ID:       "4",
		Name:     "Patricia Lebsack",
		UserName: "Karianne",
		Email:    "Julianne.OConner@kory.org",
		Phone:    "493-170-9623 x156",
		Website:  "kale.biz",
	},
	{
		ID:       "5",
		Name:     "Chelsey Dietrich",
		UserName: "Kamren",
		Email:    "Lucio_Hettinger@annie.ca",
		Phone:    "(254)954-1289",
		Website:  "demarco.info",
	},
}

func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	ID := c.Param("id")
	for _, i := range users {
		if i.ID == ID {
			c.IndentedJSON(http.StatusOK, i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found!"})
}

func addUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		fmt.Println(err)
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusOK, users)
}

func updateUser(c *gin.Context) {
	var updateUser user

	if err := c.BindJSON(&updateUser); err != nil {
		fmt.Println(err)
		return
	}
	if updateUser.ID == "" {
		c.IndentedJSON(400, gin.H{"message": "id is missing!"})
		return
	}

	for i, user := range users {
		if user.ID == updateUser.ID {

			if updateUser.Name != "" {
				user.Name = updateUser.Name
			}
			if updateUser.Email != "" {
				user.Email = updateUser.Email
			}
			if updateUser.Phone != "" {
				user.Phone = updateUser.Phone
			}
			if updateUser.Website != "" {
				user.Website = updateUser.Website
			}
			users[i] = user
			c.IndentedJSON(http.StatusOK, users)
			return
		}
	}
	c.IndentedJSON(400, gin.H{"message": "id is not found!"})
}

func RemoveIndex(s []user, index int) []user {
	return append(s[:index], s[index+1:]...)
}

func deleteUser(c *gin.Context) {
	var deleteUser user

	if err := c.BindJSON(&deleteUser); err != nil {
		fmt.Println(err)
		return
	}
	if deleteUser.ID == "" {
		c.IndentedJSON(400, gin.H{"message": "id is missing!"})
		return
	}

	for i, user := range users {
		if user.ID == deleteUser.ID {
			users = RemoveIndex(users, i)
			c.IndentedJSON(http.StatusOK, users)
			return
		}
	}
	c.IndentedJSON(400, gin.H{"message": "id is not found!"})
}

func main() {
	router := gin.Default()
	router.GET("/", getAllUsers)
	router.GET("/user/:id", getUserById)
	router.POST("/addUser", addUser)
	router.PUT("/updateUser", updateUser)
	router.DELETE("/deleteUser", deleteUser)

	router.Run("localhost:8080")
}
