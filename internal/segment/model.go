package segment

import "time"

type Segment struct {
	ID            time.Time `json:"timestamp" example:"2024-03-09T12:04:08Z"`
	TotalSegments uint      `json:"total_segments" example:"5"`
	SenderName    string    `json:"sender_name" example:"Некто"`
	SegmentNumber uint      `json:"segment_number" example:"1"`
	HadError      bool      `json:"had_error" example:"false"`
	Payload       []byte    `json:"payload"`
}

type SegmentRequest struct {
	ID            time.Time `json:"timestamp" example:"2024-03-09T12:04:08Z"`
	TotalSegments uint      `json:"total_segments" example:"5"`
	SenderName    string    `json:"sender" example:"Некто"`
	SegmentNumber uint      `json:"segment_number" example:"1"`
	Payload       string    `json:"message" example:"что-то"`
}
