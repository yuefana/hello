package structPack

type ExpStruct struct {
	Mi1 int
	Mi2 float32
}

type node struct {
	data int    `info:"data"`
	next *node  `info:"next pointer"`
	Name string `info:"name"`
}

func NewNode(num int) *node {
	// cur := new(node)
	// cur.data = num
	// cur.name = "hello"
	// return cur
	return &node{data: num, Name: "why"}
}
