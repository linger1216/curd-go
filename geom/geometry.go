package geom

type LngLat struct {
	/// the longitude of the `LngLat`
	Longitude float64 `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	/// the latitude of the `LngLat`
	Latitude float64 `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	/// the altitude of the `LngLat`
	Altitude float32 `protobuf:"fixed32,3,opt,name=altitude,proto3" json:"altitude,omitempty"`
}

type Point struct {
	Coordinates      *LngLat `protobuf:"bytes,1,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	Type             string  `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	SpatialReference string  `protobuf:"bytes,4,opt,name=spatial_reference,json=spatialReference,proto3" json:"spatialReference,omitempty"`
}

type Geometry struct {
}
