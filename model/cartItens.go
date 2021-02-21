package model

//ItensCarrinho ...
type ItensCarrinho struct {
	Base
	CarrinhoID    string  `gorm:"column:CarrinhoId;"`
	ProdutoID     string  `gorm:"column:ProdutoId;"`
	Nome          string  `gorm:"column:Nome;"`
	ValorUnitario float64 `gorm:"column:ValorUnitario;"`
	Qtd           int     `gorm:"column:Qtd;"`
	Foto          string  `gorm:"column:Foto;"`
	Peso          float64 `gorm:"column:Peso;"`
	SubTotal      float64 `gorm:"column:SubTotal;"`
	SKU           string  `gorm:"column:Sku;"`
}
