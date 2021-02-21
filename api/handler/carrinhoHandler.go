package handler

import (
	"carrinho-api-gorm/controller"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//carrinhoHandler ...
type carrinhoHandler struct {
	carrinhoController controller.ICarrinhoController
}

//NewCarrinhoHandler ...
func NewCarrinhoHandler(e *mux.Router, carrinhoController controller.ICarrinhoController) {
	handler := &carrinhoHandler{carrinhoController}

	s := e.PathPrefix("/v1").Subrouter()

	s.HandleFunc("/carrinho", handler.CarrinhoByClienteID).Methods(http.MethodGet)
}

//CarrinhoByClienteID ...
func (l *carrinhoHandler) CarrinhoByClienteID(w http.ResponseWriter, r *http.Request) {

	clienteID, presentClienteID := r.URL.Query()["clienteId"]

	if presentClienteID {
		result, error := l.carrinhoController.CarrinhoByClienteID(clienteID[0])

		if error != nil {
			fmt.Println(error)
			return
		}

		w.Header().Add(contentTypeConst, applicationJSON)

		response, _ := json.Marshal(result)

		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return

	} else {
		w.Header().Add(contentTypeConst, applicationJSON)
		response, _ := json.Marshal(map[string]interface{}{"mensagem": "Verifique os parametros obrigatórios enviados"})
		w.WriteHeader(http.StatusOK)
		w.Write(response)
		return
	}
}
