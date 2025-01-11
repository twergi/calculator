package httpctrl

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/gorilla/schema"
	"github.com/twergi/calculator/internal/model"
)

type CalculatePostForm struct {
	A  int64  `schema:"a"`
	B  int64  `schema:"b"`
	Op string `schema:"op"`
}

func (c *HTTPController) calculate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	context := map[string]any{}
	defer func() {
		c.writePage(ctx, w, context)
	}()

	if err := r.ParseForm(); err != nil {
		context["error"] = err.Error()
		return
	}

	prev, err := c.calcUsecase.GetLastResult(ctx)
	if err != nil {
		if !errors.Is(err, model.ErrNotFound) {
			slog.ErrorContext(ctx, "calculate: get last result", "error", err.Error())
		}
	} else {
		context["prev"] = prev
	}

	decoder := schema.NewDecoder()

	formData := CalculatePostForm{}
	err = decoder.Decode(&formData, r.PostForm)
	if err != nil {
		context["error"] = err.Error()
		return
	}

	context["a"] = formData.A
	context["b"] = formData.B
	context["op"] = formData.Op

	result, err := c.calcUsecase.Calculate(
		ctx,
		formData.A,
		formData.B,
		model.OperationFromString(formData.Op),
	)
	if err != nil {
		context["error"] = err.Error()
		return
	}

	context["result"] = result
}
