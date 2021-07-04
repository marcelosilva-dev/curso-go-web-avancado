package handler

import (
	"net/http"

	"github.com/marcelosilva-dev/curso-go-web-avancado/lib/contx"
)

//IndexCliente abre a pagina de gerenciamento de clientes
func Ver3D(ctx *contx.Context) {
	ctx.NativeHTML(http.StatusOK, "3D")
}
