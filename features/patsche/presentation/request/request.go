package request

import "finalproject/features/patsche"

type Patsche struct {
	AdminID int `json:"adminid"`
	Day  string `json:"day"`
	Time string `json:"time"`
}

func (req *Patsche) ToDomain() *patsche.Domain {
	return &patsche.Domain{
		AdminID: req.AdminID,
		Day:  req.Day,
		Time: req.Time,
	}
}
