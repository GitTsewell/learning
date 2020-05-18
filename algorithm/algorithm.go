package algorithm

import (
	"fmt"
)

// 冒泡排序
// 从第一个原始开始  两两比较 如果第一个比第二个大就交换他们的位置
// O(n^2)
func maopao(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		flag := false
		for j := 1; j < len(buf)-i; j++ {
			if buf[j-1] > buf[j] {
				times++
				buf[j-1], buf[j] = buf[j], buf[j-1]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
	fmt.Println("maopao times: ", times)
}

// 选择排序
// 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置。
// 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
// 重复第二步，直到所有元素均排序完毕。
// O(n^2)
func xuanze(buf []int) {
	times := 0
	for i := 0; i < len(buf)-1; i++ {
		min := i
		for j := i; j < len(buf); j++ {
			times++
			if buf[min] > buf[j] {
				min = j
			}
		}
		if min != i {
			buf[min], buf[i] = buf[i], buf[min]
		}
	}
	fmt.Println("xuanze times: ", times)
}

// 插入排序
// 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
// 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置
// O(n^2)
func charu(buf []int) {
	times := 0
	for i := 1; i < len(buf); i++ {
		for j := i; j > 0; j-- {
			if buf[j] < buf[j-1] {
				times++
				buf[j-1], buf[j] = buf[j], buf[j-1]
			} else {
				break
			}
		}
	}
	fmt.Println("charu times: ", times)
}

// 希尔排序
// 先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，待整个序列中的记录"基本有序"时，再对全体记录进行依次直接插入排序
// // O(n^1.3)
func xier(buf []int) {
	times := 0
	tmp := 0
	length := len(buf)
	incre := length
	// fmt.Println("buf: ", buf)
	for {
		incre /= 2
		for k := 0; k < incre; k++ { //根据增量分为若干子序列
			for i := k + incre; i < length; i += incre {
				for j := i; j > k; j -= incre {
					// fmt.Println("j: ", j, " data: ", buf[j], " j-incre: ", j-incre, " data: ", buf[j-incre])
					times++
					if buf[j] < buf[j-incre] {
						tmp = buf[j-incre]
						buf[j-incre] = buf[j]
						buf[j] = tmp
					} else {
						break
					}
				}
				// fmt.Println("middle: ", buf)
			}
			// fmt.Println("outer: ", buf)
		}
		// fmt.Println("outer outer: ", buf, " incre: ", incre)

		if incre == 1 {
			break
		}
	}
	// fmt.Println("after: ", buf)
	fmt.Println("xier times: ", times)
}

// 快速排序
// 从数列中挑出一个元素，称为 "基准"（pivot）;
// 重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
// 递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序
func kuaisu(buf []int) {
	kuai(buf, 0, len(buf)-1)
}

func kuai(a []int, l, r int) {
	if l >= r {
		return
	}
	i, j, key := l, r, a[l] //选择第一个数为key
	for i < j {
		for i < j && a[j] > key { //从右向左找第一个小于key的值
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}
		for i < j && a[i] < key { //从左向右找第一个大于key的值
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	//i == j
	a[i] = key
	kuai(a, l, i-1)
	kuai(a, i+1, r)
}

//归并排序
func guibing(buf []int) {
	tmp := make([]int, len(buf))
	merge_sort(buf, 0, len(buf)-1, tmp)
}

func merge_sort(a []int, first, last int, tmp []int) {
	if first < last {
		middle := (first + last) / 2
		merge_sort(a, first, middle, tmp)       //左半部分排好序
		merge_sort(a, middle+1, last, tmp)      //右半部分排好序
		mergeArray(a, first, middle, last, tmp) //合并左右部分
	}
}

func mergeArray(a []int, first, middle, end int, tmp []int) {
	// fmt.Printf("mergeArray a: %v, first: %v, middle: %v, end: %v, tmp: %v\n",
	//     a, first, middle, end, tmp)
	i, m, j, n, k := first, middle, middle+1, end, 0
	for i <= m && j <= n {
		if a[i] <= a[j] {
			tmp[k] = a[i]
			k++
			i++
		} else {
			tmp[k] = a[j]
			k++
			j++
		}
	}
	for i <= m {
		tmp[k] = a[i]
		k++
		i++
	}
	for j <= n {
		tmp[k] = a[j]
		k++
		j++
	}

	for ii := 0; ii < k; ii++ {
		a[first+ii] = tmp[ii]
	}
	// fmt.Printf("sort: buf: %v\n", a)
}

// 堆排序
func duipai(buf []int) {
	temp, n := 0, len(buf)

	for i := (n - 1) / 2; i >= 0; i-- {
		MinHeapFixdown(buf, i, n)
	}

	for i := n - 1; i > 0; i-- {
		temp = buf[0]
		buf[0] = buf[i]
		buf[i] = temp
		MinHeapFixdown(buf, 0, i)
	}
}

func MinHeapFixdown(a []int, i, n int) {
	j, temp := 2*i+1, 0
	for j < n {
		if j+1 < n && a[j+1] < a[j] {
			j++
		}

		if a[i] <= a[j] {
			break
		}

		temp = a[i]
		a[i] = a[j]
		a[j] = temp

		i = j
		j = 2*i + 1
	}
}
