package main

import (
	"net/http"

	"gopkg.in/mgo.v2"

	"go2/GoStudy/restful/common"
	"go2/GoStudy/restful/controllers/tea_controller"
	"go2/GoStudy/restful/models/tea"

	"github.com/gorilla/context"
	"github.com/justinas/alice"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	appCtx := tea_controller.NewAppContext(session.DB("test"))
	commonHandlers := alice.New(context.ClearHandler, common.LoggingHandler, common.RecoverHandler, appCtx.AcceptHandler)
	router := common.NewRouter()
	router.Get("/teas/:id", commonHandlers.ThenFunc(appCtx.TeaHandler))
	router.Put("/teas/:id", commonHandlers.Append(appCtx.ContentTypeHandler, appCtx.BodyHandler(tea.TeaResource{})).ThenFunc(appCtx.UpdateTeaHandler))
	router.Delete("teas/:id", commonHandlers.ThenFunc(appCtx.DeleteTeaHandler))
	router.Get("/teas", commonHandlers.ThenFunc(appCtx.TeasHandler))
	router.Post("teas", commonHandlers.Append(appCtx.ContentTypeHandler, appCtx.BodyHandler(tea.TeaResource{})).ThenFunc(appCtx.CreateTeaHandler))
	http.ListenAndServe(":3000", router)
}
