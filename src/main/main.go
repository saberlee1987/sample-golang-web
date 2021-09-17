package main

import (
	"Test12/src/dto"
	"Test12/src/hi"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

var templatePage *template.Template

func main() {
	fmt.Println("Hello World !!!!!!")
	fmt.Println(hi.SayHi("Saber", "Azizi"))

	http.HandleFunc("/", handle1)
	http.HandleFunc("/personPage", handle3)
	http.HandleFunc("/addPerson", handle4)
	http.HandleFunc("/github", handle2)
	http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))
	templatePage = template.Must(template.ParseGlob("template/*"))
	//
	http.ListenAndServe(":8086", nil)
}
func handle1(response http.ResponseWriter, request *http.Request) {
	//templatePage, err := template.ParseFiles("template/home.gohtml")
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//m := map[string]string{"firstName":"Saber","lastName":"Azizi"}
	p := dto.Person{}
	p.FirstName = "Saber"
	p.LastName = "Azizi"
	p.Books = []string{"java", "C#", "php", "python", "C++", "C", "javaScript", "perl", "go"}
	err := templatePage.Execute(response, p)
	if err != nil {
		log.Fatalln(err)
	}

}

func handle2(response http.ResponseWriter, request *http.Request) {

	body, err := callHttpMethod()
	if err != nil {
		response.Header().Add("statusCode", "500")
		response.Write([]byte(err.Error()))
	} else {
		response.Write([]byte(body))
	}
}

func handle3(response http.ResponseWriter, request *http.Request) {
	templatePage, err := template.ParseFiles("template/person.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = templatePage.Execute(response, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
func handle4(response http.ResponseWriter, request *http.Request) {
	firstName := request.PostFormValue("firstName")
	lastName := request.PostFormValue("lastName")
	books := request.PostFormValue("books")
	books = strings.TrimSpace(books)
	booksArray := strings.Split(books, "\n")

	p := dto.Person{
		FirstName: firstName,
		LastName:  lastName,
		Books:     booksArray,
	}
	dto.AddPerson(p)
	fmt.Println(firstName, lastName, books, booksArray)

	persons := dto.GetPersons()
	m := map[string][]dto.Person{}
	m["persons"] = persons
	templatePage, err := template.ParseFiles("template/person.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	err = templatePage.Execute(response, m)
	if err != nil {
		log.Fatalln(err)
	}

}
func callHttpMethod() (string, error) {
	c := http.Client{Timeout: time.Duration(10) * time.Second}

	req, err := http.NewRequest("GET", "https://api.github.com/gists/c960d211532f7c35aeb0c854892bf108", nil)
	if err != nil {
		return "", err
	}
	response, err := c.Do(req)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("response--", string(body))
	return (string(body)), nil
}
