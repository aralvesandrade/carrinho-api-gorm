package model

import "time"

//Carrinho ...
type Carrinho struct {
	Base
	ClienteID           string    `gorm:"column:ClienteId;"`
	CEP                 string    `gorm:"column:CEP;"`
	FormaEnvio          string    `gorm:"column:FormaEnvio;"`
	FormaPagamento      string    `gorm:"column:FormaPagamento;"`
	ValorPedido         float64   `gorm:"column:ValorPedido;"`
	ValorFrete          float64   `gorm:"column:ValorFrete;"`
	Cupom               string    `gorm:"column:Cupom;"`
	Desconto            int       `gorm:"column:Desconto;"`
	ValorDesconto       float64   `gorm:"column:ValorDesconto;"`
	BandeiraCartao      string    `gorm:"column:BandeiraCartao;"`
	Parcelas            int       `gorm:"column:Parcelas;"`
	ValorParcela        float64   `gorm:"column:ValorParcela;"`
	ValorTotalParcelado float64   `gorm:"column:ValorTotalParcelado;"`
	ValorTotal          float64   `gorm:"column:ValorTotal;"`
	PrazoEntrega        int       `gorm:"column:PrazoEntrega;"`
	DataCarrinho        time.Time `gorm:"column:DataCarrinho;"`
}

//TableName ...
func (Carrinho) TableName() string {
	return "carrinho"
}
