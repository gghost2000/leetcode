type MyQueue struct {
	a []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	this.a = append(this.a, x)
}

func (this *MyQueue) Pop() int {
	val := this.a[0]
	this.a = this.a[1:]
	return val
}

func (this *MyQueue) Peek() int {
	return this.a[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.a) == 0
}



/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */