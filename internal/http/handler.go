package http

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
	
	"github.com/sirupsen/logrus"

	"github.com/ellexo2456/NetworkingDataLinkLayer/internal/segment"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	BaseURL string
	log     *logrus.Logger
}

func NewHandler(baseURL string, log *logrus.Logger) *Handler {
	return &Handler{BaseURL: baseURL,
		log: log}
}

// EncodeSegmentSimulate
// @Summary EncodeSegmentSimulate.
// @Description Кодирует и декодирует полученный в виде байт сегмент, вносит ошибку, исправляет ее, так же с вероятностью возвращает сегмент на траспортный уровень.
// @Tags Code
// @Accept json
// @Produce json
// @Param segment body segment.SegmentRequest true "Пользовательский объект в формате JSON"
// @Success 200 "Успешный ответ"
// @Failure 400 {object} swag.ErrorResponse "Ошибка в запросе"
// @Failure 500 {object} swag.ErrorResponse "Внутренняя ошибка сервера"
// @Router /code [post]
func (h *Handler) EncodeSegmentSimulate(c *gin.Context) {
	const op = "handlers.EncodeSegmentSimulate"
	log := h.log.WithField("operation", op)

	var segReq segment.SegmentRequest
	if err := c.BindJSON(&segReq); err != nil {
		log.WithError(err).Error("ошибка чтения входящего JSON сегмента")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось прочитать JSON: " + err.Error()})
		return
	}

	log.WithField("segmentRequest", segReq).Info("на вход поступил сегмент")

	seg := segment.Segment{ID: segReq.ID,
		TotalSegments: segReq.TotalSegments,
		SenderName:    segReq.SenderName,
		SegmentNumber: segReq.SegmentNumber,
		Payload:       []byte(segReq.Payload),
	}

	randomNumber := rand.Intn(100)

	if randomNumber < 2 {
		log.WithField("segment", seg).Warn("потеря сегмента с вероятностью 2%")
		c.JSON(http.StatusBadRequest, gin.H{"error": "сегмент утерян"})
	}

	cycleCode := seg.Simulate(seg.SplitSegmentToCycleCodes(h.log), h.log)
	seg.Payload = nil
	seg.JoinCycleCodesToSegment(cycleCode, h.log)

	segResp := struct {
		ID            time.Time `json:"timestamp" example:"2024-03-09T12:04:08Z"`
		TotalSegments uint      `json:"total_segments" example:"5"`
		SenderName    string    `json:"sender" example:"Некто"`
		SegmentNumber uint      `json:"segment_number" example:"1"`
		HadError      bool      `json:"had_error" example:"false"`
		Payload       string    `json:"message"`
	}{
		ID: seg.ID,
		TotalSegments: seg.TotalSegments,
		SenderName: seg.SenderName,
		SegmentNumber: seg.SegmentNumber,
		HadError: seg.HadError,
		Payload: string(seg.Payload),
	}
	
	segmentJSON, err := json.Marshal(segResp)
	if err != nil {
		log.WithError(err).Error("ошибка при кодировании сегмента в JSON")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка при кодировании сегмента в JSON: " + err.Error()})
		return
	}
	
	log.Info("Transfer method", bytes.NewBuffer(segmentJSON))
	
	response, err := http.Post(h.BaseURL, "application/json", bytes.NewBuffer(segmentJSON))
	if err != nil {
		log.WithError(err).Error("ошибка при отправке сегмента на эндпоинт")
		c.JSON(http.StatusBadRequest, gin.H{"error": "ошибка при отправке сегмента на эндпоинт: " + err.Error()})
		return
	}
	defer response.Body.Close()
	
	if response.StatusCode != http.StatusOK {
		log.Error("неверный код состояния ответа")
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный код состояния ответа"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "сегмент успешно отправлен на эндпоинт"})
}
