// eggsbox は鶏卵の箱を描画します。
package main

import (
	"fmt"
	"log"
	"math"

	"github.com/naga718/golang-practice/ch03/ex03/polygon"
)

const (
	width, height = 600, 320            // キャンバスの大きさ（画素数）
	cells         = 100                 // 格子のマス目の数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x 単位 および y 単位当たりの画素数
	zscale        = height * 0.4        // z 単位当たりの画素数
	angle         = math.Pi / 6         // x, y軸の角度 (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf(getSurface())
}

func getSurface() string {

	polygons := make([]polygon.Polygon, 0)
	str := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var p polygon.Polygon
			oks := []bool{false, false, false, false}
			p.Ax, p.Ay, oks[0] = corner(i+1, j)
			p.Bx, p.By, oks[1] = corner(i, j)
			p.Cx, p.Cy, oks[2] = corner(i, j+1)
			p.Dx, p.Dy, oks[3] = corner(i+1, j+1)
			if !(oks[0] && oks[1] && oks[2] && oks[3]) {
				// 一つでも座標生成に失敗していたら、ポリゴンは作らない。
				continue
			}

			polygons = append(polygons, p)

			//str += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			//	ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	maxHeight := polygon.MaxHeight(polygons)
	minHeight := polygon.MinHeight(polygons)

	log.Printf("%v, %v", maxHeight, minHeight)

	for _, p := range polygons {
		str += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'",
			p.Ax, p.Ay, p.Bx, p.By, p.Cx, p.Cy, p.Dx, p.Dy)

		if polygon.IsMaxHeight(maxHeight, p) {
			str += fmt.Sprintf(" fill='red'/>\n")
		} else if polygon.IsMinHeight(minHeight, p) {
			str += fmt.Sprintf(" fill='blue'/>\n")
		} else {
			str += fmt.Sprintf("/>\n")
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
	pow1 := math.Pow(2, math.Sin(y))
	pow2 := math.Pow(2, math.Sin(x))
	r := pow1 * pow2 / 12

	if math.IsNaN(r) {
		return r, false
	}
	return r, true
}
