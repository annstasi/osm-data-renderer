package osm

import (
	"time"
)

type OSM struct {
	Name      string    `db:"name"`
	WayId     int64     `db:"way_id"`
	Polygon   string    `db:"polygon"`
	Lat       float64   `db:"lat"`
	Lon       float64   `db:"lon"`
	Tags      string    `db:"tags"`
	Type      string    `db:"type"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func (o OSM) Values() ([]interface{}, error) {
	return []interface{}{
		o.WayId,
		o.Name,
		o.Polygon,
		o.Lat,
		o.Lon,
		o.Tags,
		o.Type,
		o.CreatedAt,
		o.UpdatedAt,
	}, nil
}
