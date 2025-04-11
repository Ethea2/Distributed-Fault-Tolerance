package handlers

import (
	"net/http"

	"github.com/Ethea2/Distributed-Fault-Tolerance/gateway/internal/proxy"
)

type AuthHandler struct {
	proxy *proxy.ServiceProxy
}

func NewAuthHandler(proxy *proxy.ServiceProxy) *AuthHandler {
	return &AuthHandler{
		proxy: proxy,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	if err := h.proxy.ForwardRequestAndCopyResponse(w, r, "/login"); err != nil {
		http.Error(w, "Error forwarding request: "+err.Error(), http.StatusInternalServerError)
	}
}
