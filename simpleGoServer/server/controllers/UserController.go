package controllers

import (
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"github.com/julienschmidt/httprouter"
	"simpleGoServer/server/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"simpleGoServer/server/con"
)

func renderJson(w http.ResponseWriter, code int, res interface{}) {
	resBytes, _ := json.MarshalIndent(res, "", "	")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("code", fmt.Sprint(code))
	w.WriteHeader(code)
	w.Write(resBytes)
}

func SaveUser(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]interface{})
	user := models.User{}

	reqBytes, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(reqBytes, &user)
	if err != nil {
		res["message"] = "Error in parsing json"
		res["error"] = err.Error()
		renderJson(w, http.StatusBadRequest, res)
		return
	}

	user.DateCreated = time.Now().UTC()
	user.LastUpdated = time.Now().UTC()

	info, err := con.Db.C("user").Upsert(models.User{}, &user)
	if info != nil {
		user.ID = info.UpsertedId.(bson.ObjectId)
	}

	if err != nil {
		res["message"] = "Error in save user"
		res["error"] = err.Error()
		renderJson(w, http.StatusBadRequest, res)
		return
	}

	res["message"] = "User saved successfully"
	res["user"] = user
	renderJson(w, http.StatusOK, res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	res := make(map[string]interface{})
	user := models.User{}
	params := r.Context().Value("params").(httprouter.Params)
	id := params.ByName("id")

	if !bson.IsObjectIdHex(id) {
		res["message"] = "Invalid object ID"
		renderJson(w, http.StatusBadRequest, res)
		return
	}

	err := con.Db.C("user").FindId(bson.ObjectIdHex(id)).One(&user)
	if err != nil {
		res["message"] = "User not found"
		res["error"] = err.Error()
		renderJson(w, http.StatusNotFound, res)
		return
	}

	renderJson(w, http.StatusOK, user)
}

