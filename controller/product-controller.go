package controller

import (
	"api-go/model"
	"api-go/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(usecase usecase.ProductUseCase) productController {
	return productController{
		productUseCase: usecase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	insertProduct, err := p.productUseCase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, insertProduct)
}

func (p *productController) GetProductById(ctx *gin.Context) {
	productId, errResp := parseProductID(ctx)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)
	if product == nil {
		response := model.Response{Message: "Produto nao encontrado"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (p *productController) DeleteProduct(ctx *gin.Context) {
	productId, errResp := parseProductID(ctx)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	product, err := p.productUseCase.GetProductById(productId)
	if product == nil {
		response := model.Response{Message: "Produto nao encontrado"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	err = p.productUseCase.DeleteProduct(productId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (p *productController) UpdateProduct(ctx *gin.Context) {
	productId, errResp := parseProductID(ctx)
	if errResp != nil {
		ctx.JSON(http.StatusBadRequest, errResp)
		return
	}

	productToUpdate, err := p.productUseCase.GetProductById(productId)
	if productToUpdate == nil {
		response := model.Response{Message: "Produto nao encontrado"}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	var product model.Product
	err = ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	productToUpdate.Name = product.Name
	productToUpdate.Price = product.Price

	updatedProduct, err := p.productUseCase.UpdateProduct(*productToUpdate)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, updatedProduct)
}

func parseProductID(ctx *gin.Context) (int, *model.Response) {
	id := ctx.Param("id")
	if id == "" {
		return 0, &model.Response{Message: "Id do produto é obrigatório"}
	}

	productId, err := strconv.Atoi(id)
	if err != nil {
		return 0, &model.Response{Message: "Id do produto deve ser um número"}
	}

	return productId, nil
}
