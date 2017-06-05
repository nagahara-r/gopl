// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Copyright © 2017 Yuki Nagahara
// 練習7-13：expr.String() を実装し、構文ツリーをParse可能な文字列に変換します。

// See page 198.

// Package eval provides an expression evaluator.
package eval

import (
	"fmt"
	"math"
	"strings"
)

//!+env

type Env map[Var]float64

//!-env

//!+Eval1

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

//!-Eval1

//!+Eval2

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (a ary) Eval(env Env) float64 {
	switch a.fn {
	case "min":
		var min float64
		for i, ar := range a.args {
			if i == 0 || min > ar.Eval(env) {
				min = ar.Eval(env)
			}
		}
		return min
	case "max":
		var max float64
		for i, ar := range a.args {
			if i == 0 || max < ar.Eval(env) {
				max = ar.Eval(env)
			}
		}
		return max
	case "sum":
		var sum float64
		for _, ar := range a.args {
			sum += ar.Eval(env)
		}
		return sum
	}
	panic(fmt.Sprintf("unsupported array call: %s", a.fn))
}

//!-Eval2

// strings
func (b binary) String() string {
	switch b.op {
	case '+':
		return "(" + b.x.String() + " + " + b.y.String() + ")"
	case '-':
		return "(" + b.x.String() + " - " + b.y.String() + ")"
	case '*':
		return "(" + b.x.String() + " * " + b.y.String() + ")"
	case '/':
		return "(" + b.x.String() + " / " + b.y.String() + ")"
	}
	return ""
}

func (u unary) String() string {
	switch u.op {
	case '+':
		return "+" + (u.x.String())
	case '-':
		return "-" + (u.x.String())
	}
	return ""
}

func (c call) String() string {
	str := ""

	for _, arg := range c.args {
		str += arg.String()
		str += ", "
	}

	return c.fn + "(" + strings.TrimSuffix(str, ", ") + ")"
}

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprint(float64(l))
}

func (a ary) String() string {
	str := ""
	for _, ar := range a.args {
		str += ar.String() + ", "
	}

	return a.fn + "[" + strings.TrimSuffix(str, ", ") + "]"
}
