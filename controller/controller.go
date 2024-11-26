package controller

import (
	"github.com/gin-gonic/gin"
	// "log"
)

func MainFile(C *gin.Context) {
	// Main File
	C.JSON(200, gin.H{"result": "Hallo ! Simple Application !"})
}

func GetAllProduct(C *gin.Context) {
	// Get All Product
	C.JSON(200, gin.H{"result": "Hallo ! Our Products !"})
}

func GetProductID(C *gin.Context) {
	// Get All Product
	PID := C.Param("ProductID")
	C.JSON(200, gin.H{"result": PID})
}