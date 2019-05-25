package main

import "net/http"
import "encoding/json"
import "strings"
type weatherData struct {
	Name string `json:"name"`
	Main struct {
		Kelvin float64 `json:"temp"`
	} `json:"main"`
 }

func main(){
	
	http.HandleFunc("/hi", hi)
	
	http.HandleFunc("/weather/", func(w http.ResponseWriter, r *http.Request) {
		city := strings.SplitN(r.URL.Path, "/", 3)[2]
	 
		data, err := query(city)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	 
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(data)
	 })
	 http.HandleFunc("/", hello)
	http.ListenAndServe(":5050",nil)

}
func hello(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("hello!"))
}
func hi(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("hello Andela!"))
}
func query(city string) (weatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?APPID=7a2d3fb47d3d5a2ab03fc9db29ff3819&q=" + city)
	if err != nil {
		return weatherData{}, err
	}
 
	defer resp.Body.Close()
 
	var d weatherData
 
	if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
		return weatherData{}, err
	}
 
	return d, nil
 }