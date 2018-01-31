package basictype

//1. boolean
var a1 bool = true
var a2 bool = false

//2. number
var b1 uint8 = 20
var b2 uint16 = 3000
var b3 uint32
var b4 uint64

var b5 int8 = -2
var b6 int16 = -3000
var b7 int32
var b8 int64

var b9 float32
var b10 float64

var b11 byte // =uint8
var b12 rune // =int32

var b13 uint // = 32/64
var b14 int  // =uint

//查看最大值
var MAX_UINT = ^uint(0)
var MAX_INT = int(^uint(0) >> 1) //首位0，其余为1
var MIN_INT = ^MAX_INT           //首位为1，其余为0

//3. string
var c string = "youyou"

//4. array
var d []int

//5. 切片类型
make([]int, 50, 100)
new([100]int)[0:50]

//6. struct 结构体
type Student struct {
	id		int
	name	string 
	address string
}

//7. 指针类型
var s *Student = new(Student)
var s1 *Student = &Student{102, "jun", "hangzhou"}

//8. 函数类型
func test() *Student{
	return &Student{102, "jun", "hangzhou"}
}

//9. 接口类型

type Reader interface {
	Read(p []byte) (n int, err os.Error)
}

//10. Map

var m1 map[string]string









