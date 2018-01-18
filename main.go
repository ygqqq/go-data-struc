package main
import(
	"./list"
	"strconv"
)

func main() {
	var l = list.List{}
	l.Init();

	for i:=0;i<2;i++{
		stu := &list.Student{
			Name:"ygq"+strconv.Itoa(i),
			Age:18+i,
		}
		l.Append(stu)
	}
	//l.PrintList()
	stu := &list.Student{
		Name : "haha",
		Age : 88,
	}
	l.AppendAt(stu,0)
	l.AppendAt(stu,2)
	l.Remove(l.GetSize())
	l.PrintList()
}

