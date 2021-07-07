// 贪心算法
// 让胃口最小的先吃最小的饼干
import "fmt"
func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    child := 0
    cookies := 0
    for {
        // 当饼干没了或者小孩没了就结束
        if  child >= len(g) || cookies >= len(s) {
            break
        }
        if s[cookies] >= g[child] {
            child++
        }
        cookies++
    }
    return child
}