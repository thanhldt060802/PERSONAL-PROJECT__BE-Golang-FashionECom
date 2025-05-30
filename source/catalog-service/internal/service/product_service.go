package service

import (
	"context"
	"fmt"
	"thanhldt060802/internal/dto"
	"thanhldt060802/internal/model"
	"thanhldt060802/internal/repository"
	"time"
)

type productService struct {
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
	brandRepository    repository.BrandRepository
}

type ProductService interface {
	// Integrate with Elasticsearch
	GetAllProducts(ctx context.Context) ([]dto.ProductView, error)

	// Main features
	GetProductById(ctx context.Context, reqDTO *dto.GetProductByIdRequest) (*dto.ProductView, error)
	CreateProduct(ctx context.Context, reqDTO *dto.CreateProductRequest) error
	UpdateProductById(ctx context.Context, reqDTO *dto.UpdateProductByIdRequest) error
	DeleteProductById(ctx context.Context, reqDTO *dto.DeleteProductByIdRequest) error

	// Elasticsearch integration features
	// GetProducts()
}

func NewProductService(productRepository repository.ProductRepository, categoryRepository repository.CategoryRepository, brandRepository repository.BrandRepository) ProductService {
	return &productService{
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
		brandRepository:    brandRepository,
	}
}

//
//
// Integrate with Elasticsearch
// ######################################################################################

func (productService *productService) GetAllProducts(ctx context.Context) ([]dto.ProductView, error) {
	products, err := productService.productRepository.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("query products from postgresql failed: %s", err.Error())
	}

	categories, err := productService.categoryRepository.Get(ctx, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("query categories from postgresql failed: %s", err.Error())
	}

	brands, err := productService.brandRepository.Get(ctx, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("query brands from postgresql failed: %s", err.Error())
	}

	return dto.ToListProductView(products, categories, brands), nil
}

//
//
// Main features
// ######################################################################################

func (productService *productService) GetProductById(ctx context.Context, reqDTO *dto.GetProductByIdRequest) (*dto.ProductView, error) {
	foundProduct, err := productService.productRepository.GetById(ctx, reqDTO.Id)
	if err != nil {
		return nil, fmt.Errorf("id of product is not valid: %s", err.Error())
	}

	foundCategory, _ := productService.categoryRepository.GetById(ctx, foundProduct.CategoryId)

	foundBrand, _ := productService.brandRepository.GetById(ctx, foundProduct.BrandId)

	return dto.ToProductView(foundProduct, foundCategory, foundBrand), nil
}

func (productService *productService) CreateProduct(ctx context.Context, reqDTO *dto.CreateProductRequest) error {
	if _, err := productService.categoryRepository.GetById(ctx, reqDTO.Body.CategoryId); err != nil {
		return fmt.Errorf("id of category not found")
	}
	if _, err := productService.brandRepository.GetById(ctx, reqDTO.Body.BrandId); err != nil {
		return fmt.Errorf("id of brand not found")
	}

	newProduct := model.Product{
		Name:               reqDTO.Body.Name,
		Description:        reqDTO.Body.Description,
		Sex:                reqDTO.Body.Sex,
		Price:              reqDTO.Body.Price,
		DiscountPercentage: reqDTO.Body.DiscountPercentage,
		Stock:              reqDTO.Body.Stock,
		ImageURL:           reqDTO.Body.ImageURL,
		CategoryId:         reqDTO.Body.CategoryId,
		BrandId:            reqDTO.Body.BrandId,
	}
	if err := productService.productRepository.Create(ctx, &newProduct); err != nil {
		return fmt.Errorf("insert product to postgresql failed: %s", err.Error())
	}

	// Missing->SyncCreatingToElasticsearch

	return nil
}

func (productService *productService) UpdateProductById(ctx context.Context, reqDTO *dto.UpdateProductByIdRequest) error {
	foundProduct, err := productService.productRepository.GetById(ctx, reqDTO.Id)
	if err != nil {
		return fmt.Errorf("id of product is not valid: %s", err.Error())
	}

	if reqDTO.Body.Name != nil {
		foundProduct.Name = *reqDTO.Body.Name
	}
	if reqDTO.Body.Description != nil {
		foundProduct.Description = *reqDTO.Body.Description
	}
	if reqDTO.Body.Sex != nil {
		foundProduct.Sex = *reqDTO.Body.Sex
	}
	if reqDTO.Body.Price != nil {
		foundProduct.Price = *reqDTO.Body.Price
	}
	if reqDTO.Body.DiscountPercentage != nil {
		foundProduct.DiscountPercentage = *reqDTO.Body.DiscountPercentage
	}
	if reqDTO.Body.Stock != nil {
		foundProduct.Stock = *reqDTO.Body.Stock
	}
	if reqDTO.Body.ImageURL != nil {
		foundProduct.ImageURL = *reqDTO.Body.ImageURL
	}
	if reqDTO.Body.CategoryId != nil {
		if _, err := productService.categoryRepository.GetById(ctx, *reqDTO.Body.CategoryId); err != nil {
			return fmt.Errorf("id of category not found")
		}
		foundProduct.CategoryId = *reqDTO.Body.CategoryId
	}
	if reqDTO.Body.BrandId != nil {
		if _, err := productService.brandRepository.GetById(ctx, *reqDTO.Body.BrandId); err != nil {
			return fmt.Errorf("id of brand not found")
		}
		foundProduct.CategoryId = *reqDTO.Body.BrandId
	}
	foundProduct.UpdatedAt = time.Now().UTC()

	if err := productService.productRepository.Update(ctx, foundProduct); err != nil {
		return fmt.Errorf("update product on postgresql failed: %s", err.Error())
	}

	// Missing->SyncUpdatingToElasticsearch

	return nil
}

func (productService *productService) DeleteProductById(ctx context.Context, reqDTO *dto.DeleteProductByIdRequest) error {
	if _, err := productService.productRepository.GetById(ctx, reqDTO.Id); err != nil {
		return fmt.Errorf("id of product is not valid")
	}

	if err := productService.productRepository.DeleteById(ctx, reqDTO.Id); err != nil {
		return fmt.Errorf("delete product from postgresql failed: %s", err.Error())
	}

	// Missing->SyncDeletingToElasticsearch

	return nil
}
