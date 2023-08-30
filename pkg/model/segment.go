package model

type Segments struct {
	Percent      int    `json:"percent"`
	SegmentsName string `json:"segments_name" binding:"required"`
}
