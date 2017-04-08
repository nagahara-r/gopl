package weightconv

import "fmt"

type Pound float64
type Kilo float64

func (lb Pound) String() string { return fmt.Sprintf("%glb", lb) }
func (kg Kilo) String() string  { return fmt.Sprintf("%gKg", kg) }

// LbToKg はポンドをキログラムに変換します。
func LbToKg(lb Pound) Kilo { return Kilo(lb * 0.45) }

// KgToLb はキログラムをポンドに変換します。
func KgToLb(kg Kilo) Pound { return Pound(kg * 2.2) }
