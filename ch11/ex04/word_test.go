// Copyright © 2017 Yuki Nagahara
// 練習11-4: IsPalindrome が句読点や空白を処理することをテストします。

package word

import (
	"math/rand"
	"testing"
	"time"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		// もしここで句読点が入っても問題なし（対になる句読点が取り除かれるだけなので）
		r := rune(rng.Intn(0x1000)) // '/u0999 までのランダムなルーン'
		runes[i] = r
		runes[n-1-i] = r
	}
	runes = insertSkipLiteral(rng, runes)

	return string(runes)
}

func insertSkipLiteral(rng *rand.Rand, runes []rune) (inserted []rune) {
	skipliterals := map[int]rune{
		0: ' ',
		1: '　',
		2: ',',
		3: '\n',
		4: '、',
		5: '。',
	}

	if len(runes) <= 0 {
		return runes
	}

	n := rng.Intn(5) // 最大5つ

	for i := 0; i < n; i++ {
		ri := rng.Intn(len(runes))
		literal := skipliterals[ri%6]

		inserted = append(inserted, runes[:ri]...)
		inserted = append(inserted, literal)
		inserted = append(inserted, runes[ri:]...)
	}

	return
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
