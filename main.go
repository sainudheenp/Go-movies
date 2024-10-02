package main
import (
"fmt"
"log"
"encoding/json"
"math/rand"
"strconv"
"net/http"
"github.com/gorilla/mux"
)

type Movie struct  {
ID string `json:"id"`
Lsbn string `json:"lsbn"`
Title string `json:"title"`
Director *Director `json:"Director"`
}
type Director struct{
Firstname string `json:"firstname"`
Lastname string `json:"lastname"`
}
var movies[] Movie

func getMovies(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.vars(r)
	for index , item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter , r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
for _,item:= range movies{
	if item.ID == params["id"]{
		json.NewEncoder(w).Encode(item)
		return
	}
}

}

func createMovie(w http.ResponseWriter,r *http.Request){
w.Header().Set("Content-Type","application/json")
var movie Movie 
_ =json.NewDecoder(r.Body).Decode(&movie)
movie.ID =   strconv.Itoa (rand.Intn(10000))
movies = append(movies,movie )
json.NewEncoder(w).Encode(movie)


}



func updateMovie (){
	w.Header().Set("Content-Type","application/json")
	pa
	  
}



func mian(){
	r:= mux.NewRouter()

	movies = append(movies,Movie{ID:"1",Lsbn:"3232",Title:"WE",Director :&Director{Firstname:"john",Lastname:"Doe"}});

	r.HandleFunc("/movies",getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}",getMovie).Methods("GET")
	r.HandleFunc("/movies",createMovie).Methods("POST")
	r.HandleFunc("/movies/{}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{}",deleteMovie).Methods("DELTE")

	fmt.Println("Starting Server at 8000")
	log.Fatal(http.ListenAndServe(":8000",r))


}