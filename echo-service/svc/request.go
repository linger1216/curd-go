package svc

import (
	"context"
	"net/http"
)

type CreateEchoRequest struct {
}

type CreateEchoResponse struct {
}

func DecodeHTTPCreateEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

type DeleteEchoRequest struct {
}

type DeleteEchoResponse struct {
}

func DecodeHTTPDeleteEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

type UpdateEchoRequest struct {
}

type UpdateEchoResponse struct {
}

func DecodeHTTPUpdateEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

type ListEchoRequest struct {
	Header          bool     `form:"header" json:"header"   yaml:"header" `
	BoundingAreaIds []string `form:"bounding_area_ids" json:"bounding_area_ids" yaml:"bounding_area_ids" `
	AccessKeys      []string `form:"access_keys"  json:"access_keys" yaml:"access_keys" `
	Names           []string `form:"names"  json:"names" yaml:"names" `
	GeofenceIds     []string `form:"geofence_ids" json:"geofence_ids" yaml:"geofence_ids" `
	Floors          []string `form:"floors" json:"floors" yaml:"floors" `
	FloorsIds       []string `form:"floors_ids" json:"floors_ids,omitempty" yaml:"floors_ids" `
	RoomIds         []string `form:"room_ids" json:"room_ids" yaml:"room_ids" `
	StartTime       int64    `form:"start_time" json:"start_time" yaml:"start_time" `
	EndTime         int64    `form:"end_time" json:"end_time" yaml:"end_time" `
	CurrentPage     uint64   `form:"current_page" json:"current_page" yaml:"current_page" `
	PageSize        uint64   `form:"page_size" json:"page_size" yaml:"page_size" `
}

type ListEchoResponse struct {
}

func DecodeHTTPListEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func DecodeHTTPListHeadEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

type GetEchoRequest struct {
}

type GetEchoResponse struct {
}

func DecodeHTTPGetEchoRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
