package polygon

// ポリゴン座標
type Polygon struct {
	Ax float64
	Ay float64
	Bx float64
	By float64
	Cx float64
	Cy float64
	Dx float64
	Dy float64
}

// Polygon の頂点を返します。
func MinHeight(polygons []Polygon) float64 {
	minHeight := polygons[0].Ay

	for _, polygon := range polygons {
		minHeight, _ = comp(minHeight, polygon.Ay)
		minHeight, _ = comp(minHeight, polygon.By)
		minHeight, _ = comp(minHeight, polygon.Cy)
		minHeight, _ = comp(minHeight, polygon.Dy)
	}

	return minHeight
}

// Polygon の頂点を返します。
func MaxHeight(polygons []Polygon) float64 {
	maxHeight := polygons[0].Ay

	for _, polygon := range polygons {
		_, maxHeight = comp(maxHeight, polygon.Ay)
		_, maxHeight = comp(maxHeight, polygon.By)
		_, maxHeight = comp(maxHeight, polygon.Cy)
		_, maxHeight = comp(maxHeight, polygon.Dy)
	}

	return maxHeight
}

// 現在のポリゴンが頂点であれば trueを返します
func IsMaxHeight(maxHeight float64, polygon Polygon) bool {
	var polygons []Polygon = []Polygon{polygon}
	return maxHeight == MaxHeight(polygons)
}

// 現在のポリゴンが谷であれば trueを返します
func IsMinHeight(minHeight float64, polygon Polygon) bool {
	var polygons []Polygon = []Polygon{polygon}
	return minHeight == MinHeight(polygons)
}

// 比較結果を大きい値、小さい値で返します。
func comp(a float64, b float64) (less float64, greater float64) {
	if a < b {
		return a, b
	} else {
		return b, a
	}
}
