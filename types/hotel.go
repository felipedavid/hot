package types

type HotelParamsErrors struct {
	NameError     string `json:"name,omitempty"`
	LocationError string `json:"location,omitempty"`
}

type UpdateHotelParams struct {
	Name     *string `bson:"name" json:"name"`
	Location *string `bson:"location" json:"location"`

	HotelParamsErrors
}

func (p *UpdateHotelParams) Validate() (valid bool) {
	if p.Name != nil && len(*p.Name) < 2 {
		p.NameError = "name too short"
	}

	if p.Location != nil && len(*p.Location) < 2 {
		p.LocationError = "location too short"
	}

	return p.NameError == "" && p.LocationError == ""
}

type CreateHotelParams struct {
	Name     string `bson:"name" json:"name"`
	Location string `bson:"location" json:"location"`

	HotelParamsErrors
}

func (p *CreateHotelParams) Validate() (valid bool) {
	if len(p.Name) < 2 {
		p.NameError = "name too short"
	}

	if len(p.Location) < 2 {
		p.LocationError = "location too short"
	}

	return p.NameError == "" && p.LocationError == ""
}

type Hotel struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Name     string `bson:"name" json:"name"`
	Location string `bson:"location" json:"location"`
}

func NewHotel(p *CreateHotelParams) *Hotel {
	return &Hotel{
		Name:     p.Name,
		Location: p.Location,
	}
}

type RoomType int

const (
	_ RoomType = iota
	SingleRT
	DoubleRT
	SeaSideRT
	DeluxRT
)

type Room struct {
	ID        string   `bson:"_id,omitempty" json:"id"`
	Type      RoomType `bson:"type" json:"type"`
	BasePrice float64  `bson:"base_price" json:"base_price"`
	Price     float64  `bson:"price" json:"price"`
	HotelID   string   `bson:"hotel_id" json:"hotel_id"`
}
