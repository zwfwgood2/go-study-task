package basetype

import (
	"fmt"
	"slices"
)

func RunMerge(intervals [][]int) {
	ans, e := merge(intervals)
	fmt.Printf("合并后切片为:%v;e%v\n", ans, e)
}
func merge(intervals [][]int) (ans [][]int, e error) {
	slices.SortFunc(intervals, func(p, q []int) int { return p[0] - q[0] }) // 按照左端点从小到大排序
	for _, p := range intervals {
		m := len(ans)
		if m > 0 && p[0] <= ans[m-1][1] { // 可以合并
			ans[m-1][1] = max(ans[m-1][1], p[1]) // 更新右端点最大值
		} else { // 不相交，无法合并
			ans = append(ans, p) // 新的合并区间
		}
	}
	return
}
