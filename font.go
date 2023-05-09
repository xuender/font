package font

import (
	"os"
	"strings"

	"github.com/flopp/go-findfont"
	"github.com/samber/lo"
	"github.com/xuender/font/wqy/microhei"
)

// nolint: gochecknoglobals
var _fonts = [...]string{
	"wqy-zenhei",
	"wqy",
	"simhei.ttf",
	"simkai.ttf",
	"simsun.ttf",
	"STHeiti",
	"ukai.ttc",
	"uming.ttc",
}

func DefaultFontPath() (string, bool) {
	return lo.Find(findfont.List(), func(path string) bool {
		for _, font := range _fonts {
			if strings.Contains(path, font) {
				return true
			}
		}

		return false
	})
}

func DefaultFont() []byte {
	path, has := DefaultFontPath()
	if has {
		if ret, err := os.ReadFile(path); err == nil {
			return ret
		}
	}

	return microhei.TTF
}
