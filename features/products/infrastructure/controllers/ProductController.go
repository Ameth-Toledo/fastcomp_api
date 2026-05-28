package controllers

import (
	"fastcomp_api/core"
	"fastcomp_api/features/products/application"
	"fastcomp_api/features/products/domain/dto"
	"fastcomp_api/features/products/domain/entities"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	service *application.ProductService
}

func NewProductController(service *application.ProductService) *ProductController {
	return &ProductController{service: service}
}

func (pc *ProductController) Create(c *gin.Context) {
	var req dto.CreateProductRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var imageURL string
	fileHeader, err := c.FormFile("image")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer la imagen"})
			return
		}
		defer file.Close()

		imageURL, err = core.UploadImage(file, "fastcomp/products")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	product := &entities.Product{
		Brand:       req.Brand,
		Name:        req.Name,
		Category:    entities.Category(req.Category),
		Price:       req.Price,
		Stock:       req.Stock,
		Condition:   entities.Condition(req.Condition),
		Featured:    req.Featured,
		Description: req.Description,
		Specs:       req.Specs,
		ImageURL:    imageURL,
	}

	created, err := pc.service.Create(product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, created)
}

func (pc *ProductController) GetByID(c *gin.Context) {
	id := c.Param("id")
	product, err := pc.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetAll(c *gin.Context) {
	category := c.Query("category")
	featured := c.Query("featured")

	if featured == "true" {
		products, err := pc.service.GetFeatured()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
		return
	}

	if category != "" {
		products, err := pc.service.GetByCategory(entities.Category(category))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, products)
		return
	}

	products, err := pc.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) Update(c *gin.Context) {
	id := c.Param("id")
	var req dto.UpdateProductRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existing, err := pc.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	imageURL := existing.ImageURL
	fileHeader, err := c.FormFile("image")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer la imagen"})
			return
		}
		defer file.Close()

		imageURL, err = core.UploadImage(file, "fastcomp/products")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	product := &entities.Product{
		Brand:       req.Brand,
		Name:        req.Name,
		Category:    entities.Category(req.Category),
		Price:       req.Price,
		Stock:       req.Stock,
		Condition:   entities.Condition(req.Condition),
		Featured:    req.Featured,
		Description: req.Description,
		Specs:       req.Specs,
		ImageURL:    imageURL,
	}

	updated, err := pc.service.Update(id, product)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updated)
}

func (pc *ProductController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := pc.service.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Producto eliminado"})
}
