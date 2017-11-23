package service

import (
    "net/http"
    "github.com/unrolled/render"
)

func loginHandler(formatter *render.Render) http.HandlerFunc {

    return func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()

        formatter.HTML(w, http.StatusOK, "logined", struct {
            Username      string `json:"username"`
            Hobby 	  string `json:"hobby"`
        }{Username: req.Form["username"][0], Hobby: req.Form["hobby"][0]})
    }
}
func NotFoundHandler(formatter *render.Render) http.HandlerFunc {
    
        return func(w http.ResponseWriter, req *http.Request) {
            http.Error(w, "501 page not implemented", 501)
        }
    }