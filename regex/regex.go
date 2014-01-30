package regex

import (
	"regexp"
)

const (
	KVPAIR_CAP = 16
)

type KVPair map[string]string

func MatchAll(r *regexp.Regexp, data string) (captures []KVPair, ok bool) {
	captures = make([]KVPair, 0, KVPAIR_CAP)
	names := r.SubexpNames()
	length := len(names)
	matches := r.FindAllStringSubmatch(data, -1)
	for _, match := range matches {
		cmap := make(KVPair, length)
		for pos, val := range match {
			name := names[pos]
			if name != "" {
				cmap[name] = val
			}
		}
		captures = append(captures, cmap)
	}
	if len(captures) > 0 {
		ok = true
	}
	return
}
