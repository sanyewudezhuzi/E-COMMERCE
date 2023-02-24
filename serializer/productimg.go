package serializer

import (
	"github.com/sanyewudezhuzi/E-COMMERCE/conf"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
)

type ProductImg struct {
	ProductID uint   `json:"product_id"`
	ImgPath   string `json:"img_path"`
}

func BuildProductImg(productimg model.ProductImg) ProductImg {
	return ProductImg{
		ProductID: productimg.ID,
		ImgPath:   conf.Host + conf.HttpPort + conf.ProductPath + productimg.ImgPath,
	}
}

func BuildProductImgs(productimgs []model.ProductImg) []ProductImg {
	var ProductImgs []ProductImg = make([]ProductImg, len(productimgs))
	for k, v := range productimgs {
		ProductImgs[k] = BuildProductImg(v)
	}
	return ProductImgs
}
