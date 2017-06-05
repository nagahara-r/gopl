// Copyright © 2017 Yuki Nagahara

package eval

import "testing"

func TestAryEval(t *testing.T) {
	tests := []struct {
		expr     Expr
		env      Env
		expected float64
	}{
		{
			ary{"min", []Expr{literal(3), literal(2), literal(1), literal(-1), literal(4)}},
			Env{},
			-1,
		},
		{
			ary{"max", []Expr{literal(3), call{"pow", []Expr{literal(2), literal(3)}}, literal(1), literal(-1), literal(4)}},
			Env{},
			8,
		}, {
			ary{"sum", []Expr{Var("A"), call{"pow", []Expr{literal(2), literal(3)}}, literal(1), literal(-1), literal(4)}},
			Env{"A": 3},
			15,
		}, {
			ary{"sum", []Expr{}},
			Env{"A": 3},
			0,
		},
	}

	for _, test := range tests {
		result := test.expr.Eval(test.env)
		if test.expected != result {
			t.Errorf("expected = %v, Evel(env) = %v", test.expected, result)
		}
	}
}
