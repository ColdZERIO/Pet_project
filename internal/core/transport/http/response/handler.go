package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/ColdZERIO/Pet_project/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	w   http.ResponseWriter
}

func NewHTTPResponseHandler(log *core_logger.Logger, w http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		w:   w,
	}
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.w.WriteHeader(http.StatusInternalServerError)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.w).Encode(response); err != nil {
		h.log.Error("write HTTP response error", zap.Error(err))
	}
}
