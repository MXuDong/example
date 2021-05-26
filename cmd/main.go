package main

import (
	"github.com/sirupsen/logrus"
	"github.io/MXuDong/example/internal/server"
)

// ====================
// @author: MXuDong
// project for:
//   The project for mock the application in product environment, and
//   use it to provide some func or handle to mock the action of
//   application.
// Create at: 2021-04-30
// ====================

// main function, the application run here
func main() {
	go func() {
		err := server.TcpServerStart()
		if err != nil {
			logrus.Error(err)
		}
	}()
	server.Run()
}
