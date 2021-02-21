package repository

import (
	"carrinho-api-gorm/model"

	"github.com/jinzhu/gorm"
)

const (
	tableNameItensCarrinho = "itenscarrinho"
)

//IItensCarrinhoRepository ...
type IItensCarrinhoRepository interface {
	ItensCarrinhoByCarrinhoID(carrinhoID string) (*[]model.ItensCarrinho, error)
}

//ItensCarrinhoRepository ...
type ItensCarrinhoRepository struct {
	db *gorm.DB
}

//NewItensCarrinhoRepository ...
func NewItensCarrinhoRepository(db *gorm.DB) IItensCarrinhoRepository {
	return &ItensCarrinhoRepository{db}
}

//ItensCarrinhoByCarrinhoID ...
func (l *ItensCarrinhoRepository) ItensCarrinhoByCarrinhoID(carrinhoID string) (*[]model.ItensCarrinho, error) {
	var err error
	itensCarrinho := []model.ItensCarrinho{}

	err = l.db.Table(tableNameItensCarrinho).Where("CarrinhoId = ?", carrinhoID).Find(&itensCarrinho).Error

	if err != nil {
		return &[]model.ItensCarrinho{}, err
	}

	return &itensCarrinho, nil
}
