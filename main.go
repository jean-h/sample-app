/**
 * Copyright 2023 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"log"
	"net/http"
)

func main() {
	// Root handler for health checks and default landing
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received request for %s from %s", r.URL.Path, r.RemoteAddr)
		fmt.Fprintf(w, "Hello, the Go server is running!")
	})

	http.HandleFunc("/blue", blueHandler)
	log.Println("Starting server on :8080...")
	http.ListenAndServe(":8080", nil)
}

func blueHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling /blue request")
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{0, 0, 255, 255}}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}

func redHandler(w http.ResponseWriter, r *http.Request) {
	img := image.NewRGBA(image.Rect(0, 0, 100, 100))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{255, 0, 0, 255}}, image.ZP, draw.Src)
	w.Header().Set("Content-Type", "image/png")
	png.Encode(w, img)
}
