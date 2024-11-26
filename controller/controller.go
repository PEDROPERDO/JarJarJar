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
	result, err := utilities.OpenJSON("product.json")
	if err != nil {
		log.Fatalf("Fail : %v", err)
	}
	C.JSON(200, result)
}

func GetProductID(C *gin.Context) {
	// Get Product ID
	PID := C.Param("ProductID")
	allresult, err := utilities.OpenJSON("product.json")
	if err != nil {
		log.Fatalf("Fail : %v", err)
	}
	for _, result := range allresult {
		if result.ProductID == PID {
			C.JSON(200, result)
			return
		}
	}
	C.JSON(404, gin.H{"result": "Product Unavailable !"})
}