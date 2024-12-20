package main

import "net/http"

func loginPage(w http.ResponseWriter, r *http.Request) {
	exeTmpl(w, r, nil, "loginPage.tmpl")
}
