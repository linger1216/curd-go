package core

import (
	"fmt"
	"github.com/linger1216/go-front/echo-service/svc"
	"testing"
)

func TestNewMetaTable(t *testing.T) {
	j := `{
  "Name": "echo_table",
  "columns": [
    {
      "Name": "id",
      "type": "string",
      "primary": true
    },
    {
      "Name": "access_key",
      "type": "string"
    },
    {
      "Name": "bounding_area_id",
      "type": "string"
    },
    {
      "Name": "geofence_id",
      "type": "string"
    },
    {
      "Name": "room_id",
      "type": "string"
    },
    {
      "Name": "floor",
      "type": "string"
    },
    {
      "Name": "floor_id",
      "type": "string"
    },
    {
      "Name": "create_time",
      "type": "int64"
    },
    {
      "Name": "update_time",
      "type": "int64"
    }
  ]
}`

	var query string
	var args []interface{}
	table := NewMetaTable([]byte(j))

	fmt.Printf("%s\n", table.CreateTable())

	query, args = table.Delete("id1", "id2", "id3")
	fmt.Printf("%s %v\n", query, args)

	query, args = table.Upsert(&Echo{
		BoundingAreaId: "boundingId1",
		AccessKey:      "ak1",
		GeofenceId:     "geofence1",
		FloorId:        "floorId1",
		Floor:          "7",
		RoomId:         "roomId1",
		CreateTime:     100000,
		UpdateTime:     200000,
		Macs:           []string{"mac1,mac2"},
	}, &Echo{
		Id:             "id2",
		BoundingAreaId: "boundingId2",
		AccessKey:      "ak2",
		GeofenceId:     "geofence2",
		FloorId:        "floorId2",
		Floor:          "8",
		RoomId:         "roomId2",
		CreateTime:     200000,
		UpdateTime:     300000,
		Macs:           []string{"mac3,mac4"},
	})
	fmt.Printf("%s %v\n", query, args)

	query, args = table.List(&svc.ListEchoRequest{
		Header:          false,
		BoundingAreaIds: []string{"boundingId1"},
		AccessKeys:      []string{"ak1"},
		Names:           []string{"name1"},
		GeofenceIds:     []string{"geofence1"},
		Floors:          []string{"9"},
		FloorsIds:       []string{"floorId1"},
		RoomIds:         []string{"roomId1"},
		StartTime:       100000,
		EndTime:         200000,
		CurrentPage:     0,
		PageSize:        100,
	})
	fmt.Printf("%s %v\n", query, args)

	query, args = table.List(&svc.ListEchoRequest{
		Header:          true,
		BoundingAreaIds: []string{"boundingId1"},
		AccessKeys:      []string{"ak1"},
		Names:           []string{"name1"},
		GeofenceIds:     []string{"geofence1"},
		Floors:          []string{"9"},
		FloorsIds:       []string{"floorId1"},
		RoomIds:         []string{"roomId1"},
		StartTime:       100000,
		EndTime:         200000,
		CurrentPage:     0,
		PageSize:        100,
	})
	fmt.Printf("%s %v\n", query, args)

	query, args = table.Get("id1", "id2")
	fmt.Printf("%s %v\n", query, args)
}
