package controller

import (
	userobj "project/dto"
	"project/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Version1Controller struct{}

var users []userobj.Users

// var mod_users = new(models.User)

// func (v1 Version1Controller) Init() {
// 	// u := new(dto.userobj)
// 	users = userobj.Init()
// }

// type Users1 struct {
//     Id        int
// 	Firstname string
// 	Middlename string
// 	Lastname  string
// 	Email string
// 	Phone string
// }
// var users1 = []Users1{
// 	Users1{Id: 1, Firstname: "Oliver", Middlename: "dfg", Lastname: "Queen", Email: "oliver@gmail.com", Phone:"1234567890"},
// 	Users1{Id: 2, Firstname: "Malcom", Middlename: "dfg",Lastname: "Merlyn", Email: "oliver@gmail.com", Phone:"1234567890"},
// 	Users1{Id:3, Firstname: "Ritik", Middlename: "dfg",Lastname:"Gupta", Email: "oliver@gmail.com", Phone:"1234567890"},
// }

func (v1 Version1Controller) Contains(x userobj.Users) bool {
	for _, n := range users {
		if x.Id == n.Id {
			return true
		}
	}
	return false
}

func (v1 Version1Controller) PostUser(c *gin.Context) {

	var jsonData userobj.Users

	// mod_users.add(c)

	if c.BindJSON(&jsonData) == nil {
		// users = append(users, jsonData)
		err := models.Add(jsonData)
		if err == nil {
			c.JSON(200, gin.H{"status": "Uploaded Successful"})
		} else {
			c.JSON(200, gin.H{"status": err.Error()})
		}

	}

}

func (v1 Version1Controller) getIndex(id int) int {

	var index userobj.Users
	var i int
	for i, index = range users {
		if id == index.Id {
			break
		}
	}
	return i
}

func (v1 Version1Controller) GetUsers(c *gin.Context) {

	allusers, err := models.GetAll()
	if err == nil {
		c.JSON(200, allusers)
	} else {
		c.JSON(500, gin.H{"message": "Kuch gadbd ho raha hai!!!"})
	}

}
func (v1 Version1Controller) GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	user_id, _ := strconv.Atoi(id)
	puser, err := models.GetById(user_id)

	if err != nil {
		c.JSON(500, gin.H{
			"response": "Record not found!" + err.Error(),
		})
	} else {
		c.JSON(200, puser)
	}

	// var puser userobj.Users
	// switch(user_id){
	// 	case 1:  puser = users[0]
	// 			 break;
	// 	case 2: puser = users[1]
	// 			 break;
	// 	case 3: puser = users[2]
	// 			 break;
	// 	// default: errordefault()
	// }
	// copy(puser,users1[v1.getIndex(user_id)])
	// puser = users[v1.getIndex(user_id)]

	// fmt.Print(puser.Firstname)

	// content := gin.H{"id": user_id, "firstname": puser.Firstname, "middlename": puser.Middlename, "lastname": puser.Lastname, "email": puser.Email, "phone": puser.Phone}
	// c.JSON(200, content)

	// errordefault(){
	// content := gin.H{"error": "user with id#" + id + " not found"}
	// c.JSON(404, content)
}

func (v1 Version1Controller) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	var jsonData userobj.Users
	var err error
	if c.BindJSON(&jsonData) == nil {
		err = models.Update(id, jsonData)
	}
	c.JSON(200, gin.H{
		"status": err.Error(),
	})

}

func (v1 Version1Controller) DeleteUser(c *gin.Context) {

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
	// var jsonData userobj.Users
	// var list []Users
	// copy(list,users)
	// var mod_users []userobj.Users
	// // if c.BindJSON(&jsonData) == nil{
	// for i, cuser := range users {
	// 	if user_id != cuser.Id {
	// 		mod_users = append(mod_users, users[i])
	// 	}
	// 	// }
	// 	copy(users, mod_users)
	// 	// slice := users[0:jsonData.Id-1]

	// 	c.JSON(200, gin.H{"status": "Deleted Successful"})
	// }
}
