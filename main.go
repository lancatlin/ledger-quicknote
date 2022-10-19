package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
	"time"
)

var ledgerTpl *template.Template
var htmlTpl *template.Template

const LEDGER_FILE = "test.txt"

type TxData struct {
	Name    string
	Date    string
	Amount  string
	Account string
}

func init() {
	ledgerTpl = template.Must(template.ParseGlob("tx/*.txt"))
	log.Println(ledgerTpl.DefinedTemplates())
	htmlTpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		htmlTpl.ExecuteTemplate(w, "index.html", ledgerTpl)
	})

	http.HandleFunc("/new", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		fmt.Println(r.Form)
		tx, err := newTx(r.Form)
		if err != nil {
			http.Error(w, err.Error(), 400)
			log.Println(err, r.Form)
			return
		}
		if err := htmlTpl.ExecuteTemplate(w, "new.html", struct {
			Tx string
		}{tx}); err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), 400)
			return
		}

		tx := r.FormValue("tx")
		if err := appendToFile(tx); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		if err := htmlTpl.ExecuteTemplate(w, "success.html", struct {
			Tx string
		}{tx}); err != nil {
			http.Error(w, err.Error(), 500)
		}
	})

	log.Println("Listen on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func newTx(params url.Values) (result string, err error) {
	name := params.Get("action")
	data := TxData{
		Date:    time.Now().Format("2006/01/02"),
		Amount:  params.Get("amount"),
		Account: params.Get("account"),
		Name:    params.Get("name"),
	}
	var buf bytes.Buffer
	err = ledgerTpl.ExecuteTemplate(&buf, name, data)
	return buf.String(), nil
}

func appendToFile(tx string) (err error) {
	f, err := os.OpenFile(LEDGER_FILE, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := strings.NewReader(strings.ReplaceAll(tx, "\r", "")) // Remove CR generated from browser
	_, err = io.Copy(f, buf)
	return err
}
