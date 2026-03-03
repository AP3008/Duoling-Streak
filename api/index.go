package api

import (
	_ "embed"
	"fmt"
	"net/http"
	"strings"

	"duolingo-api/duolingo"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /user/{username}", handleUserStreak)
	mux.HandleFunc("GET /user/svg/{username}", handleSVG)

	mux.ServeHTTP(w, r)
}

//go:embed duolingo-bird.svg
var birdSVG string

func handleUserStreak(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	streak, err := duolingo.GetStreak(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-type", "text/plain")
	fmt.Fprint(w, streak)
}

func handleSVG(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")

	streak, err := duolingo.GetStreak(username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	nestedBird := strings.Replace(birdSVG, `<svg viewBox="0 0 237 227" fill="none" xmlns="http://www.w3.org/2000/svg">`, `<svg x="60" y="0" width="155" height="140" viewBox="0 0 237 227" preserveAspectRatio="xMidYMid meet" fill="none" xmlns="http://www.w3.org/2000/svg">`, 1)

	svg := fmt.Sprintf(`<svg width="400" height="140" viewBox="0 0 400 140" fill="none" xmlns="http://www.w3.org/2000/svg">
	%s
	<g transform="translate(292, 70)">
		<circle r="52" stroke="#FF9600" stroke-width="4" fill="none"/>
		<text x="0" y="5" text-anchor="middle" font-family="Segoe UI, Arial, sans-serif" font-weight="800" font-size="38" fill="#FF9600">%d</text>
		<text x="0" y="23" text-anchor="middle" font-family="Segoe UI, Arial, sans-serif" font-weight="700" font-size="13" fill="#555" letter-spacing="2">DAYS</text>
	</g>
</svg>`, nestedBird, streak)

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	fmt.Fprint(w, svg)}
