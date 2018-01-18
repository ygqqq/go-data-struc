package list
import(
	"fmt"
)

type Student struct{
	Name string
	Age int
}
type Node struct{
	Student *Student
	next *Node
}
type List struct{
	size int
	first *Node
	last *Node
}

func (list *List)Init(){
	list.size = 0
	list.first , list.last = nil , nil
}

func (list *List)GetSize() int{
	return list.size
}
func (list *List)GetNode(index int) *Node{
	if index <= 0 || index > list.size{
		return nil
	}
	node := list.first
	for i:=1;i<index;i++{
		node =node.next
	}
	return node
}
func (list *List)Append(stu *Student) bool{
	if stu == nil{
		return false
	}
	node := &Node{
		Student:stu,
		next:nil,
	}
	if list.size == 0{
		list.first = node
	}else{
		list.last.next = node
	}
	list.last = node
	list.size++
	return true
}
func (list *List)AppendAt(stu *Student,index int) bool{
	if index <= 0{
		oldFirst := list.first
		list.first = &Node{
			Student : stu,
			next: oldFirst,
		}
		list.size++
		return true
	}

	if index >= list.size{
		list.Append(stu)
		return true
	}
	node := &Node{
		Student : stu,
	}
	list.size++
	insertPositionNode := list.GetNode(index)
	insertPositionNodeNext := insertPositionNode.next

	insertPositionNode.next =node
	node.next = insertPositionNodeNext

	return true
}
func (list *List)Remove(index int) bool{
	if index<=0 || index > list.GetSize() {
		return false
	}
	//如果只有一个节点,可直接调用init方法初始化
	if(list.size == 1){
		list.Init()
		return true
	}
	
	//如果删除first
	if index == 1 {
		list.first = list.GetNode(2)
		list.size--
		return true
	}
	//如果删除的是尾部
	if index == list.size {
		list.last = list.GetNode(index - 1)
		list.size--
		return true
	}
	//删除的是中间节点
	preNode := list.GetNode(index - 1)
	preNode.next = preNode.next.next
	list.size--
	return true	
}
func (list *List)PrintList(){
	for i:= 1;i<=list.GetSize();i++{
		fmt.Println(*list.GetNode(i).Student)
	}
	fmt.Println()
}

