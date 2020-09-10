package core

type Echo struct {
	Id             string   `json:"id" yaml:"id" `
	BoundingAreaId string   `json:"boundingAreaId" yaml:"boundingAreaId" `
	AccessKey      string   `json:"accessKey" yaml:"accessKey" `
	GeofenceId     string   `json:"geofenceId" yaml:"geofenceId" `
	FloorId        string   `json:"floorId" yaml:"floorId" `
	Floor          string   `json:"floor" yaml:"floor" `
	RoomId         string   `json:"roomId" yaml:"roomId" `
	CreateTime     int64    `json:"createTime" yaml:"createTime" `
	UpdateTime     int64    `json:"updateTime" yaml:"updateTime" `
	Macs           []string `json:"macs"`
}
