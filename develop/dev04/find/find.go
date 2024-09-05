package find

import (
	"sort"
	"strings"
)

func GetMap(arr []string) map[string][]string {
	mp := make(map[string][]string, len(arr))
	for _, v := range arr {
		vvv := strings.ToLower(v)
		str := []rune(vvv)
		sort.Slice(str, func(i, j int) bool {
			return str[i] < str[j]
		})
		mp[string(str)] = append(mp[string(str)], vvv)
	}
	newMp := make(map[string][]string, len(mp))
	for k, v := range mp {
		if len(v) < 2 {
			delete(mp, k)
			continue
		}
		fstr := v[0]
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
		newMp[fstr] = append(newMp[fstr], v...)
	}

	return newMp
}
