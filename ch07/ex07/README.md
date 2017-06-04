### 問題
```
練習問題7.7
20.0 のデフォルト値は ° を含んでいないのに、ヘルプメッセージが ° を含む理由を説明しなさい。
```

プログラミング言語Go (ADDISON-WESLEY PROFESSIONAL COMPUTING SERIES) (Alan A.A. Donovan (著), Brian W. Kernighan (著), 柴田 芳樹 (翻訳), p.209) より引用

### 回答
tempconv にCelsius型の String() が実装されているため。（下記）

```go
// tempcomv.go
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
```

ヘルプメッセージの生成時、デフォルト値表示はStringで生成されるため、ペルプメッセージに °C が含まれていると考えられる。

```go
// flag.go
func (f *FlagSet) Var(value Value, name string, usage string) {
	// Remember the default value as a string; it won't change.
	flag := &Flag{name, usage, value, value.String()}
```
