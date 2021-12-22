package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MongoURI = "mongodb+srv://PolyakovDOTTech:123abc@polyakovtechdb.n6fvv.mongodb.net/PolyakovTechDB?retryWrites=true&w=majority"

func acceptPostFromContact(w http.ResponseWriter, r *http.Request) {
	message, _ := io.ReadAll(r.Body)
	fmt.Println(string(message))
	client, _ := mongo.NewClient(options.Client().ApplyURI(MongoURI))
	fmt.Println(client)
	thankyouPage, _ := ioutil.ReadFile("./Public/HTML/thankyou.html")
	w.WriteHeader(http.StatusOK)
	w.Write(thankyouPage)
}

//Send public resource
func sendPublicResources(w http.ResponseWriter, r *http.Request) {
	//make a mime converter to convert extenstion to mime type
	mimeconverter := make(map[string]string)
	mimeconverter[".css"] = "text/css"
	mimeconverter[".jpg"] = "image/jpeg"
	mimeconverter[".svg"] = "image/svg+xml"
	mimeconverter[".js"] = "text/javascript"
	extension := "." + strings.Split(r.URL.Path, ".")[1]
	resourceToSend, err := ioutil.ReadFile("./" + r.URL.Path)
	if err != nil {
		sendError(w)
		return
	}
	w.Header().Set("content-type", string(mimeconverter[extension]))
	w.WriteHeader(http.StatusOK)
	w.Write(resourceToSend)
}

//Servs the basic polyakov.tech/ paths for get only
func serveHTMLForHomeSite(w http.ResponseWriter, r *http.Request) {
	//if public then send to public folder acces
	if strings.Contains(r.URL.Path, "/Public") && r.Method == "GET" {
		sendPublicResources(w, r)
		return
	} else if r.URL.Path == "/" {
		r.URL.Path = "/home"
	}
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		//set path
		path := "./Public/HTML" + r.URL.Path + ".html"
		path = strings.ToLower(path)
		//Reaf file check for erro and then send correspinding html
		htmlTosend, err := ioutil.ReadFile(path)
		if err != nil {
			sendError(w)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(htmlTosend)
	case "POST":
		//check if post is on contact page
		if r.URL.Path != "/contact" {
			sendError(w)
			return
		}
		acceptPostFromContact(w, r)
	default:
		sendError(w)
	}
}

//Used to send the error 404page
func sendError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	errorPage, _ := ioutil.ReadFile("./Public/HTML/404.html")
	w.Write(errorPage)
}
func sendIcon(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader((http.StatusOK))
	icon, _ := ioutil.ReadFile("./favicon.ico")
	w.Write((icon))
}

//main function for server
func main() {
	listOFHomeSite := [10]string{"/home", "/Home", "/resume", "/Resume", "/projects", "/Projects", "/contact", "/Contact", "/Public/", "/"}
	//register each polyakov.tech html page handler
	for i := 0; i < len(listOFHomeSite); i++ {
		http.HandleFunc(listOFHomeSite[i], serveHTMLForHomeSite)
	}
	//if favorite icon is request use this path
	http.HandleFunc("/favicon.ico", sendIcon)
	//run the server
	if err := http.ListenAndServe(":8081", nil); err != nil {
		fmt.Println("cant start")
	}
}
