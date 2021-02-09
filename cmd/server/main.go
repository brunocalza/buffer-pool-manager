package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/brunocalza/bpm"
)

func newPage(bufferPool *bpm.BufferPoolManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		bufferPool.NewPage()

		response := bpm.NewResponse(bufferPool)

		data, _ := json.Marshal(response)

		w.Write(data)
	}
}

func flushPage(bufferPool *bpm.BufferPoolManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		pageParam := r.URL.Query().Get("page")
		if pageParam != "" {
			pageID, _ := strconv.Atoi(pageParam)
			bufferPool.FlushPage(bpm.PageID(pageID))

			response := bpm.NewResponse(bufferPool)
			data, _ := json.Marshal(response)
			w.Write(data)
		} else {
			w.Write([]byte("{}"))
		}
	}
}

func deletePage(bufferPool *bpm.BufferPoolManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		pageParam := r.URL.Query().Get("page")
		if pageParam != "" {
			pageID, _ := strconv.Atoi(pageParam)
			bufferPool.DeletePage(bpm.PageID(pageID))

			response := bpm.NewResponse(bufferPool)
			data, _ := json.Marshal(response)
			w.Write(data)
		} else {
			w.Write([]byte("{}"))
		}
	}
}

func unpinPage(bufferPool *bpm.BufferPoolManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		pageParam := r.URL.Query().Get("page")
		if pageParam != "" {
			pageID, _ := strconv.Atoi(pageParam)
			bufferPool.UnpinPage(bpm.PageID(pageID), false)

			response := bpm.NewResponse(bufferPool)
			data, _ := json.Marshal(response)
			w.Write(data)
		} else {
			w.Write([]byte("{}"))
		}
	}
}

func fetchPage(bufferPool *bpm.BufferPoolManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		pageParam := r.URL.Query().Get("page")
		if pageParam != "" {
			pageID, _ := strconv.Atoi(pageParam)
			bufferPool.FetchPage(bpm.PageID(pageID))

			response := bpm.NewResponse(bufferPool)
			data, _ := json.Marshal(response)
			w.Write(data)
		} else {
			w.Write([]byte("{}"))
		}
	}
}

func flushAll(bufferPool *bpm.BufferPoolManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		bufferPool.FlushAllpages()

		response := bpm.NewResponse(bufferPool)

		data, _ := json.Marshal(response)

		w.Write(data)
	}
}

func main() {

	clockReplacer := bpm.NewClockReplacer(bpm.MaxPoolSize)
	diskManager := bpm.NewDiskManagerMock()

	bufferPool := bpm.NewBufferPoolManager(diskManager, clockReplacer)

	http.HandleFunc("/new", newPage(bufferPool))
	http.HandleFunc("/flush", flushPage(bufferPool))
	http.HandleFunc("/delete", deletePage(bufferPool))
	http.HandleFunc("/unpin", unpinPage(bufferPool))
	http.HandleFunc("/fetch", fetchPage(bufferPool))
	http.HandleFunc("/flush-all", flushAll(bufferPool))

	log.Fatal("HTTP server error: ", http.ListenAndServe("localhost:3000", nil))
}
