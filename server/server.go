package server

func Init() {
	// config := config.GetConfig()
	r := NewRouter()
	// r.Run(config.GetString("server.port"))
	r.Run(":8080")
}