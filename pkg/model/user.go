package model

type User struct {
	Id int `json:"id"`
}

type UserSegments struct {
	Id           int      `json:"id"`
	SegmentsName []string `json:"segments_name"`
}
