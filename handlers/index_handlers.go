package handlers

import (
	"encoding/json"
	"errors"
	"kafekoding-client/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Index(ctx *gin.Context) {
	session := sessions.Default(ctx)

	user := session.Get("user")
	if user == nil {
		flasher.Set("danger", "Please login first")
		ctx.Redirect(http.StatusFound, "/login?redirect_to=/")
		return
	}

	req, err := http.NewRequest("GET", CLASSES_URL, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot connecting to the server"))
		return
	}

	userData := user.(UserPayload)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userData.Token)

	client := http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for key, val := range via[0].Header {
			req.Header[key] = val
		}
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, errors.New("cannot connecting to the server"))
		return
	}

	classes := []models.Class{}
	_ = json.NewDecoder(resp.Body).Decode(&classes)
	ctx.HTML(http.StatusOK, "index", gin.H{
		"user":       userData,
		"classes":    classes,
		"server_url": SERVER_URL,
	})
}

// Detail Of Class
//
// # Detail of class
//
// This is for class detail
func Detail(ctx *gin.Context) {
	session := sessions.Default(ctx)

	id := ctx.Param("id")

	user := session.Get("user")
	if user == nil {
		flasher.Set("danger", "Please login first before accessing this page.")
		ctx.Redirect(http.StatusFound, "/login?redirect_to=/classes/"+id)
		return
	}

	userData := user.(UserPayload)
	req, err := http.NewRequest("GET", CLASSES_URL+"/"+id, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+userData.Token)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	client := &http.Client{}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		for k, v := range via[0].Header {
			req.Header[k] = v
		}
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	class := models.Class{}
	err = json.NewDecoder(resp.Body).Decode(&class)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusOK)
		return
	}

	ctx.HTML(http.StatusOK, "detail", gin.H{
		"user":       user,
		"class":      class,
		"server_url": SERVER_URL,
	})
}
