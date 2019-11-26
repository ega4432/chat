package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func uploadHandler(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("userid")
	file, header, err := req.FormFile("avatarFile")
	if err != nil {
		_, _ = io.WriteString(w, err.Error())
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		_, _ = io.WriteString(w, err.Error())
		return
	}
	filename := filepath.Join("avatars", userId+filepath.Ext(header.Filename))
	err = ioutil.WriteFile(filename, data, 0777)
	if err != nil {
		_, _ = io.WriteString(w, err.Error())
		return
	}
	_, _ = io.WriteString(w, "成功！")
}
