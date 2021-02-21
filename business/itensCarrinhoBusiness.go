package business

import (
	"carrinho-api-gorm/model"
	"carrinho-api-gorm/repository"
)

//IItensCarrinhoBusiness ...
type IItensCarrinhoBusiness interface {
	ItensCarrinhoByCarrinhoID(carrinhoID string) (*[]model.ItensCarrinho, error)
}

//ItensCarrinhoBusiness ...
type ItensCarrinhoBusiness struct {
	ItensCarrinhoRepository repository.IItensCarrinhoRepository
}

//NewItensCarrinhoBusiness ...
func NewItensCarrinhoBusiness(ItensCarrinhoRepository repository.IItensCarrinhoRepository) IItensCarrinhoBusiness {
	return &ItensCarrinhoBusiness{ItensCarrinhoRepository}
}

//ItensCarrinhoByCarrinhoID ...
func (l *ItensCarrinhoBusiness) ItensCarrinhoByCarrinhoID(carrinhoID string) (*[]model.ItensCarrinho, error) {
	return l.ItensCarrinhoRepository.ItensCarrinhoByCarrinhoID(carrinhoID)
}
