package main

import (
	"sort"
	"testing"
)

func BenchmarkMultiTrackSort(b *testing.B) {
	var testtracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"aaa", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"bbb", "bbb", "Smash", 2011, length("50m20s")},
		{"ccc", "title title", "Smash", 2011, length("40m54s")},
		{"ddd", "Martin Solveig", "Smash", 2011, length("6m44s")},
	}

	for i := 0; i < b.N; i++ {
		MultiTrackSort(byYear(testtracks), byArtist(testtracks))
	}

}

func BenchmarkStableSort(b *testing.B) {
	var testtracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
		{"Go", "Moby", "Moby", 1992, length("3m37s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"aaa", "Martin Solveig", "Smash", 2011, length("4m24s")},
		{"bbb", "bbb", "Smash", 2011, length("50m20s")},
		{"ccc", "title title", "Smash", 2011, length("40m54s")},
		{"ddd", "Martin Solveig", "Smash", 2011, length("6m44s")},
	}

	for i := 0; i < b.N; i++ {
		sort.Stable(byArtist(testtracks))
		sort.Stable(byYear(testtracks))
	}
}
