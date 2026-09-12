package code

import (
	"fmt"
	"os"
	"strings"
	"errors"
	"path/filepath"
	"strconv"
)

func GetPathSize(path string, human bool, hidden bool, recurs bool) (string, error) {
	if hidden==false && strings.HasPrefix(path, ".") {
		return "", errors.New("файл не найден")
	}
	if path == "" {
		return "", nil
	}
	inf, err := os.Stat(path)
	var size int64
	if err != nil {
		return "", errors.New("Файл не найден")
	}
	if err == nil && inf.IsDir() {
		list_files, _ := os.ReadDir(path)
		for _, v := range list_files {
		    if (hidden==false) && (strings.HasPrefix(v.Name(),".")){
			continue
		    }
		    if v.IsDir() && recurs {
			dirSize, _ := GetPathSize(filepath.Join(path, v.Name()), human, hidden, recurs)
			dirSizeInt, _ := strconv.Atoi(strings.Split(dirSize, " ")[0])
			size += int64(dirSizeInt)
		    }
		    r, _ := v.Info()
		    size += r.Size()
		}
		endSize := outcomeSize(human, size)
		return endSize, nil

	}
	size = inf.Size()
	endSize := outcomeSize(human, size)
	return endSize, nil
}

func outcomeSize(human bool, size int64) string {
	var outcome string
	if !human || size < 1024 {
		outcome = fmt.Sprintf("%d B", size)
		return outcome
	}
	switch {
		case (size>=1024) && (size<1048576): outcome = fmt.Sprintf("%.1f KB", float32(size)/1024)
		case (size>=1048576) && (size<1073741824): outcome= fmt.Sprintf("%.1f MB", float32(size)/1048576)
		case (size>=1073741824) && (size<1099511627776): outcome=fmt.Sprintf("%.1f GB", float32(size)/1073741824)
		case (size>=1099511627776) && (size<1125899906842624): outcome=fmt.Sprintf("%.1f TB", float32(size)/1099511627776)
		case (size>=1125899906842624) && (size<1152921504606846976): outcome=fmt.Sprintf("%.1f PB", float32(size)/1125899906842624)
		case (float64(size)>=1152921504606846976) && (float64(size)<1180591620717411303424): outcome=fmt.Sprintf("%.1 EB", float64(size)/1152921504606846976)
	
	}
	return outcome
} 
