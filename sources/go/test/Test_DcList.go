package test

import (
	"fmt"
	dc "structure/list/dclist"
)

func TestDcList() {

	fmt.Println("*********** 测试开始 ***********")
	l := dc.NewDcList()

	fmt.Println("----------- 追加元素 -----------")
	l.Append(7)
	l.Append(3)
	l.Display()

	fmt.Println("----------- 插入元素 -----------")
	l.Insert(1, 2)
	l.Insert(3, 9)
	l.Insert(5, 10)
	l.Display()

	fmt.Println("----------- 删除元素 -----------")
	l.Delete(1)
	l.Delete(4)
	l.Display()

	fmt.Println("----------- 修改元素 -----------")
	l.Update(2, 22)
	l.Display()

	fmt.Println("----------- 查找元素 -----------")
	ge, _ := l.GetElem(22)
	le, _ := l.Locate(22)
	fmt.Printf("按位序查找值：%d\n", ge)
	fmt.Printf("按值查找位序：%d\n", le)

	fmt.Println("----------- 查找前驱 -----------")
	pe, _ := l.PrevElem(22)
	fmt.Printf("元素前驱：%d\n", pe)

	fmt.Println("----------- 查找后继 -----------")
	ne, _ := l.NextElem(22)
	fmt.Printf("元素后继：%d\n", ne)

	fmt.Println("*********** 测试结束 ***********")

}