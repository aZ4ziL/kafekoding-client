package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const (
	SERVER_URL  = "http://localhost:8000"
	TOKEN_URL   = SERVER_URL + "/v1/auth/get-token"
	LOGIN_URL   = SERVER_URL + "/v1/auth/user"
	CLASSES_URL = SERVER_URL + "/v1/classes"
)

// Login
// handler for login user
func Login(ctx *gin.Context) {
	session := sessions.Default(ctx)
	if user := session.Get("user"); user != nil {
		ctx.Redirect(http.StatusFound, "/")
		return
	}

	if ctx.Request.Method == "GET" {
		defer flasher.Del()

		ctx.HTML(http.StatusOK, "login", gin.H{
			"flasher": flasher,
		})
		return
	}

	if ctx.Request.Method == "POST" {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		dataByte := bytes.NewBufferString(fmt.Sprintf(`{"username": "%s", "password": "%s"}`, username, password))
		req, err := http.NewRequest("POST", TOKEN_URL, dataByte)
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			flasher.Set("danger", err.Error())
			ctx.Redirect(http.StatusFound, "/login")
			return
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			flasher.Set("danger", err.Error())
			ctx.Redirect(http.StatusFound, "/login")
			return
		}

		defer resp.Body.Close()
		respData := resp.Body

		data := struct {
			Status  string `json:"status"`
			Message string `json:"message"`
			Token   string `json:"token"`
		}{}
		_ = json.NewDecoder(respData).Decode(&data)
		if data.Status == "success" {
			req, err := http.NewRequest("GET", LOGIN_URL, nil)
			if err != nil {
				flasher.Set("danger", err.Error())
				ctx.Redirect(http.StatusFound, "/login")
				return
			}
			req.Header.Set("Accept", "application/json")
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Authorization", "Bearer "+data.Token)

			client := http.Client{}
			client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
				for key, val := range via[0].Header {
					req.Header[key] = val
				}
				return err
			}
			resp, err := client.Do(req)
			if err != nil {
				flasher.Set("danger", err.Error())
				ctx.Redirect(http.StatusFound, "/login")
				return
			}
			defer resp.Body.Close()
			userData := &UserPayload{}
			_ = json.NewDecoder(resp.Body).Decode(&userData)
			if resp.StatusCode == 200 {
				userData.Token = data.Token
				session.Set("user", userData)
				_ = session.Save()
				// check if query redirect
				redirectTo := ctx.Request.URL.Query().Get("redirect_to")
				log.Println(redirectTo)
				if redirectTo != "" {
					ctx.Redirect(http.StatusFound, redirectTo)
					return
				} else {
					ctx.Redirect(http.StatusFound, "/")
					return
				}
			} else {
				flasher.Set("danger", "Your session has ended.")
				ctx.Redirect(http.StatusFound, "/login")
				return
			}
		} else {
			flasher.Set(data.Status, data.Message)
			ctx.Redirect(http.StatusFound, "/login")
			return
		}
	}
}

// Logout Handler
func Logout(ctx *gin.Context) {
	session := sessions.Default(ctx)
	session.Delete("user")
	session.Clear()
	session.Save()
	ctx.Redirect(http.StatusFound, "/login")
}
