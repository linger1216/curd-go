package svc

import (
	"compress/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/linger1216/go-front/geom"
	"github.com/linger1216/go-front/utils"
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
	req := &CreateEchoRequest{}
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
	req := &DeleteEchoRequest{}
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
	return req, err
}

type UpdateEchoRequest struct {
	Echos []*Echo
}

type UpdateEchoResponse struct {
}

func DecodeHTTPUpdateEchoRequest(c *gin.Context) (interface{}, error) {
	req := &UpdateEchoRequest{}

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
	Header      bool        `form:"header" json:"header"   yaml:"header" `
	Ages        []int       `form:"age" json:"ages" yaml:"ages" `
	Names       []string    `form:"name"  json:"names" yaml:"names" `
	Books       []string    `form:"book" json:"books" yaml:"books" `
	Tags        []int       `form:"tag" json:"tags,omitempty" yaml:"tags" `
	Point       *geom.Point `form:"point" json:"point,omitempty" yaml:"point" `
	Radius      float64     `form:"radius" json:"radius,omitempty" yaml:"radius" `
	StartTime   int64       `form:"start_time" json:"start_time" yaml:"start_time" `
	EndTime     int64       `form:"end_time" json:"end_time" yaml:"end_time" `
	CurrentPage uint64      `form:"current_page" json:"current_page" yaml:"current_page" `
	PageSize    uint64      `form:"page_size" json:"page_size" yaml:"page_size" `
}

type ListEchoResponse struct {
	Headers []*KVResponse `form:"headers" json:"headers,omitempty" yaml:"headers"`
	Echos   []*Echo       `form:"echos" json:"echos,omitempty" yaml:"echos"`
}

func DecodeHTTPListEchoRequest(c *gin.Context) (interface{}, error) {
	req := &ListEchoRequest{}
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

	err = decodeListEchoRequest(c, req)
	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeListEchoRequest(c *gin.Context, req *ListEchoRequest) error {
	err := c.BindQuery(&req)
	if err != nil {
		return err
	}
	lng := c.Query("longitude")
	lat := c.Query("latitude")
	req.Point = &geom.Point{}
	if len(lng) > 0 && len(lat) > 0 {
		req.Point.Coordinates = []float64{utils.StringToFloat(lng), utils.StringToFloat(lat)}
	}
	req.Point.SpatialReference = c.Query("spatial_reference")
	req.Radius = utils.StringToFloat(c.Query("radius"))
	if len(req.Point.Coordinates) == 0 {
		req.Point = nil
	}
	return nil
}

func DecodeHTTPListHeadEchoRequest(c *gin.Context) (interface{}, error) {
	req := &ListEchoRequest{}
	req.Header = true
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

	err = decodeListEchoRequest(c, req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

type GetEchoRequest struct {
	Ids []string `form:"ids" json:"ids,omitempty" yaml:"ids"`
}

type GetEchoResponse struct {
	Echos []*Echo `form:"echos" json:"echos,omitempty" yaml:"echos"`
}

func DecodeHTTPGetEchoRequest(c *gin.Context) (interface{}, error) {
	req := &GetEchoRequest{}
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
	return req, err
}
