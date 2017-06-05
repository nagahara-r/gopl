// Copyright © 2017 Yuki Nagahara

package eval

import "fmt"

func ArySample() {
	exprs := []Expr{
		ary{"min", []Expr{literal(3), literal(2), literal(1), literal(-1), literal(4)}},
		ary{"max", []Expr{literal(3), call{"pow", []Expr{literal(2), literal(3)}}, literal(1), literal(-1), literal(4)}},
		ary{"sum", []Expr{literal(3), call{"pow", []Expr{literal(2), literal(3)}}, literal(1), literal(-1), literal(4)}},
	}
	env := Env{}

	for _, e := range exprs {
		fmt.Println(e.String())
		fmt.Printf("=> %v\n", e.Eval(env))
	}
}
