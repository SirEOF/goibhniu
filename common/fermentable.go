package common

import (
	"errors"
	"time"
)

type Fermentables struct {
	data   map[int]Fermentable
	nextid int
}

func NewFermentables() Fermentables {
	return Fermentables{
		data:   make(map[int]Fermentable),
		nextid: 0,
	}
}

func (this Fermentables) GetAll() map[int]Fermentable {
	return this.data
}

func (this Fermentables) Get(id int) (*Fermentable, error) {
	ferm, ok := this.data[id]
	if !ok {
		return nil, errors.New("Unable to get fermentable, invalid ID")
	}

	return &ferm, nil
}

func (this *Fermentables) Update(newdata Fermentable) error {
	if !this.Exist(newdata.ID) {
		return errors.New("ID does not exist, unable to update the data.")
	}

	this.data[newdata.ID] = newdata
	return nil
}

func (this *Fermentables) Delete(id int) {
	delete(this.data, id)
}

func (this Fermentables) Exist(id int) bool {
	_, ok := this.data[id]
	return ok
}

func (this *Fermentables) Add(newdata Fermentable) error {
	if this.Exist(newdata.ID) {
		return errors.New("ID already exists, unable to create vessel.")
	}

	newdata.ID = this.nextid
	this.data[this.nextid] = newdata
	this.nextid++

	return nil
}

type Fermentable struct {
	ID        int       `json:"id"`
	Recipe    string    `json:"recipe"`
	Tags      []string  `json:"tags"`
	Fermenter string    `json:"fermenter"`
	StartDate time.Time `json:"startdate"`
	EndDate   time.Time `json:"enddate"`
	Readings  []Reading `json:"readings"`
	Notes     string    `json:"notes"`
}

func NewFermentable(id int, name string, recipe string, vessel string, start time.Time) Fermentable {
	return Fermentable{
		ID:        id,
		Name:      name,
		Recipe:    recipe,
		Tags:      make([]string, 0),
		Fermenter: vessel,
		StartDate: start,
		Readings:  make([]Reading, 0),
	}
}

func (this *Fermentable) End(endtime time.Time) {
	this.EndDate = endtime
}

func (this *Fermentable) AddReading(update Reading) {
	this.Readings = append(this.Readings, update)
}

func (this *Fermentable) TakeReading(gravity float64) {
	this.Readings = append(this.Readings, MakeReading(gravity))
}

func (this Fermentable) GetReadings() []Reading {
	return this.Readings
}
