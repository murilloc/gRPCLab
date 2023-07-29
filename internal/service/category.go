package service

import (
	"context"
	"github.com/devfullcycke/14-gRPC/internal/database"
	"github.com/devfullcycke/14-gRPC/internal/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{CategoryDB: categoryDB}
}

func (c *CategoryService) AddCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {

	category, err := c.CategoryDB.AddCategory(in.Name, in.Description)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while trying to create category")
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return &pb.CategoryResponse{
		Category: categoryResponse,
	}, nil

}
