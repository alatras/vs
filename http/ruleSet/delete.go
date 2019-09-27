package ruleSet

import (
	"bitbucket.verifone.com/validation-service/app/deleteRuleSet"
	"bitbucket.verifone.com/validation-service/http/errorResponse"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func (rs Resource) Delete(w http.ResponseWriter, r *http.Request) {
	app := rs.deleteRuleSetAppFactory()

	ctx := r.Context()
	entityId := chi.URLParam(r, "id")
	ruleSetId := chi.URLParam(r, "ruleSetId")

	err := app.Execute(ctx, entityId, ruleSetId)

	if err != nil {
		switch err {
		case deleteRuleSet.NotFound:
			_ = render.Render(w, r, errorResponse.ResourceNotFound("ruleSet", ruleSetId))
		case deleteRuleSet.UnexpectedError:
			_ = render.Render(w, r, errorResponse.UnexpectedError(err.Error()))
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
