package model

import (
	"strconv"

	"github.com/sanyewudezhuzi/E-COMMERCE/cache"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name          string
	CategoryID    uint
	Title         string
	Info          string
	ImgPath       string
	Price         string
	DiscountPrice string
	OnSale        bool `gorm:"default:false"`
	Num           int
	BossID        uint
	BossName      string
	BossAvatar    string
}

func (product *Product) View() int64 {
	countStr := cache.RedisClient.Get(cache.RedisClient.Context(), cache.ProductViewKey(product.ID))
	count, _ := strconv.Atoi(countStr.String())
	return int64(count)
}

func (product *Product) AddView() {
	// 增加商品点击数
	cache.RedisClient.Incr(cache.RedisClient.Context(), cache.ProductViewKey(product.ID))
	cache.RedisClient.ZIncrBy(cache.RedisClient.Context(), cache.RankKey, 1, strconv.Itoa(int(product.ID)))
}
