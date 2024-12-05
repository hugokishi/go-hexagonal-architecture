package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hugokishi/hexagonal-go/internal/core/models"
	"github.com/hugokishi/hexagonal-go/internal/core/services"
)

type ProductHandler struct {
	svc services.ProductService
}

func NewProductHandler(ProductService services.ProductService, routerGroup *gin.RouterGroup) {
	h := &ProductHandler{svc: ProductService}

	routerGroup.GET("/products", h.Save)
}

func (h *ProductHandler) Save(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.svc.Save(ctx, product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"product": product})
}
