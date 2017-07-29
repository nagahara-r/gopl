// Copyright © 2017 Yuki Nagahara
// 練習11-3: IsPalindrome を回文でないことをテストします。

package palindrome

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // '/u0999 までのランダムなルーン'
		runes[i] = r
		runes[n-1-i] = r
	}
	return string(runes)
}

func randomNonPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	if n <= 1 {
		// 0文字, 1文字の場合、必ず回文になるので作り直し
		return randomNonPalindrome(rng)
	}
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // '/u0999 までのランダムなルーン'
		runes[i] = r
		runes[n-1-i] = r
	}
	runes[0] = rune('a')
	runes[len(runes)-1] = rune('b')

	return string(runes)
}

func TestRandomPalindrome(t *testing.T) {
	// 擬似乱数生成器を初期化する
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%v) is false", p)
		}
	}
}

func TestRandomNonPalindrome(t *testing.T) {
	// 擬似乱数生成器を初期化する
	seed := time.Now().UTC().UnixNano()
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%v) is true", p)
		}
	}
}
