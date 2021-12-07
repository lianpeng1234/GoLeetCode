package main

import (
	"container/list"
	"sort"
)

// 数组区间，答案错误
func eraseOverlapIntervals(intervals [][]int) int {
	// 按照每行的第一个元素排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 变长数组会议室列表
	timeList := []*list.List{}

	for i := 0; i < len(intervals); i++ {
		start := intervals[i][0]

		add := false
		for j := 0; j < len(timeList); j++ {
			aaa := timeList[j].Back().Value.([]int)
			if aaa[1] <= start {
				timeList[j].PushBack(intervals[i])
				add = true
				break
			}
		}
		if !add {
			list := list.List{}
			list.PushBack(intervals[i])
			timeList = append(timeList, &list)
		}
	}

	max := 0
	for i := 0; i < len(timeList); i++ {
		if max < timeList[i].Len() {
			max = timeList[i].Len()
		}
	}

	return len(intervals) - max
}
