package api

import (
	"net/http"
	"duolingo-api/duolingo"	
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /user/{username}", handleUserStreak)
	mux.HandleFunc("GET /user/svg/{username}", handleSVG)

	mux.ServeHTTP(w, r)
}

func handleUserStreak(w http.ResponseWriter, r *http.Request){
	username := r.PathValue("username")

	streak, err := duolingo.GetStreak(username)
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}

	w.Header().Set("Content-type", "text/plain")
	fmt.Fprint(w, streak)
}

func handleSVG(w http.ResponseWriter, r *http.Request){
	username := r.PathValue("username")

	streak, err := duolingo.GetStreak(username)
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}

	svg := fmt.Sprintf(`
	<svg width="220" height="70" viewBox="0 0 220 70" xmlns="http://www.w3.org/2000/svg">
		<defs>
			<filter id="shadow" x="0" y="0" width="200%%" height="200%%">
				<feDropShadow dx="1" dy="1" stdDeviation="1" flood-color="#444" flood-opacity="0.3"/>
			</filter>
		</defs>

		<rect x="2" y="2" width="216" height="66" rx="15" fill="#58CC02" filter="url(#shadow)"/>

		<g transform="translate(15, 12) scale(0.9)">
			<path d="M22.5 0C10.07 0 0 10.07 0 22.5S10.07 45 22.5 45 45 34.93 45 22.5 34.93 0 22.5 0Zm0 41a18.5 18.5 0 1 1 18.5-18.5A18.52 18.52 0 0 1 22.5 41Z" fill="white"/>
			<circle cx="15.5" cy="18.5" r="3.5" fill="white"/>
			<circle cx="29.5" cy="18.5" r="3.5" fill="white"/>
			<path d="M22.5 28c-3.3 0-6.1-1.9-7.4-4.7l-1.3.6C15.4 27.5 18.7 30 22.5 30s7.1-2.5 8.7-6.1l-1.3-.6C28.6 26.1 25.8 28 22.5 28Z" fill="white"/>
		</g>

		<text x="65" y="30" font-family="Arial, Helvetica, sans-serif" font-size="12" fill="#FFFFFF" font-weight="bold" letter-spacing="0.5">
			Duolingo
		</text>
		<text x="65" y="48" font-family="Arial, Helvetica, sans-serif" font-size="16" fill="#FFFFFF" font-weight="bold">
			Streak
		</text>

		<circle cx="175" cy="35" r="25" fill="white" stroke="#FF9600" stroke-width="3"/>
		
		<text x="175" y="42" font-family="Arial, Helvetica, sans-serif" font-size="22" fill="#FF9600" font-weight="bold" text-anchor="middle">
			%d
		</text>
	</svg>`, streak)

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")
	
	fmt.Fprint(w, svg)
}
