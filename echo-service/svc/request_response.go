package svc

import (
	"compress/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type KVResponse struct {
	Key   string
	Value string
}

type CreateEchoRequest struct {
	Echos []*Echo `form:"echos" json:"echos,omitempty" yaml:"echos"`
}

type CreateEchoResponse struct {
	Ids []string `form:"ids" json:"ids,omitempty" yaml:"ids"`
}

func DecodeHTTPCreateEchoRequest(c *gin.Context) (interface{}, error) {
	var req CreateEchoRequest
	var reader io.ReadCloser
	var err error
	switch c.GetHeader("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(c.Request.Body)
		defer reader.Close()
		if err != nil {
			return nil, NewError(http.StatusBadRequest, "failed to read the gzip content")
		}
	default:
		reader = c.Request.Body
	}

	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = jsoniter.ConfigFastest.Unmarshal(buf, &req.Echos); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
	}
	return req, nil
}

type DeleteEchoRequest struct {
	Ids []string `form:"ids" json:"ids,omitempty" yaml:"ids"`
}

type DeleteEchoResponse struct {
}

func DecodeHTTPDeleteEchoRequestV2(c *gin.Context) (interface{}, error) {
	var req DeleteEchoRequest
	var reader io.ReadCloser
	var err error
	switch c.GetHeader("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(c.Request.Body)
		defer reader.Close()
		if err != nil {
			return nil, NewError(http.StatusBadRequest, "failed to read the gzip content")
		}
	default:
		reader = c.Request.Body
	}

	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = jsoniter.ConfigFastest.Unmarshal(buf, &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
	}

	ids := make([]string, 0)
	queryIds := c.QueryArray("id")
	ids = append(ids, queryIds...)

	pathIds := strings.Split(c.Param("ids"), ",")
	ids = append(ids, pathIds...)

	if len(ids) == 0 {
		return nil, NewError(http.StatusBadRequest,
			fmt.Sprintf("request ids '%s': cannot parse request params", ids))
	}
	req.Ids = ids
	return &req, err
}

type UpdateEchoRequest struct {
	Echos []*Echo
}

type UpdateEchoResponse struct {
}

func DecodeHTTPUpdateEchoRequest(c *gin.Context) (interface{}, error) {
	var req UpdateEchoRequest

	var reader io.ReadCloser
	var err error
	switch c.GetHeader("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(c.Request.Body)
		defer reader.Close()
		if err != nil {
			return nil, NewError(http.StatusBadRequest, "failed to read the gzip content")
		}
	default:
		reader = c.Request.Body
	}

	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = jsoniter.ConfigFastest.Unmarshal(buf, &req.Echos); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
	}

	return req, nil
}

type ListEchoRequest struct {
	Header          bool     `form:"header" json:"header"   yaml:"header" `
	BoundingAreaIds []string `form:"bounding_area_id" json:"bounding_area_ids" yaml:"bounding_area_ids" `
	AccessKeys      []string `form:"access_key"  json:"access_keys" yaml:"access_keys" `
	Names           []string `form:"name"  json:"names" yaml:"names" `
	GeofenceIds     []string `form:"geofence_id" json:"geofence_ids" yaml:"geofence_ids" `
	Floors          []string `form:"floor" json:"floors" yaml:"floors" `
	FloorsIds       []string `form:"floors_id" json:"floors_ids,omitempty" yaml:"floors_ids" `
	RoomIds         []string `form:"room_id" json:"room_ids" yaml:"room_ids" `
	StartTime       int64    `form:"start_time" json:"start_time" yaml:"start_time" `
	EndTime         int64    `form:"end_time" json:"end_time" yaml:"end_time" `
	CurrentPage     uint64   `form:"current_page" json:"current_page" yaml:"current_page" `
	PageSize        uint64   `form:"page_size" json:"page_size" yaml:"page_size" `
}

type ListEchoResponse struct {
	Headers []*KVResponse `form:"headers" json:"headers,omitempty" yaml:"headers"`
	Echos   []*Echo       `form:"echos" json:"echos,omitempty" yaml:"echos"`
}

func DecodeHTTPListEchoRequest(c *gin.Context) (interface{}, error) {
	var req ListEchoRequest
	var reader io.ReadCloser
	var err error
	switch c.GetHeader("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(c.Request.Body)
		defer reader.Close()
		if err != nil {
			return nil, NewError(http.StatusBadRequest, "failed to read the gzip content")
		}
	default:
		reader = c.Request.Body
	}

	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = jsoniter.ConfigFastest.Unmarshal(buf, &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
	}

	err = c.BindQuery(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

func DecodeHTTPListHeadEchoRequest(c *gin.Context) (interface{}, error) {
	var req ListEchoRequest
	var reader io.ReadCloser
	var err error
	switch c.GetHeader("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(c.Request.Body)
		defer reader.Close()
		if err != nil {
			return nil, NewError(http.StatusBadRequest, "failed to read the gzip content")
		}
	default:
		reader = c.Request.Body
	}

	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = jsoniter.ConfigFastest.Unmarshal(buf, &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
	}

	err = c.BindQuery(&req)
	if err != nil {
		return nil, err
	}
	req.Header = true
	return req, nil
}

type GetEchoRequest struct {
	Ids []string `form:"ids" json:"ids,omitempty" yaml:"ids"`
}

type GetEchoResponse struct {
	Echos []*Echo `form:"echos" json:"echos,omitempty" yaml:"echos"`
}

func DecodeHTTPGetEchoRequest(c *gin.Context) (interface{}, error) {
	var req GetEchoRequest
	var reader io.ReadCloser
	var err error
	switch c.GetHeader("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(c.Request.Body)
		defer reader.Close()
		if err != nil {
			return nil, NewError(http.StatusBadRequest, "failed to read the gzip content")
		}
	default:
		reader = c.Request.Body
	}

	buf, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, NewError(http.StatusBadRequest, "cannot read body of http request")
	}
	if len(buf) > 0 {
		if err = jsoniter.ConfigFastest.Unmarshal(buf, &req); err != nil {
			const size = 8196
			if len(buf) > size {
				buf = buf[:size]
			}
			return nil, NewError(http.StatusBadRequest,
				fmt.Sprintf("request body '%s': cannot parse non-json request body", buf))
		}
	}

	ids := make([]string, 0)
	queryIds := c.QueryArray("id")
	ids = append(ids, queryIds...)

	pathIds := strings.Split(c.Param("ids"), ",")
	ids = append(ids, pathIds...)

	if len(ids) == 0 {
		return nil, NewError(http.StatusBadRequest,
			fmt.Sprintf("request ids '%s': cannot parse request params", ids))
	}
	req.Ids = ids
	return &req, err
}
