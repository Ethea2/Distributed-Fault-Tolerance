package handlers

import (
	"net/http"

	"github.com/Ethea2/Distributed-Fault-Tolerance/gateway/internal/proxy"
)

type GradeHandler struct {
	proxy *proxy.ServiceProxy
}

func NewGradeHandler(proxy *proxy.ServiceProxy) *GradeHandler {
	return &GradeHandler{
		proxy: proxy,
	}
}

func (h *GradeHandler) GetGrades(w http.ResponseWriter, r *http.Request) {
	if err := h.proxy.ForwardRequestAndCopyResponse(w, r, "/student_grades"); err != nil {
		http.Error(w, "Error forwarding request: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *GradeHandler) GetStudentGrades(w http.ResponseWriter, r *http.Request) {
	if err := h.proxy.ForwardRequestAndCopyResponse(w, r, "/faculty_grades"); err != nil {
		http.Error(w, "Error forwarding request: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *GradeHandler) GradeStudent(w http.ResponseWriter, r *http.Request) {
	if err := h.proxy.ForwardRequestAndCopyResponse(w, r, "/grade"); err != nil {
		http.Error(w, "Error forwarding request: "+err.Error(), http.StatusInternalServerError)
	}
}
