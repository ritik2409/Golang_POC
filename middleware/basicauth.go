package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_poc")

type Credentials struct {
	UserName string
	Password string
}

func basicauth() {

	var accounts []Credentials
	// var account Account
	db, _ := gorm.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/golang_poc")
	db.Debug().AutoMigrate(&Credentials{})
	// rows, err := db.Query("select * from golang_poc.credentials")
	log.Println("Abhi tak toh thik hai...")
	if err := db.Find(&accounts).Error; err != nil {
		log.Println("Dikkat toh yahin hai...", err)
	}

	log.Println(accounts)
	// if err == nil{
	// for rows.Next() {
	//     rows.Scan(&account.UserName, &account.Password)
	//     accounts = append(accounts,account)
	// }}

	var key_value map[string]string

	key_value = make(map[string]string)
	// x := make(map[string]string)

	for i := range accounts {
		// x[accounts[i].UserName] = accounts[i].Password
		// key_value = append(key_value,x[accounts[i].UserName])
		key_value[accounts[i].UserName] = accounts[i].Password
	}

	router := gin.Default()
	authorized := router.Group("/", gin.BasicAuth(key_value))

	authorized.GET("/secret", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"secret": "This is the secret!!",
		})
	})
	router.Run(":8080")

}
