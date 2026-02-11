package handlers

import (
	"net/http"

	"github.com/builders-lab/trailblazer-frontend/internal/utils"
)

func HandlerReadiness(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJson(w, 200, struct{}{})
}
