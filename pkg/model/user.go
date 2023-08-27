package model

type User struct {
	Id int `json:"id"`
}

type UserSegments struct {
	Id         int `json:"id"`
	SegmentsId int `json:"segments_id"`
}
