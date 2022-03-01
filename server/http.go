package main

import (
	"encoding/json"
	"net/http"
)

func getClips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(readJson())
	errHanding(err)
}

func addClip(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	clips := readJson()
	var newClips Clipboard
	_ = json.NewDecoder(r.Body).Decode(&newClips)
	clips = append(clips, newClips)
	writeJson(clips)
	err := json.NewEncoder(w).Encode(newClips)
	errHanding(err)
}

//func deleteClip(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	params := mux.Vars(r)
//	clips := readJson()
//	for index, clip := range clips {
//		paraIndex, err := strconv.Atoi(params["index"])
//		errHanding(err)
//		if paraIndex == clip.Index {
//			clips = append(clips[:index], clips[index+1:]...)
//			break
//		}
//	}
//	writeJson(clips)
//	err := json.NewEncoder(w).Encode(clips)
//	errHanding(err)
//}

func clearClips(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	clear()
	clips := readJson()
	err := json.NewEncoder(w).Encode(clips)
	errHanding(err)
}

// TODO In a web server, if you panic you miss the opportunity to respond appropriately to the client with say, an http 500
