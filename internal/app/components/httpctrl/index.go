package httpctrl

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/twergi/calculator/internal/model"
)

func (c *HTTPController) index(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	context := map[string]any{}
	defer func() {
		c.writePage(ctx, w, context)
	}()

	prev, err := c.calcUsecase.GetLastResult(ctx)
	if err != nil {
		if !errors.Is(err, model.ErrNotFound) {
			slog.ErrorContext(ctx, "calculate: get last result", "error", err.Error())
		}
	} else {
		context["prev"] = prev
	}
}
