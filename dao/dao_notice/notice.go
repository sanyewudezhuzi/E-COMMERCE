package daonotice

import (
	"context"

	"github.com/sanyewudezhuzi/E-COMMERCE/model"
	"gorm.io/gorm"
)

type NoticeDao struct {
	*gorm.DB
}

func NewNoticeDao(ctx context.Context) *NoticeDao {
	return &NoticeDao{model.NewDBClient(ctx)}
}

func NewNoticeDaoByDB(db *gorm.DB) *NoticeDao {
	return &NoticeDao{db}
}

func (dao *NoticeDao) GetNoticeByID(id uint) (*model.Notice, error) {
	var notice model.Notice
	err := dao.DB.Model(&model.Notice{}).Where("id = ?", id).First(&notice).Error
	return &notice, err
}
