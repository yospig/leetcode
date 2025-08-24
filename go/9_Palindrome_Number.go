package main
import(
    "fmt"
    "strconv"
)

func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    runes := []rune(strconv.Itoa(x))
    for i, j := 0, len(runes)-1;i < j;i, j = i+1, j-1 {
        if runes[i] != runes[j] {
            return false
        }
    }
    return true
}

func main() {
    fmt.Println(isPalindrome(121))  // true
    fmt.Println(isPalindrome(-121)) // false
    fmt.Println(isPalindrome(10))   // false
}
