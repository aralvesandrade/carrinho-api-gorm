package business

import (
	"carrinho-api-gorm/model"
	"carrinho-api-gorm/repository"
)

//ICarrinhoBusiness ...
type ICarrinhoBusiness interface {
	CarrinhoByClienteID(clienteID string) (*model.Carrinho, error)
}

//CarrinhoBusiness ...
type CarrinhoBusiness struct {
	carrinhoRepository repository.ICarrinhoRepository
}

//NewCarrinhoBusiness ...
func NewCarrinhoBusiness(carrinhoRepository repository.ICarrinhoRepository) ICarrinhoBusiness {
	return &CarrinhoBusiness{carrinhoRepository}
}

//CarrinhoByClienteID ...
func (l *CarrinhoBusiness) CarrinhoByClienteID(clienteID string) (*model.Carrinho, error) {
	return l.carrinhoRepository.CarrinhoByClienteID(clienteID)
}
