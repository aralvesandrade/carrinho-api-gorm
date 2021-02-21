package response

//Carrinho ...
type Carrinho struct {
	ResponseBase
	ID          string          `json:"id"`
	TotalItens  int             `json:"totalItens"`
	Total       float64         `json:"total"`
	PesoTotal   float64         `json:"pesoTotal"`
	RemoverItem bool            `json:"removerItem"`
	Itens       []ItensCarrinho `json:"itens"`
}

//ResponseBase ...
type ResponseBase struct {
	OK       bool   `json:"ok"`
	Mensagem string `json:"mensagem"`
}
