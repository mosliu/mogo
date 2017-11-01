package mosDict

import (
    "unicode"
)

func IsChinese(word string) bool {
    return isChinese(word)
}
func isChinese(word string) bool {
    for _, r := range word {
        if unicode.Is(unicode.Scripts["Han"], r) {
            return true
        }
    }
    return false
}
