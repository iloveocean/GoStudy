package tea_controller

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"go2/GoStudy/restful/common"
	"go2/GoStudy/restful/models/tea"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type AppContext struct {
	db *mgo.Database
}

func (ctx *AppContext) TeasHandler(w http.ResponseWriter, r *http.Request) {
	repo := tea.TeaRepo{ctx.db.C("teas")}
	teas, err := repo.All()
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", common.JASON_API_IDENTIFIER)
	json.NewEncoder(w).Encode(teas)
}

func (ctx *AppContext) TeaHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	repo := tea.TeaRepo{ctx.db.C("teas")}
	tea, err := repo.Find(params.ByName("id"))
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	json.NewEncoder(w).Encode(tea)
}

func (ctx *AppContext) CreateTeaHandler(w http.ResponseWriter, r *http.Request) {
	body := context.Get(r, "body").(*tea.TeaResource)
	repo := tea.TeaRepo{ctx.db.C("teas")}
	err := repo.Create(&body.Data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", common.JASON_API_IDENTIFIER)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(body)
}

func (ctx *AppContext) UpdateTeaHandler(w http.ResponseWriter, r *http.Request) {
	repo := tea.TeaRepo{ctx.db.C("teas")}
	params := context.Get(r, "params").(httprouter.Params)
	body := context.Get(r, "body").(*tea.TeaResource)
	body.Data.Id = bson.ObjectIdHex(params.ByName("id"))
	err := repo.Update(&body.Data)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(204)
	w.Write([]byte{'\n'})
}

func (ctx *AppContext) DeleteTeaHandler(w http.ResponseWriter, r *http.Request) {
	params := context.Get(r, "params").(httprouter.Params)
	repo := tea.TeaRepo{ctx.db.C("teas")}
	err := repo.Delete(params.ByName("id"))
	if err != nil {
		panic(err)
	}
	w.WriteHeader(204)
	w.Write([]byte{'\n'})
}

func (ctx *AppContext) AcceptHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		log.Printf("enter AcceptHandler\n")
		if r.Header.Get("Accept") != common.JASON_API_IDENTIFIER {
			common.WriteError(w, common.ErrNotAcceptable)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (ctx *AppContext) ContentTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != common.JASON_API_IDENTIFIER {
			common.WriteError(w, common.ErrUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func (ctx *AppContext) BodyHandler(v interface{}) func(http.Handler) http.Handler {
	t := reflect.TypeOf(v)

	m := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			val := reflect.New(t).Interface()
			err := json.NewDecoder(r.Body).Decode(val)
			if err != nil {
				common.WriteError(w, common.ErrBadRequest)
				return
			}
			if next != nil {
				context.Set(r, "body", val)
				next.ServeHTTP(w, r)
			}
		}
		return http.HandlerFunc(fn)
	}
	return m
}

func NewAppContext(db *mgo.Database) *AppContext {
	return &AppContext{db}
}
