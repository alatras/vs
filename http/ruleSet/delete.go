package ruleSet

import (
	"bitbucket.verifone.com/validation-service/app/deleteRuleSet"
	appd "bitbucket.verifone.com/validation-service/appdynamics"
	"bitbucket.verifone.com/validation-service/http/errorResponse"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
)

func (rs Resource) Delete(w http.ResponseWriter, r *http.Request) {
	appDCorrelationHeader := r.Header.Get(appd.APPD_CORRELATION_HEADER_NAME)
	businessTransaction := appd.StartBT("Delete ruleset", appDCorrelationHeader)
	appd.SetBTURL(businessTransaction, r.URL.Path)
	defer appd.EndBT(businessTransaction)

	app := rs.deleteRuleSetAppFactory()

	ctx := r.Context()
	entityId := chi.URLParam(r, "id")
	ruleSetId := chi.URLParam(r, "ruleSetId")

	err := app.Execute(ctx, entityId, ruleSetId)

	if err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), false)

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
