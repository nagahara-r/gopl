// Copyright © 2017 Yuki Nagahara

// 練習7-16 Webベース電卓

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/naga718/golang-practice/ch07/ex16/eval"
)

var exprTemp = template.Must(template.New("expr").Parse(`
<html>
<head>
<title>Go Calculator</title>
</head>
<body>
<h1>Calculator</h1>
<form action="." method="get"></p>
<input type="text" name="expr" value="{{.Expr}}" size="100" /></p>
{{range .Vars}}
	{{.Name}} = <input type="text" name="${{.Name}}" value="{{.Value}}" size="20" /></p>
{{end}}
 <input type="submit" /></p>
{{if .Done}}
 = <b>{{.Value}}</b></p>
{{end}}
Detecting Error: <b> {{.Err}} </b>
</body></html>
`))

type htmldata struct {
	Expr  string
	Err   string
	Vars  []htmlvar
	Value float64
	Done  bool
}

type htmlvar struct {
	Name  string
	Value float64
}

func main() {
	http.HandleFunc("/", process)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func process(w http.ResponseWriter, req *http.Request) {
	htmldata := htmldata{}
	expstr := req.URL.Query().Get("expr")

	htmldata.Expr = expstr
	if htmldata.Expr == "" {
		htmldata.Expr = "1 + pow(2, 10)"
		exprTemp.Execute(w, &htmldata)
		return
	}

	expr, err := eval.Parse(expstr)
	if err != nil {
		htmldata.Err = err.Error()
		exprTemp.Execute(w, &htmldata)
		return
	}

	vars := eval.VarSearch(expr, nil)

	htmldata.Vars, err = parseVars(req.URL.Query(), vars)
	if err != nil {
		htmldata.Err = err.Error()
		exprTemp.Execute(w, &htmldata)
		return
	}

	env := getEnv(htmldata.Vars)

	htmldata.Value = expr.Eval(env)
	htmldata.Done = true
	exprTemp.Execute(w, &htmldata)
}

func parseVars(query url.Values, vars []string) (hvar []htmlvar, err error) {
	errs := []error{}
	for _, v := range vars {
		q := query.Get("$" + v)
		if q == "" {
			//env[eval.Var(v)] = 0
			hvar = append(hvar, htmlvar{v, 0})
			errs = append(errs, fmt.Errorf("Var value isn't set: %v\n", v))
			continue
		}

		f, err := strconv.ParseFloat(q, 64)
		if err != nil {
			hvar = append(hvar, htmlvar{v, 0})
			errs = append(errs, fmt.Errorf("Invalid Value: %v = %v\n", v, q))
			continue
		}

		hvar = append(hvar, htmlvar{v, f})
	}

	if len(errs) != 0 {
		errm := ""
		for _, e := range errs {
			errm += e.Error()
		}
		err = fmt.Errorf(errm)
	}

	return
}

func getEnv(hvar []htmlvar) (env eval.Env) {
	env = make(eval.Env)

	for _, v := range hvar {
		env[eval.Var(v.Name)] = v.Value
	}

	return
}
