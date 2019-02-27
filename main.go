package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"text/template"
	"time"
)

func sosista(valikud map[string]string) {
	rand.Seed(time.Now().UTC().UnixNano())
	juhuArv := rand.Intn(len(valikud))
	var i int
	for key, value := range valikud {
		if i == juhuArv {
			fmt.Printf("\n\n\nValitavate erakondade hulgast (%d) kostub juhuslik valimissosin:\n%7s - %s\n", len(valikud), key, value)
		}
		i++
	}
}

func main() {
	var valitavad = map[string]string{
		"RE":   "Eesti Reformierakond",
		"K":    "Eesti Keskerakond",
		"SDE":  "Sotsiaaldemokraatlik Erakond",
		"IRL":  "Erakond Isamaa ja Res Publica Liit",
		"EVA":  "Eesti Vabaerakond",
		"EKRE": "Eesti Konservatiivne Rahvaerakond",
		"EER":  "Erakond Eestimaa Rohelised",
		"RÜE":  "Rahva Ühtsuse Erakond",
		"EIP":  "Eesti Iseseisvuspartei",
		"EÜVP": "Eestimaa Ühendatud Vasakpartei",
		"EVP":  "Eesti Vabaduspartei - Põllumeeste Kogu", // EVP-PK
	}
	var kindelEi = map[string]string{}

	tmpl := template.Must(template.New("base").Parse(base))
	reader := bufio.NewReader(os.Stdin)
	UPDATE := true

	for {
		if UPDATE {
			if len(valitavad) > 0 {
				data := Data{Pealkiri: "VALITAVAD ERAKONNAD", Sisu: valitavad}
				tmpl.ExecuteTemplate(os.Stdout, "base", data)
			} else {
				fmt.Println("Ei ole ühtegi valitavat erakonda\n==================================================")
			}

			if len(kindelEi) > 0 {
				data := Data{Pealkiri: "VÄLISTATUD ERAKONNAD", Sisu: kindelEi}
				tmpl.ExecuteTemplate(os.Stdout, "base", data)
			} else {
				fmt.Println("Ei ole ühtegi välistatud erakonda\n==================================================")
			}

			fmt.Println(juhised)
		}

		juhis, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Sisendi lugemine ebaõnnestus, palun proovige uuesti.")
			continue
		}
		juhis = strings.TrimSpace(strings.ToUpper(juhis))
		switch juhis {
		case "SOSISTA":
			sosista(valitavad)
			UPDATE = false
		case "AITAB":
			os.Exit(0)
		default:
			if valitavad[juhis] != "" {
				kindelEi[juhis] = valitavad[juhis]
				delete(valitavad, juhis)
				UPDATE = true
			} else if kindelEi[juhis] != "" {
				valitavad[juhis] = kindelEi[juhis]
				delete(kindelEi, juhis)
				UPDATE = true
			} else {
				fmt.Printf("Erakonda lühendiga \"%s\" ei leitud valitavate ega välistatute hulgast\n", juhis)
				UPDATE = false

			}

		}
	}
}

var base = `{{define "base"}}
{{ .Pealkiri }}
==================================================
{{ range $key, $value := .Sisu -}}
	{{printf "%7s - %s\n" $key $value}}
{{- end -}}
==================================================
{{ end }}
`

var juhised string = `Sisesta:
* Erakonna lühend, et liigutada erakonda välistatute hulka või tagasi
	Näiteks "np" välistamaks "näitepartei"
* "sosista", et kuulda valimissosinat
* "aitab", et väljuda`

type Data struct {
	Pealkiri string
	Sisu     map[string]string
}
