package handlers

import (
	"net/http"

	"github.com/builders-lab/trailblazer-frontend/internal/utils"
)

func HandlerErr(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, 400, "Something Went Wrong")
}
