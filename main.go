package main
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main(){
	getStreak("APE3MP")
}

func handleUserStreak(w http.ResponseWriter, r *http.Request){
	username := r.PathValue("username")

	streak, err := getStreak(username)
	if err != nil{
		http.Error(w, err.Error(), http.StatusNotFound)
		return 
	}

	w.Header().Set("Content-type", "text/plain")
	fmt.Fprint(w, streak)
}
