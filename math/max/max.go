package max

import (
	"github.com/XIAOZHUXUEJAVA/Go-Algorithms-DataStructures-Journey/constraints"
)

// constraints.Integer 表示该泛型类型参数 T 必须满足 Integer 约束
// 即它必须是一个整数类型，可以是有符号整数类型或无符号整数类型。
func Int[T constraints.Integer](values ...T) T {
	max := values[0]
	for _, value := range values {
		if value > max {
			max = value
		}
	}
	return max
}
