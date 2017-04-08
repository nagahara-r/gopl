package langeconv

import "fmt"

type Metre float64
type Feet float64

func (m Metre) String() string { return fmt.Sprintf("%gm", m) }
func (ft Feet) String() string { return fmt.Sprintf("%gft", ft) }

// MToFt はメートル単位の数値をフィート単位に変換します。
func MToFt(m Metre) Feet { return Feet(m * 3.28) }

// FtToM はフィート単位をメートルに変換します。
func FtToM(ft Feet) Metre { return Metre(ft * 0.3048) }
