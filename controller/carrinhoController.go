package controller

import (
	"carrinho-api-gorm/business"
	"carrinho-api-gorm/dto/response"
)

//ICarrinhoController ...
type ICarrinhoController interface {
	CarrinhoByClienteID(clienteID string) (*response.Carrinho, error)
}

//CarrinhoController ...
type CarrinhoController struct {
	carrinhoBusiness      business.ICarrinhoBusiness
	itensCarrinhoBusiness business.IItensCarrinhoBusiness
}

//NewCarrinhoController ...
func NewCarrinhoController(carrinhoBusiness business.ICarrinhoBusiness, itensCarrinhoBusiness business.IItensCarrinhoBusiness) ICarrinhoController {
	return &CarrinhoController{carrinhoBusiness, itensCarrinhoBusiness}
}

//CarrinhoByClienteID ...
func (l *CarrinhoController) CarrinhoByClienteID(clienteID string) (*response.Carrinho, error) {
	carrinho, error := l.carrinhoBusiness.CarrinhoByClienteID(clienteID)

	if error != nil {
		return nil, error
	}

	itensCarrinho, error := l.itensCarrinhoBusiness.ItensCarrinhoByCarrinhoID(carrinho.ID)

	if error != nil {
		return nil, error
	}

	var responseCarrinho response.Carrinho
	responseItens := []response.ItensCarrinho{}

	responseCarrinho.OK = true
	responseCarrinho.ID = carrinho.ID

	var totalItens int
	var total, pesoTotal float64

	for _, item := range *itensCarrinho {
		totalItens += item.Qtd
		total += item.SubTotal
		pesoTotal += item.Peso

		itemCarrinho := response.ItensCarrinho{
			ProdutoID:     item.ProdutoID,
			Nome:          item.Nome,
			Peso:          item.Peso,
			ValorUnitario: item.ValorUnitario,
			Qtd:           item.Qtd,
			SubTotal:      item.SubTotal,
			Foto:          item.Foto,
			SKU:           item.SKU,
		}

		responseItens = append(responseItens, itemCarrinho)
	}

	responseCarrinho.TotalItens = totalItens
	responseCarrinho.Total = total
	responseCarrinho.PesoTotal = pesoTotal
	responseCarrinho.Itens = responseItens

	return &responseCarrinho, nil
}
