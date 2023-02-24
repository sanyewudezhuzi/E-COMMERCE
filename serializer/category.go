package serializer

import "github.com/sanyewudezhuzi/E-COMMERCE/model"

type Category struct {
	ID           uint   `json:"id"`
	CategoryName string `json:"category_name"`
	CreateAt     int64  `json:"create_at"`
}

func BuildCategory(item *model.Category) Category {
	return Category{
		ID:           item.ID,
		CategoryName: item.CategoryName,
		CreateAt:     item.CreatedAt.Unix(),
	}
}

func BuildCategorys(item []model.Category) []Category {
	var Category []Category = make([]Category, len(item))
	for k, v := range item {
		Category[k] = BuildCategory(&v)
	}
	return Category
}
