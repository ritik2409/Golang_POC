package main

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type Credentials struct {
	UserName string
	Password string
}

var identityKey = "id"

func helloHandler(c *gin.Context) {
	log.Println("Sab badiya 4")
	claims := jwt.ExtractClaims(c)
	log.Println("Sab badiya1")
	user, _ := c.Get(identityKey)
	log.Println("Sab badiya 2")
	c.JSON(200, gin.H{
		"userID":   claims["id"],
		"userName": user.(*User).UserName,
		"text":     "Hello World.",
	})

	// reqtoken := c.Request.Header.Get("Authorization")
	// splitoken := strings.Split(reqtoken, "Bearer")
	// reqtoken = splitoken[1]
	// log.Println(reqtoken)

	// SecretKey := "this_is_random"

	// token, err := jwt.ParseWithClaims(reqToken, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte(SecretKey), nil
	// })
	// if err == nil && token.Valid {
	// 	fmt.Println("valid token")
	// 	return true
	// } else {
	// 	fmt.Println("invalid token")
	// 	return false
	// }
}

// User demo
type User struct {
	UserName  string
	FirstName string
	LastName  string
}

func main() {

	r := gin.Default()
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:            "test zone",
		SigningAlgorithm: "HS256",
		Key:              []byte("this_is_random"),
		Timeout:          time.Hour,
		MaxRefresh:       time.Hour,
		IdentityKey:      identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims["id"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Credentials
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.UserName
			password := loginVals.Password

			if (userID == "admin" && password == "admin") || (userID == "test" && password == "test") {
				return &User{
					UserName:  userID,
					LastName:  "Bo-Yi",
					FirstName: "Wu",
				}, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok && v.UserName == "admin" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	// authMiddleware.SigningAlgorithm = "RS256"
	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	// Refresh time can be longer than token timeout
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", helloHandler)
		log.Println("Sab Badiya 3")
	}

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

// func getdetails(c *gin.Context) {
// 	c.JSON(200, gin.H{
// 		"Status":        "verified",
// 		"CustomerName":  "Ritik",
// 		"CustomerEmail": "ritik@zestmoney.in",
// 		"CustomerPhone": "1234567890",
// 	})
// }
