package common

import "errors"

type Vessels []Vessel

func NewVessels() Vessels {
	return make(Vessels, 0)
}

func (this Vessels) Get(id string) (Vessel, error) {
	for _, vessel := range this {
		if vessel.ID == id {
			return vessel, nil
		}
	}
	return Vessel{}, errors.New("Unable to get vessel, invalid vessel ID")
}

func (this *Vessels) Update(newdata Vessel) error {
	for _, cur := range *this {
		if cur.ID == newdata.ID {
			cur = newdata
			return nil
		}
	}
	return errors.New("Unable to update vessel, invalid vessed ID")
}

func (this *Vessels) Delete(id string) error {
	for i, cur := range *this {
		if cur.ID == id {
			temp := *this
			temp = append(temp[:i], temp[i+1:]...)
			this = &temp
			return nil
		}
	}

	return errors.New("Invalid ID, vessel does not exist.")
}

func (this Vessels) Exist(id string) bool {
	for _, cur := range this {
		if cur.ID == id {
			return true
		}
	}
	return false
}

func (this Vessels) Add(newdata Vessel) error {
	if this.Exist(newdata.ID) {
		return errors.New("ID already exists, unable to create vessel.")
	}

	this = append(this, newdata)
	return nil
}

type Vessel struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Size int    `json:"size"`
}
