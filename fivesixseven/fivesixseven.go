package fivesixseven

import (
	"sort"
	"strings"
)

func CheckInclusion(s1 string, s2 string) bool {
	s1Split := strings.Split(s1, "")
	sort.Strings(s1Split)
	s1Sorted := strings.Join(s1Split, "")
	if len(s1) > len(s2) {
		return false
	}
	for i := 0; i <= len(s2)-len(s1); i++ {
		s2Sub := s2[i : i+len(s1)]
		s2SubSplit := strings.Split(s2Sub, "")
		sort.Strings(s2SubSplit)
		s2SubSorted := strings.Join(s2SubSplit, "")
		if s1Sorted == s2SubSorted {
			return true
		}
	}
	return false
}
