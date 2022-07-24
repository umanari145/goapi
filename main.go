package main

import (
	"net"
)

func init() {
	//err := godotenv.Load()
	//if err != nil {
	//	logrus.Fatal("Error loading .env")
	//}

	//logrus.SetLevel(logrus.DebugLevel)
	//logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	listen, err := net.Listen("tcp", "localhost:12345")
	if err != nil {
		return errors.WithStack(err)
	}
}
