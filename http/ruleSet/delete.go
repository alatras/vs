package ruleSet

import (
	"net/http"
	"validation-service/app/deleteRuleSet"
	appd "validation-service/appdynamics"
	"validation-service/enums/contextKey"
	"validation-service/http/errorResponse"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (rs Resource) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	app := rs.deleteRuleSetAppFactory()

	entityId := chi.URLParam(r, "id")
	ruleSetId := chi.URLParam(r, "ruleSetId")

	err := app.Execute(ctx, entityId, ruleSetId)

	if err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)

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
