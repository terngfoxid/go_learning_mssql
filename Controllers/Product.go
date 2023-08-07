package Controllers

import (
	"fmt"
	"go-mssql/Models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var product []Models.Product
	product, err := Models.GetAllProducts()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func CreateProduct(c *gin.Context) {
	var product Models.Product
	c.BindJSON(&product)
	status, err := Models.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Printf("Status:%d", status)
		fmt.Println("")
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "status ok",
			"Detail":  "Product:" + product.Name + " is created"})
	}
}

func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var product Models.Product
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&product, id)
	if err != nil {
		c.JSON(http.StatusNotFound, product)
	}
	var bufferpro Models.Product
	bufferpro.Id = product.Id
	c.BindJSON(&product)
	product.Id = bufferpro.Id
	status, err := Models.UpdateProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Printf("Status:%d", status)
		fmt.Println("")
		c.JSON(http.StatusOK, product)
	}
}

func DeleteProduct(c *gin.Context) {
	var product Models.Product
	id := c.Params.ByName("id")
	status, err := Models.DeleteProduct(&product, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Printf("Status:%d", status)
		fmt.Println("")
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "status ok",
			"Detail":  "Product id " + id + " is deleted"})
	}
}
