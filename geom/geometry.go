package geom

type Point struct {
	Coordinates      []float64 `protobuf:"bytes,1,opt,name=coordinates,proto3" json:"coordinates,omitempty"`
	Type             string    `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	SpatialReference string    `protobuf:"bytes,4,opt,name=spatial_reference,json=spatialReference,proto3" json:"spatialReference,omitempty"`
}

type Geometry struct {
	Point
}
