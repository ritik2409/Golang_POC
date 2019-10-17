package models

import (
	"errors"
	"log"
	"project/db"
	"project/dto"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// var db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/user")

// type User struct {
// 	Id          int
// 	FirstName   string `db:"first_name" form:"first_name"`
// 	MiddleName  string `db:"middle_name" form:"middle_name"`
// 	LastName    string `db:"last_name" form:"last_name"`
// 	Email       string `db:"email" form:"email"`
// 	MobilePhone string `db:"mobile_phone" form:"mobile_phone"`
// }

func GetAll() ([]dto.Users, error) {
	db := db.GetDB()
	var (
		// user  dto.Users
		users []dto.Users
	)
	if err := db.Find(&users).Error; err != nil {
		log.Println("Dikkat toh yahin hai...", err)
		return nil, err
	}

	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	// if err == nil {
	// 	return users, nil
	// }

	// for rows.Next() {
	// 	rows.Scan(&user.Id, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.MobilePhone)
	// 	users = append(users, user)
	// }

	// defer rows.Close()
	// c.JSON(http.StatusOK, users)
	log.Println("Sab badiya chal raha...bus ab records dikhata hu")
	return users, nil
}

func Add(crecord dto.Users) error {
	db := db.GetDB()
	var nuser dto.Users
	nuser.Id = crecord.Id
	var err error
	if db.First(&nuser, nuser.Id).RecordNotFound() {
		db.Create(&crecord)
		err = nil
		log.Println("Save toh ho gaya..")
	} else {
		log.Println("Sab badiya1")
		db.First(&nuser, crecord.Id)
		log.Println("Sab badiya2")
		db.Model(&crecord).UpdateColumns(dto.Users{Firstname: crecord.Firstname, Middlename: crecord.Middlename,
			Lastname: crecord.Lastname, Phone: crecord.Phone, Email: crecord.Email})
		log.Println("Sab badiya3")
		err = errors.New("Record already exists. Hence, Updated!")
		log.Println("Nhi ho raha save...", err)
	}

	// db := db.GetDB()
	// Id := c.PostForm("id")
	// FirstName := c.PostForm("first_name")
	// MiddleName := c.PostForm("middle_name")
	// LastName := c.PostForm("last_name")
	// Email := c.PostForm("email")
	// MobilePhone := c.PostForm("mobile_phone")

	// if FirstName == "" || Email == "" {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": fmt.Sprintf("Please fill all mandatory field"),
	// 	})
	// return
	return err
}

// if err != nil {
// 	panic(err)
// }

// stmt, err := db.Prepare("insert into details (id,first_name, middle_name, last_name, email, mobile_phone) values(?,?,?,?,?,?);")

// if err != nil {
// 	fmt.Print(err.Error())

// }

// _, err = stmt.Exec(Id, FirstName, MiddleName, LastName, Email, MobilePhone)
// if err != nil {
// 	fmt.Print(err.Error())
// }

// defer stmt.Close()

// c.JSON(http.StatusOK, gin.H{
//     "message": fmt.Sprintf("successfully created"), new User{Id:5,Firstname:"fghj",Lastname:"hj",Middlename:"fghj",Email:"fghj",MobilePh}
// })
// return nil
// }

func GetById(userId int) (dto.Users, error) {
	db := db.GetDB()
	var user dto.Users

	// id := c.Params.ByName("id")
	// userId, _ := strconv.Atoi(id)
	// row := db.QueryRow("select id, first_name, middle_name, last_name, email, mobile_phone from details where id = ?;", id)

	if db.First(&user, userId).RecordNotFound() {
		log.Println("Aisa koi record nhi...")
		return user, errors.New("record not found")
	}
	// err = row.Scan(&user.Id, &user.FirstName, &user.MiddleName, &user.LastName, &user.Email, &user.MobilePhone)
	// if err != nil {
	// 	c.JSON(http.StatusOK, nil)
	// } else {
	// 	c.JSON(http.StatusOK, user)
	// }
	return user, nil
}

func Update(id int, urecord dto.Users) error {

	db := db.GetDB()
	// var nuser dto.Users
	// nuser.Id = urecord.Id
	log.Println(urecord)
	var nrecord dto.Users
	var err error
	if db.First(&nrecord, id).RecordNotFound() {
		db.Create(&urecord)
		err = errors.New("No record found...New User Added to the Database!")
		log.Println("Naya record aaya hai...")
	} else {
		err = errors.New("Updated Successfully!")
		db.First(&nrecord, id)
		log.Println(nrecord)
		// nrecord.Firstname = urecord.Firstname
		// nrecord.Middlename = urecord.Middlename
		// nrecord.Lastname = urecord.Lastname
		// nrecord.Phone = urecord.Phone
		// nrecord.Email = urecord.Email
		// db.Save(&nrecord)
		log.Println(urecord)

		db.Model(&urecord).UpdateColumns(dto.Users{Firstname: urecord.Firstname, Middlename: urecord.Middlename,
			Lastname: urecord.Lastname, Phone: urecord.Phone, Email: urecord.Email})

		// copy(&nrecord, &urecord)
		// copyRecord()
		// nrecord = urecord
		// db.Save(&nrecord)

		log.Println("Kardiya update...", err)
	}

	return err
}

func Delete(id int) error {
	var duser dto.Users
	duser.Id = id

	db := db.GetDB()
	if db.First(&duser, id).RecordNotFound() {
		log.Println("Aisa koi record nhi...")
		return errors.New("record not found")
	} else {
		db.Delete(&duser)
	}

	return nil
	// id := c.Param("id")
	// stmt, err := db.Prepare("delete from details where id= ?;")

	// if err != nil {
	// 	fmt.Print(err.Error())
	// }
	// _, err = stmt.Exec(id)

	// if err != nil {
	// 	fmt.Print(err.Error())
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": fmt.Sprintf("Successfully deleted user with ID : %s", id),
	// })
}

// func createTable() {
// 	// db := db.GetDB()
// 	// stmt, err := db.Prepare("CREATE TABLE details (id int NOT NULL AUTO_INCREMENT,first_name varchar(40), middle_name varchar(40), last_name varchar(40), email varchar(60), mobile_phone varchar(15), PRIMARY KEY (id));")
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// }
// 	// _, err = stmt.Exec()
// 	// if err != nil {
// 	// 	fmt.Println(err.Error())
// 	// } else {
// 	// 	fmt.Println("Table is successfully created")
// 	// }
// }

// func db() {
//     createTable()

//     if err != nil {
//         fmt.Print(err.Error())
//     }
//     defer db.Close()

//     err = db.Ping()
//     if err != nil {
//         fmt.Print(err.Error())
//     }

//     router := gin.Default()
//     router.GET("/api/user/:id", getById)
//     router.GET("/api/users", getAll)
//     router.POST("/api/user", add)
//     router.DELETE("/api/user/:id", delete)
//     router.Run(":8000")
// }
