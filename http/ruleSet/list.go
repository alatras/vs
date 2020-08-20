package ruleSet

import (
	"bitbucket.verifone.com/validation-service/app/listAncestorsRuleSet"
	"bitbucket.verifone.com/validation-service/app/listDescendantsRuleSet"
	appd "bitbucket.verifone.com/validation-service/appdynamics"
	"bitbucket.verifone.com/validation-service/http/errorResponse"
	"bitbucket.verifone.com/validation-service/ruleSet"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
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
	appDCorrelationHeader := r.Header.Get(appd.APPD_CORRELATION_HEADER_NAME)
	businessTransaction := appd.StartBT("List rulesets", appDCorrelationHeader)
	appd.SetBTURL(businessTransaction, r.URL.Path)
	defer appd.EndBT(businessTransaction)

	app := rs.listRuleSetAppFactory()

	ctx := r.Context()

	entityId := chi.URLParam(r, "id")

	ruleSets, err := app.Execute(ctx, entityId)

	if err != nil {
		_ = render.Render(w, r, errorResponse.UnexpectedError(err.Error()))
		return
	}

	if err := render.RenderList(w, r, NewListRuleSetResponse(ruleSets)); err != nil {
		rs.logger.Error.WithError(err).Error("error rendering response")
		return
	}
}

func (rs Resource) ListAncestors(w http.ResponseWriter, r *http.Request) {
	app := rs.listAncestorsRuleSetAppFactory()

	ctx := r.Context()

	entityId := chi.URLParam(r, "id")

	ruleSets, err := app.Execute(ctx, []string{entityId}) // TODO: replace with entity ids list from the request

	if err.HasError() {
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
		rs.logger.Error.WithError(err).Error("error rendering response")
		return
	}
}

func (rs Resource) ListDescendants(w http.ResponseWriter, r *http.Request) {
	app := rs.listDescendantsRuleSetAppFactory()

	ctx := r.Context()

	entityId := chi.URLParam(r, "id")

	ruleSets, err := app.Execute(ctx, []string{entityId}) // TODO: replace with entity ids from the request

	if err.HasError() {
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
		rs.logger.Error.WithError(err).Error("error rendering response")
		return
	}
}
