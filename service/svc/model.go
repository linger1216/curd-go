package svc

import "github.com/linger1216/go-front/geom"

type Echo struct {
	Id         string         `json:"id" yaml:"id" `
	Age        int            `json:"age"`
	Name       string         `json:"name"`
	Geometry   *geom.Geometry `json:"geometry"`
	Books      []string       `json:"books"`
	Tags       []int          `json:"tags"`
	CreateTime int64          `json:"create_time"`
	UpdateTime int64          `json:"update_time"`
}
