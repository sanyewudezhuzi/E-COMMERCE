package servicecarousel

import (
	"context"

	daocarousel "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_carousel"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/e"
	"github.com/sanyewudezhuzi/E-COMMERCE/serializer"
)

type CarouselService struct {
}

func (s *CarouselService) Show(ctx context.Context) serializer.Response {
	carouselDao := daocarousel.NewCarouselDao(ctx)
	code := e.Success
	carousels, err := carouselDao.ShowCarousel()
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	res := serializer.BuildListResponse(serializer.BuildCarousels(carousels), len(carousels))
	res.StatusCode = code
	res.Msg = e.GetMsg(code)
	return res
}
