package fixtures

import (
	"fmt"
)

const AA = "你好 aa" // MATCH /existence of non-English characters/
const BB = "aa good" // 你好 aa
// 你好 aa

/*
你好 aa
 */

var c = map[string]string{"good": "你好"} // MATCH /existence of non-English characters/

func Foo() {
	var b = "aa good你好" // MATCH /existence of non-English characters/
	c := "aa good 你好"   // MATCH /existence of non-English characters/
	fmt.Sprint(b)
	fmt.Sprint(c)
	utils.Tr("检查链路隔离成功")
	utils.Tr("good")
}
