package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"

	"project/pkg/fibonacci"
	"project/pkg/grpcserver"
	"project/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	//HTTP REST
	router := gin.Default()
	go router.GET("/numbers", getNumbers)
	go router.Run("localhost:8080")

	//GRPC
	s := grpc.NewServer()
	srv := &grpcserver.GRPCServer{}
	proto.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}

}

func getNumbers(c *gin.Context) {

	firstNumber := c.Query("start")
	secondNumber := c.Query("end")
	intFirstNumber, err := strconv.Atoi(firstNumber)
	if err != nil {
		fmt.Println(err)
	}
	intSecondNumber, err := strconv.Atoi(secondNumber)

	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, fibonacci.Culc(intFirstNumber, intSecondNumber))
}
