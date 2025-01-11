package httpctrl

import (
	"context"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/twergi/calculator/internal/app/usecases/calculator"
)

var (
	templ = template.Must(template.ParseFiles("internal/app/components/httpctrl/page.html"))
)

type HTTPController struct {
	calcUsecase *calculator.Usecase

	server *http.Server
}

func New(addr string, calcUsecase *calculator.Usecase) *HTTPController {
	c := &HTTPController{
		calcUsecase: calcUsecase,
	}

	c.initServer(addr)

	return c
}

func (c *HTTPController) initServer(addr string) {
	mux := http.NewServeMux()
	mux.Handle("/calculate", http.HandlerFunc(c.calculate))
	mux.Handle("/", http.HandlerFunc(c.index))

	server := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	c.server = server
}

func (c *HTTPController) writePage(ctx context.Context, w http.ResponseWriter, context map[string]any) {
	err := templ.ExecuteTemplate(w, "page.html", context)
	if err != nil {
		slog.ErrorContext(ctx, "write page", "error", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *HTTPController) Run(ctx context.Context) error {
	slog.InfoContext(ctx, "starting http server", "addr", c.server.Addr)
	return c.server.ListenAndServe()
}

func (c *HTTPController) Stop(ctx context.Context) error {
	slog.InfoContext(ctx, "stopping http server")
	return c.server.Shutdown(ctx)
}
