// Copyright © 2017 Yuki Nagahara
// delstring は隣り合う文字列が一致した場合にスライス内で除去します。

package delstring

import "log"

// DeleteSideDuplicate は
func DeleteSideDuplicate(strs []string) []string {
	detect := false

	for i := range strs {
		if i >= len(strs)-1 {
			break
		}

		if strs[i] == strs[i+1] {
			log.Printf("detect: %v", strs[i])
			copy(strs[i:], strs[i+1:])

			detect = true
			break
		}
	}

	if detect {
		// まだ重複がないか探索
		return DeleteSideDuplicate(strs[:len(strs)-1])
	}

	return strs
}
