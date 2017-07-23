package routes

import (
	"autenticacao-jwt/domain"
	"autenticacao-jwt/utils"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	response := domain.Response{"Access protected area!!!"}

	utils.JsonResponse(response, w)

}
