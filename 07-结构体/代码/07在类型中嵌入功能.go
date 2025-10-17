package main
/*主要有两种方法来实现在类型中嵌入功能：

A：聚合（或组合）：包含一个所需功能类型的具名字段。

B：内嵌：内嵌（匿名地）所需功能类型，像前一节 10.6.5 所演示的那样。

为了使这些概念具体化，假设有一个 Customer 类型，我们想让它通过 Log 类型来包含日志功能，Log 类型只是简单地包含一个累积的消息（当然它可以是复杂的）。如果想让特定类型都具备日志功能，你可以实现一个这样的 Log 类型，然后将它作为特定类型的一个字段，并提供 Log()，它返回这个日志的引用。

*/
import "fmt"

type Log struct {
	msg string
}

type Custom struct {
	Name string
	log *Log
}

func main() {
	// c := new(Custom)
	// c.Name = "Tom"
	// c.log = new(Log)
	// c.log.msg = "1 - Yes we can!"

	fmt.Println("-----------")
	c := &Custom{"Barak Obama", &Log{"1 - Yes we can!"}}
	c.log.Add("2 - Yes we can111")
	fmt.Println(c.Log())
}

func (l *Log) Add(s string){
	l.msg += "\n" + s
}

func (c *Custom) Log() string{
	return c.log.msg
}


