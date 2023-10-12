package users

type drama int
type Karma drama

type User struct {
	Name     string `goofy:"required,gte=3,lte=75"`
	IsActive bool   `goofy:"required"`
	Karma    Karma  `goofy:"gte=0"`
	Address  Address
}

type Address struct {
	City    string `goofy:"gte=3,lte=75"`
	Country string `goofy:"required,iso-3166-alpha2"`
}
