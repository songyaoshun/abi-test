package main

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

func main() {
	// 定义结构体
	type T struct {
		X *big.Int
		Y *big.Int
	}

	// 创建运行时结构体
	var fields []reflect.StructField
	fields = append(fields, reflect.StructField{
		Name: "X",
		Type: reflect.TypeOf(new(big.Int)),
		Tag:  reflect.StructTag(`json:"x"`),
	})
	fields = append(fields, reflect.StructField{
		Name: "Y",
		Type: reflect.TypeOf(new(big.Int)),
		Tag:  reflect.StructTag(`json:"y"`),
	})

	// 构造运行时实例并赋值
	val := reflect.New(reflect.StructOf(fields))
	val.Elem().Field(0).Set(reflect.ValueOf(big.NewInt(1)))
	val.Elem().Field(1).Set(reflect.ValueOf(big.NewInt(2)))

	// 使用ConvertType进行类型转换
	out := abi.ConvertType(val.Interface(), new(T)).(*T)
	fmt.Printf("X: %v, Y: %v\n", out.X, out.Y)
}
