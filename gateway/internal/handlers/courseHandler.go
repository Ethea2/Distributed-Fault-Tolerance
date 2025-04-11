package handlers

import (
	"net/http"

	"github.com/Ethea2/Distributed-Fault-Tolerance/gateway/internal/proxy"
)

type CourseHandler struct {
	proxy *proxy.ServiceProxy
}

func NewCourseHandler(proxy *proxy.ServiceProxy) *CourseHandler {
	return &CourseHandler{
		proxy: proxy,
	}
}

func (h *CourseHandler) GetCourses(w http.ResponseWriter, r *http.Request) {
	if err := h.proxy.ForwardRequestAndCopyResponse(w, r, "/courses"); err != nil {
		http.Error(w, "Error forwarding request: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *CourseHandler) GetAvailableCourses(w http.ResponseWriter, r *http.Request) {
	if err := h.proxy.ForwardRequestAndCopyResponse(w, r, "/available_courses"); err != nil {
		http.Error(w, "Error forwarding request: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	if err := h.proxy.ForwardRequestAndCopyResponse(w, r, "/create_course"); err != nil {
		http.Error(w, "Error forwarding request: "+err.Error(), http.StatusInternalServerError)
	}
}

func (h *CourseHandler) EnrollCourse(w http.ResponseWriter, r *http.Request) {
	if err := h.proxy.ForwardRequestAndCopyResponse(w, r, "/enroll_course"); err != nil {
		http.Error(w, "Error forwarding request: "+err.Error(), http.StatusInternalServerError)
	}
}
