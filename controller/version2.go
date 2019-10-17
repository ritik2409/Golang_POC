package controller

import (
	userobj2 "project/dto"
	"project/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Version2Controller struct{}

var users2 []userobj2.Users

// func (v2 Version2Controller) Init() {
// 	// u := new(dto.userobj)
// 	users2 = userobj2.Init()
// }

// type Users2 struct {
//     Id        int
// 	Firstname string
// 	Middlename string
// 	Lastname  string
// 	Email string
// 	Phone string
// }
// var users2 = []Users2{
// 	Users2{Id: 1, Firstname: "Oliver", Middlename: "dfg", Lastname: "Queen", Email: "oliver@gmail.com", Phone:"1234567890"},
// 	Users2{Id: 2, Firstname: "Malcom", Middlename: "dfg",Lastname: "Merlyn", Email: "oliver@gmail.com", Phone:"1234567890"},
// 	Users2{Id:3, Firstname: "Ritik", Middlename: "dfg",Lastname:"Gupta", Email: "oliver@gmail.com", Phone:"1234567890"},
// }

func (v2 Version2Controller) Contains(x userobj2.Users) bool {
	for _, n := range users {
		if x.Id == n.Id {
			return true
		}
	}
	return false
}

func (v2 Version2Controller) PostUser(c *gin.Context) {

	var jsonData userobj2.Users

	// if c.BindJSON(&jsonData) == nil {
	// 	users2 = append(users2, jsonData)
	// 	c.JSON(200, gin.H{"status": "Uploaded Successful"})
	// }
	if c.BindJSON(&jsonData) == nil {
		// users = append(users, jsonData)
		err := models.Add(jsonData)
		if err == nil {
			c.JSON(200, gin.H{"status": "Uploaded Successful"})
		} else {
			c.JSON(200, gin.H{"status": err})
		}

	}

}

func (v2 Version2Controller) getIndex(id int) int {

	var index userobj2.Users
	var i int
	for i, index = range users2 {
		if id == index.Id {
			break
		}
	}
	return i

}

func (v2 Version2Controller) GetUsers(c *gin.Context) {

	// c.JSON(200, users)
	allusers, err := models.GetAll()
	if err == nil {
		c.JSON(200, allusers)
	} else {
		c.JSON(500, "Kuch gadbd ho raha hai!!!")
	}

}

func (v2 Version2Controller) GetUser(c *gin.Context) {

	id := c.Params.ByName("id")
	user_id, _ := strconv.Atoi(id)
	puser, err := models.GetById(user_id)

	if err != nil {
		c.JSON(500, gin.H{
			"response": "Record not found!",
		})
	} else {
		c.JSON(200, puser)
	}
	// id := c.Params.ByName("id")
	// user_id, _ := strconv.Atoi(id)

	// var puser = userobj2.Users{}

	// // switch(user_id){
	// // 	case 1:  puser = users[0]
	// // 			 break;
	// // 	case 2: puser = users[1]
	// // 			 break;
	// // 	case 3: puser = users[2]
	// // 			 break;
	// // 	// default: errordefault()
	// // }
	// puser = users2[v2.getIndex(user_id)]

	// fmt.Print(puser.Firstname)

	// content := gin.H{"id": user_id, "firstname": puser.Firstname, "middlename": puser.Middlename, "lastname": puser.Lastname, "email": puser.Email, "phone": puser.Phone}
	// c.JSON(200, content)

	// errordefault(){
	// content := gin.H{"error": "user with id#" + id + " not found"}
	// c.JSON(404, content)
}

func (v2 Version2Controller) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var jsonData userobj2.Users
	var err error
	if c.BindJSON(&jsonData) == nil {
		err = models.Update(id, jsonData)
	}
	c.JSON(200, gin.H{
		"status": err.Error(),
	})

}

func (v2 Version2Controller) DeleteUser(c *gin.Context) {
	//
	id := c.Params.ByName("id")
	user_id, _ := strconv.Atoi(id)

	err := models.Delete(user_id)
	if err == nil {
		c.JSON(200, gin.H{
			"status": "Deletion Successful!",
		})
	} else {
		c.JSON(500, gin.H{
			"status": "Record not found" + err.Error(),
		})
	}
}
