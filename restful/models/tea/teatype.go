package tea

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Tea struct {
	Id       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name"`
	Category string        `json:"category"`
}

type TeaResource struct {
	Data Tea `json:"data"`
}

type TeasCollection struct {
	Data []Tea `json:"data"`
}

type TeaRepo struct {
	Collection *mgo.Collection
}

func (r *TeaRepo) All() (TeasCollection, error) {
	result := TeasCollection{[]Tea{}}
	err := r.Collection.Find(nil).All(&result.Data)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *TeaRepo) Find(id string) (TeaResource, error) {
	result := TeaResource{}
	err := r.Collection.FindId(bson.ObjectIdHex(id)).One(&result.Data)
	if err != nil {
		return result, err
	}
	return result, nil
}

func (r *TeaRepo) Create(tea *Tea) error {
	id := bson.NewObjectId()
	_, err := r.Collection.UpsertId(id, tea)
	if err != nil {
		return err
	}
	tea.Id = id
	return nil
}

func (r *TeaRepo) Update(tea *Tea) error {
	err := r.Collection.UpdateId(tea.Id, tea)
	if err != nil {
		return err
	}
	return nil
}

func (r *TeaRepo) Delete(id string) error {
	err := r.Collection.RemoveId(bson.ObjectIdHex(id))
	if err != nil {
		return err
	}
	return nil
}
