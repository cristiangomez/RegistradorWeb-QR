// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/cristiangomez/registrdor-qr/pkg"

	"github.com/gin-gonic/gin"
	"fmt"
	"github.com/skip2/go-qrcode"
	"math/rand"
	"encoding/base64"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}

func main() {
	//flag.Parse()
	hub := pkg.NewHub()
	go hub.Run()
	/*
	http.HandleFunc("/", serveHome)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		pkg.ServeWs(hub, w, r)
	})

	http.HandleFunc("/token/{name:\\d+}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("registrado marcacion")
		png, _ := qrcode.Encode(fmt.Sprintf("http://MacBook-Pro-de-Cristian.local:8080/token/%d",rand.Intn(100)), qrcode.Medium, 256)
		//ioutil.WriteFile("dat1", png, 0644)
		hub.Broadcast<- []byte(b64.StdEncoding.EncodeToString(png))
		http.ServeFile(w, r, "websockets.html")
	})


	http.ListenAndServe(":8080", nil)
	*/
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Posts",
		})
	})
	router.GET("/ws", func(c *gin.Context) {
		pkg.ServeWs(hub, c.Writer,c.Request)
	})
	router.GET("/token/:tk", func(c *gin.Context) {
		fmt.Println("registrado marcacion")
		png, _ := qrcode.Encode(fmt.Sprintf("http://MacBook-Pro-de-Cristian.local:8080/token/%d",rand.Intn(100)), qrcode.Medium, 256)
		//ioutil.WriteFile("dat1", png, 0644)
		hub.Broadcast<- []byte(base64.StdEncoding.EncodeToString(png))
		c.String(200,"OK")
	})
	fmt.Println("iniciado")
	router.Run(":8080")
	/*if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}*/
}