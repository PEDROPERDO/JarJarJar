package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"SimpleGolang/utilities"
)

func MainFile(C *gin.Context) {
	// Main File
	C.JSON(200, gin.H{"result": "Hallo ! Simple Application !"})
}

func GetAllProduct(C *gin.Context) {
	// Get All Product
	result, err := utilities.OpenJSON("productfile.json")
	if err != nil {
		log.Fatalf("Fail : %v", err)
	}
	C.JSON(200, result)
}

func GetProductID(C *gin.Context) {
	// Get All Product
	PID := C.Param("ProductID")
	C.JSON(200, gin.H{"result": PID})
}