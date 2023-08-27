package model

type Segments struct {
	Id           int    `json:"id"`
	SegmentsName string `json:"segments_name" binding:"required"`
}
