package repository

import (
	"carrinho-api-gorm/model"

	"github.com/jinzhu/gorm"
)

const (
	tableNameCarrinho = "carrinho"
)

//ICarrinhoRepository ...
type ICarrinhoRepository interface {
	CarrinhoByClienteID(clienteID string) (*model.Carrinho, error)
}

//CarrinhoRepository ...
type CarrinhoRepository struct {
	db *gorm.DB
}

//NewCarrinhoRepository ...
func NewCarrinhoRepository(db *gorm.DB) ICarrinhoRepository {
	return &CarrinhoRepository{db}
}

//CarrinhoByClienteID ...
func (l *CarrinhoRepository) CarrinhoByClienteID(clienteID string) (*model.Carrinho, error) {
	var err error
	carrinho := model.Carrinho{}

	err = l.db.Table(tableNameCarrinho).Where("ClienteId = ?", clienteID).Find(&carrinho).Error

	if err != nil {
		return &model.Carrinho{}, err
	}

	return &carrinho, nil
}
