// eggsbox は鶏卵の箱を描画します。
package main

import (
	"fmt"
	"log"
	"math"
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

// ポリゴン構造体
type Polygon struct {
	Ax float64
	Ay float64
	Az float64
	Bx float64
	By float64
	Bz float64
	Cx float64
	Cy float64
	Cz float64
	Dx float64
	Dy float64
	Dz float64
}

func main() {
	fmt.Printf(getSurface())
}

func getSurface() string {

	polygons := make([]Polygon, 0)
	str := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			var p Polygon
			oks := []bool{false, false, false, false}
			p.Ax, p.Ay, p.Az, oks[0] = corner(i+1, j)
			p.Bx, p.By, p.Bz, oks[1] = corner(i, j)
			p.Cx, p.Cy, p.Cz, oks[2] = corner(i, j+1)
			p.Dx, p.Dy, p.Dz, oks[3] = corner(i+1, j+1)
			if !(oks[0] && oks[1] && oks[2] && oks[3]) {
				// 一つでも座標生成に失敗していたら、ポリゴンは作らない。
				continue
			}

			polygons = append(polygons, p)
		}
	}

	maxHeight := maxHeight(polygons)
	minHeight := minHeight(polygons)

	log.Printf("%v, %v", maxHeight, minHeight)

	for _, p := range polygons {
		str += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'",
			p.Ax, p.Ay, p.Bx, p.By, p.Cx, p.Cy, p.Dx, p.Dy)

		if isMaxHeight(maxHeight, p) {
			str += fmt.Sprintf(" fill='#ff0000'/>\n")
		} else if isMinHeight(minHeight, p) {
			str += fmt.Sprintf(" fill='#0000ff'/>\n")
		} else {
			str += fmt.Sprintf("/>\n")
		}
	}

	str += fmt.Sprintln("</svg>")

	return str
}

func corner(i, j int) (float64, float64, float64, bool) {
	// マス目（i, j）の角の点 (x,y) を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さ z の計算。
	z, ok := f(x, y)
	if !ok {
		return 0, 0, 0, false
	}

	// (x,y,z) を 2-D SVGキャンバス (sx,sy) へ等角的に投影。
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z, true
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // (0,0)からの距離
	r = math.Sin(r) / r

	if math.IsNaN(r) {
		return r, false
	}
	return r, true
}

func minHeight(polygons []Polygon) float64 {
	minHeight := polygons[0].Az

	for _, polygon := range polygons {
		minHeight, _ = comp(minHeight, polygon.Az)
		minHeight, _ = comp(minHeight, polygon.Bz)
		minHeight, _ = comp(minHeight, polygon.Cz)
		minHeight, _ = comp(minHeight, polygon.Dz)
	}

	return minHeight
}

func maxHeight(polygons []Polygon) float64 {
	maxHeight := polygons[0].Az

	for _, polygon := range polygons {
		_, maxHeight = comp(maxHeight, polygon.Az)
		_, maxHeight = comp(maxHeight, polygon.Bz)
		_, maxHeight = comp(maxHeight, polygon.Cz)
		_, maxHeight = comp(maxHeight, polygon.Dz)
	}

	return maxHeight
}

func isMaxHeight(max float64, polygon Polygon) bool {
	var polygons []Polygon = []Polygon{polygon}
	return max == maxHeight(polygons)
}

func isMinHeight(min float64, polygon Polygon) bool {
	var polygons []Polygon = []Polygon{polygon}
	return min == minHeight(polygons)
}

func comp(a float64, b float64) (less float64, greater float64) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
