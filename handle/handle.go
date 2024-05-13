package handle

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/charmbracelet/log"
)

type HandlerContext struct {
	w       http.ResponseWriter
	r       *http.Request
	Context map[string]any
}

func (c *HandlerContext) String(status int, msg string) error {
	c.w.WriteHeader(status)
	fmt.Fprint(c.w, msg)
	return nil
}

func (c *HandlerContext) Stringf(status int, msg string, args ...any) error {
	c.w.WriteHeader(status)
	fmt.Fprintf(c.w, msg, args...)
	return nil
}

func (c *HandlerContext) NoContent(status int) error {
	c.w.WriteHeader(status)
	return nil
}

func (c *HandlerContext) LogErr(status int, err error) error {
	c.w.WriteHeader(status)
	log.Error(err)
	return nil
}

func (c *HandlerContext) JSON(status int, v interface{}) error {
	err := json.NewEncoder(c.w).Encode(v)
	c.w.WriteHeader(status)
	return err
}

func (c *HandlerContext) BodyJson(v any) error {
	return json.NewDecoder(c.r.Body).Decode(v)
}

func (c *HandlerContext) Param(name string) string {
	return c.r.PathValue(name)
}

type HandlerFunc func(c *HandlerContext) error

func HTTPHandlerFunc(h HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := HandlerContext{w: w, r: r}
		err := h(&ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func Attach(mux *http.ServeMux, path string, h HandlerFunc) {
	mux.Handle(path, HTTPHandlerFunc(h))
}
