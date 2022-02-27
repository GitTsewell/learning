package skip_list

import (
	"math/rand"
	"time"
)

const MaxLevel = 32

const Probability = 0.25 // 基于时间与空间综合 best practice 值, 越上层概率越小

func randLevel() (level int) {
	rand.Seed(time.Now().UnixNano())
	for level = 1; rand.Float32() < Probability && level < MaxLevel; level++ {
	}
	return
}

type node struct {
	nextNodeArray []*node
	val           int
}

func newNode(val, level int) *node {
	return &node{val: val, nextNodeArray: make([]*node, level)}
}

type SkipList struct {
	head  *node
	level int
}

func Constructor() SkipList {
	return SkipList{head: newNode(0, MaxLevel), level: 1}
}

func (sl *SkipList) Add(num int) {
	current := sl.head
	update := make([]*node, MaxLevel)
	// 这一步相当于找最左边的那一个node节点
	// 循环level层数 从最上层开始循环
	for i := sl.level - 1; i >= 0; i-- {
		// 判断当前node next_node 是否是nil 如果是nil 就在当前node的new_node下add
		// 或者判断当前node的next_node的val 是否大于num 如果大于 那num 肯定要插入到当前node和next_node之间
		if current.nextNodeArray[i] == nil || current.nextNodeArray[i].val > num {
			update[i] = current
		} else { // 如果不是,那就一直移动指针 找到next_node != nil  && next_node.val < num 的地方
			for current.nextNodeArray[i] != nil && current.nextNodeArray[i].val < num {
				current = current.nextNodeArray[i]
			}
			update[i] = current
		}
	}

	// 如果this.level 大于 当前level 说明要把当前level扩充
	level := randLevel()
	if level > sl.level {
		for i := sl.level; i < level; i++ {
			update[i] = sl.head
		}
		sl.level = level
	}

	// 开始拼接右边的node节点
	node := newNode(num, level)
	for i := 0; i < level; i++ {
		node.nextNodeArray[i] = update[i].nextNodeArray[i] // 先把右边的node节点连接上
		update[i].nextNodeArray[i] = node                  // 再把左边的拼接上
	}
}

func (sl *SkipList) Search(target int) (times int, ok bool) {
	current := sl.head
	// 先从最大层开始寻找
	// 如果 == target return
	// 如果 > target break
	// 如果 < target next
	ok = false
	for i := sl.level - 1; i >= 0; i-- {
		for current.nextNodeArray[i] != nil {
			times++
			if current.nextNodeArray[i].val == target {
				ok = true
				return
			} else if current.nextNodeArray[i].val > target {
				break
			} else {
				current = current.nextNodeArray[i]
			}
		}
	}
	return
}

func (sl *SkipList) Erase(target int) (ok bool) {
	current := sl.head
	// 先从最大层开始寻找
	// 如果 == target break
	// 如果 > target break
	// 如果 < target next
	ok = false
	for i := sl.level - 1; i >= 0; i-- {
		for current.nextNodeArray[i] != nil {
			if current.nextNodeArray[i].val == target {
				tmp := current.nextNodeArray[i]
				current.nextNodeArray[i] = tmp.nextNodeArray[i]
				tmp.nextNodeArray[i] = nil
				ok = true
				break
			} else if current.nextNodeArray[i].val > target {
				break
			} else {
				current = current.nextNodeArray[i]
			}
		}
	}
	return
}
