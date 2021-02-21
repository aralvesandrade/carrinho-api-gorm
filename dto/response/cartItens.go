package response

//ItensCarrinho ...
type ItensCarrinho struct {
	ProdutoID     string  `json:"produtoId"`
	Nome          string  `json:"nome"`
	Peso          float64 `json:"peso"`
	ValorUnitario float64 `json:"valorUnitario"`
	Qtd           int     `json:"qtd"`
	SubTotal      float64 `json:"subTotal"`
	Foto          string  `json:"foto"`
	SKU           string  `json:"sku"`
}
