package products

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	commonError "training/go-products/common-error"
	"training/go-products/persistance"
	"training/go-products/validation"
)

func checkAlreadyExists(product Product) error {
	log.Println(`product name to be inserted`, product.Name)
	var existingProduct Product
	results := persistance.DB.First(&existingProduct, "name = ? ", product.Name)
	if results.RowsAffected != 0 {
		return &commonError.CommonError{
			Code:    "PRODUCT_ALREADY_EXISTS",
			Message: "product " + product.Name + " already exists",
		}
	}
	return nil
}

func validateProduct(product Product) error {
	err := validation.Validate.Struct(product)
	return err
}

func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, Product{
		Name:        "p",
		Description: "ABC",
		Price:       34.5,
	})
}

func GetAllProductsController(c *gin.Context) {
	pagination := persistance.GeneratePaginationFromRequest(c)
	var product Product
	allProducts, err := GetAllProducts(&product, &pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": allProducts,
	})

}

func GetAllProducts(product *Product, pagination *persistance.Pagination) (*[]Product, error) {
	var products []Product
	offset := (pagination.Page - 1) * pagination.Limit
	queryBuilder := persistance.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort)
	result := queryBuilder.Model(&Product{}).Where(product).Find(&products)
	if result.Error != nil {
		msg := result.Error
		return nil, msg
	}
	return &products, nil
}

func GetById(context *gin.Context) {
	var product Product
	persistance.DB.First(&product, context.Param("id"))
	context.JSON(http.StatusOK, product)
}

func SaveProduct(c *gin.Context) {
	var product Product
	err := c.BindJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	err = validateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, &commonError.CommonError{
			Code:    "VALIDATION_FAILED",
			Message: err.Error(),
		})
		return
	}
	err = checkAlreadyExists(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	persistance.DB.Create(&product)
	c.JSON(http.StatusOK, product)
}
