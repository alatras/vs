package ruleSet

import (
	"net/http"
	"validation-service/app/listAncestorsRuleSet"
	"validation-service/app/listDescendantsRuleSet"
	appd "validation-service/appdynamics"
	"validation-service/enums/contextKey"
	"validation-service/http/errorResponse"
	"validation-service/ruleSet"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func (resp ListRuleSetResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, http.StatusOK)
	return nil
}

func NewListRuleSetResponse(ruleSets []ruleSet.RuleSet) []render.Renderer {
	list := make([]render.Renderer, len(ruleSets))

	for i, r := range ruleSets {
		list[i] = ListRuleSetResponse{r}
	}

	return list
}

func (rs Resource) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	app := rs.listRuleSetAppFactory()

	entityId := chi.URLParam(r, "id")

	ruleSets, err := app.Execute(ctx, entityId)

	if err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)
		_ = render.Render(w, r, errorResponse.UnexpectedError(err.Error()))
		return
	}

	if err := render.RenderList(w, r, NewListRuleSetResponse(ruleSets)); err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)
		rs.logger.Error.WithError(err).Error("error rendering response")
		return
	}
}

func (rs Resource) ListAncestors(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	app := rs.listAncestorsRuleSetAppFactory()

	entityId := chi.URLParam(r, "id")

	ruleSets, err := app.Execute(ctx, []string{entityId}) // TODO: replace with entity ids list from the request

	if err.HasError() {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)

		if err.Is(listAncestorsRuleSet.EntityIdNotFoundErr) {
			_ = render.Render(w, r, errorResponse.ResourceNotFound("entity", entityId))
		} else if err.Is(listAncestorsRuleSet.EntityIdFormatIncorrectErr) {
			_ = render.Render(w, r, errorResponse.MalformedParameters(map[string]string{
				"params.entityId": err.Error(),
			}))
		} else {
			_ = render.Render(w, r, errorResponse.UnexpectedError(err.Error()))
		}
		return
	}

	if err := render.RenderList(w, r, NewListRuleSetResponse(ruleSets)); err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)
		rs.logger.Error.WithError(err).Error("error rendering response")
		return
	}
}

func (rs Resource) ListDescendants(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var businessTransaction appd.BtHandle

	if businessTransactionUid, ok := ctx.Value(contextKey.BusinessTransaction).(string); ok {
		businessTransaction = appd.GetBT(businessTransactionUid)
	}

	app := rs.listDescendantsRuleSetAppFactory()

	entityId := chi.URLParam(r, "id")

	ruleSets, err := app.Execute(ctx, []string{entityId}) // TODO: replace with entity ids from the request

	if err.HasError() {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)

		if err.Is(listDescendantsRuleSet.EntityIdNotFoundErr) {
			_ = render.Render(w, r, errorResponse.ResourceNotFound("entity", entityId))
		} else if err.Is(listDescendantsRuleSet.EntityIdFormatIncorrectErr) {
			_ = render.Render(w, r, errorResponse.MalformedParameters(map[string]string{
				"params.entityId": err.Error(),
			}))
		} else {
			_ = render.Render(w, r, errorResponse.UnexpectedError(err.Error()))
		}
		return
	}

	if err := render.RenderList(w, r, NewListRuleSetResponse(ruleSets)); err != nil {
		appd.AddBTError(businessTransaction, appd.APPD_LEVEL_ERROR, err.Error(), true)
		rs.logger.Error.WithError(err).Error("error rendering response")
		return
	}
}
