package async

import (
	"context"
	"net/http"

	"github.com/ameprizzo/mpesago/api"
	"github.com/techcraftlabs/base"
)

type Handler struct {
	rv base.Receiver
	rp base.Replier
	fn func(request interface{}) (response interface{}, err error)
}

func NewHandler(rv base.Receiver, rp base.Replier,
	fn func(request interface{}) (response interface{}, err error)) *Handler {
	return &Handler{rv, rp, fn}
}

func (h *Handler) HandleCallback(category api.Category) http.HandlerFunc {
	var v interface{}

	return func(w http.ResponseWriter, r *http.Request) {
		_, err := h.rv.Receive(context.Background(), "", r, v)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		res, err := h.fn(v)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		h.rp.Reply(w, base.NewResponse(200, res))
	}
}
