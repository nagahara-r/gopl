// Surface は 3-D 面の関数のSVGレンダリングを計算します。
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320            // キャンバスの大きさ（画素数）
	cells         = 100                 // 格子のマス目の数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x 単位 および y 単位当たりの画素数
	zscale        = height * 4          //* 0.4        // z 単位当たりの画素数
	angle         = math.Pi / 6         // x, y軸の角度 (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf(getSurface())
}

func getSurface() string {
	str := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok1 := corner(i+1, j)
			bx, by, ok2 := corner(i, j)
			cx, cy, ok3 := corner(i, j+1)
			dx, dy, ok4 := corner(i+1, j+1)
			if !(ok1 && ok2 && ok3 && ok4) {
				// 一つでも座標生成に失敗していたら、ポリゴンは作らない。
				continue
			}

			str += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	str += fmt.Sprintln("</svg>")

	return str
}

func corner(i, j int) (float64, float64, bool) {
	// マス目（i, j）の角の点 (x,y) を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さ z の計算。
	z, ok := f(x, y)
	if !ok {
		return 0, 0, false
	}

	// (x,y,z) を 2-D SVGキャンバス (sx,sy) へ等角的に投影。
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // (0,0)からの距離
	r = math.Sin(r) / r

	if math.IsNaN(r) {
		return r, false
	}
	return r, true
}
