package coupon

import (
	"errors"
	"strings"

	"github.com/asdine/storm"
	"github.com/asdine/storm/q"
)

//Coupon entitity
type Coupon struct {
	Name      string `json:"name" storm:"id,unique"`
	Brand     string `json:"brand"`
	Value     int    `json:"value"`
	CreatedAt string `json:"createdAt"`
	Expiry    string `json:"expiry"`
}

const (
	dbPath = "coupons.db"
)

// errors
var (
	ErrRecordInvalid = errors.New("record is invalid")
)

// All retrieves all vouchers from the database
func All() ([]Coupon, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	coupons := []Coupon{}
	err = db.All(&coupons)
	if err != nil {
		return nil, err
	}
	return coupons, nil
}

//Filter return
func Filter(filters map[string][]string) ([]Coupon, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	coupons := []Coupon{}
	fmatchers := []q.Matcher{}
	for k, v := range filters {
		fmatchers = append(fmatchers, q.Eq(strings.Title(k), v[0]))
	}
	err = db.Select(fmatchers...).Find(&coupons)
	if err != nil {
		return nil, err
	}

	return coupons, nil
}

// One returns a single coupon record from the database
func One(couponName string) (*Coupon, error) {
	db, err := storm.Open(dbPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	c := new(Coupon)
	err = db.One("Name", couponName, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// Save updates or creates a given coupon in the database
func (c *Coupon) Save() error {
	if err := c.validate(); err != nil {
		return err
	}
	db, err := storm.Open(dbPath)
	if err != nil {
		return err
	}
	defer db.Close()
	return db.Save(c)
}

// validate makes sure that the record contains valid data
func (c *Coupon) validate() error {
	if c.Name == "" {
		return ErrRecordInvalid
	}
	return nil
}
