package tempconv

// CToF は摂氏の温度を華氏へ変換します
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC は華氏の温度を摂氏へ変換します
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// 絶対温度 -> 摂氏温度
func KToC(k Kelvin) Celsius { return Celsius(k - CelsiusZeroK) }

// 摂氏温度 -> 絶対温度
func CToK(c Celsius) Kelvin { return Kelvin(c - AbsoluteZeroC) }

// 絶対温度 -> 華氏温度
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(KToC(k)*9/5 + 32) }

// 華氏温度 -> 絶対温度
func FToK(f Fahrenheit) Kelvin { return Kelvin((FToC(f) - AbsoluteZeroC)) }
