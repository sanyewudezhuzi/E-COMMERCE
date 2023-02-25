package serviceaddress

import (
	"context"

	"strconv"

	daoaddress "github.com/sanyewudezhuzi/E-COMMERCE/dao/dao_address"
	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"github.com/sanyewudezhuzi/E-COMMERCE/pkg/e"
	"github.com/sanyewudezhuzi/E-COMMERCE/serializer"
)

type AddressService struct {
	Name    string `form:"name" json:"name"`
	Phone   string `form:"phone" json:"phone"`
	Address string `form:"address" json:"address"`
}

func (service *AddressService) Create(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := daoaddress.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uId,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	err := addressDao.CreateAddress(address)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	addressDao = daoaddress.NewAddressDaoByDB(addressDao.DB)
	var addresses []*model.Address
	addresses, err = addressDao.ListAddressByUid(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.BuildAddresses(addresses),
		Msg:        e.GetMsg(code),
	}
}

func (service *AddressService) Show(ctx context.Context, aId string) serializer.Response {
	code := e.Success
	addressDao := daoaddress.NewAddressDao(ctx)
	addressId, _ := strconv.Atoi(aId)
	address, err := addressDao.GetAddressByAid(uint(addressId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.BuildAddress(address),
		Msg:        e.GetMsg(code),
	}
}

func (service *AddressService) List(ctx context.Context, uId uint) serializer.Response {
	code := e.Success
	addressDao := daoaddress.NewAddressDao(ctx)
	address, err := addressDao.ListAddressByUid(uId)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.BuildAddresses(address),
		Msg:        e.GetMsg(code),
	}
}

func (service *AddressService) Delete(ctx context.Context, aId string) serializer.Response {
	addressDao := daoaddress.NewAddressDao(ctx)
	code := e.Success
	addressId, _ := strconv.Atoi(aId)
	err := addressDao.DeleteAddressById(uint(addressId))
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	return serializer.Response{
		StatusCode: code,
		Msg:        e.GetMsg(code),
	}
}

func (service *AddressService) Update(ctx context.Context, uid uint, aid string) serializer.Response {
	code := e.Success

	addressDao := daoaddress.NewAddressDao(ctx)
	address := &model.Address{
		UserID:  uid,
		Name:    service.Name,
		Phone:   service.Phone,
		Address: service.Address,
	}
	addressId, _ := strconv.Atoi(aid)
	addressDao.UpdateAddressById(uint(addressId), address)
	addressDao = daoaddress.NewAddressDaoByDB(addressDao.DB)
	var addresses []*model.Address
	addresses, err := addressDao.ListAddressByUid(uid)
	if err != nil {
		code = e.Error
		return serializer.Response{
			StatusCode: code,
			Msg:        e.GetMsg(code),
			Error:      err,
		}
	}
	return serializer.Response{
		StatusCode: code,
		Data:       serializer.BuildAddresses(addresses),
		Msg:        e.GetMsg(code),
	}
}
