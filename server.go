package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Para marcar qndo o programa começou
var startedAt = time.Now()

func main() {
	// Endpoints
	http.HandleFunc("/healthz", Healthz)
	http.HandleFunc("/secret", Secret)
	http.HandleFunc("/configmap", ConfigMap)
	http.HandleFunc("/", Hello)
	// 8000 porta do container
	http.ListenAndServe(":8000", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// Criando variáveis de ambiente 'name' e 'age'
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	fmt.Fprintf(w, "Hello, I'm %s. I'm %s.", name, age)
}

func Secret(w http.ResponseWriter, r *http.Request) {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	fmt.Fprintf(w, "User: %s. Password: %s", user, password)
}

/*
	Ler um arquivo de configuração para injetar no container de forma
	mais dinâmica e não precisar rebuildar a aplicação sempre
*/
func ConfigMap(w http.ResponseWriter, r *http.Request) {
	// Ler arquivo family.txt
	data, err := ioutil.ReadFile("/go/myfamily/family.txt")
	if err != nil {
		log.Fatalf("Error reading file: ", err)
	}
	fmt.Fprintf(w, "My Family: %s.", string(data))
}

/*
	Checando a saúde da aplicação
*/
func Healthz(w http.ResponseWriter, r *http.Request) {

	// A qnto tempo estamos online?
	duration := time.Since(startedAt)

	if duration.Seconds() < 10 {
		w.WriteHeader(500) // Falha
		w.Write([]byte(fmt.Sprintf("Duration: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200) // Sucesso
		w.Write([]byte("ok"))
	}

}
