package stream_test

import (
	"reflect"

	"github.com/Pallinder/go-randomdata"
	"github.com/hrz123/stream"
	"github.com/hrz123/stream/types"
)

// - Q1: 输入 employees，返回 年龄 >22岁 的所有员工，年龄总和
func Question1Sub1(employees []*Employee) int64 {
	return int64(stream.OfSlice(employees).Filter(func(e types.T) bool {
		return *(e.(*Employee).Age) > 22
	}).ReduceWith(0, func(acc types.R, e types.T) types.R {
		return acc.(int) + *(e.(*Employee).Age)
	}).(int))
}

// - Q2: - 输入 employees，返回 id 最小的十个员工，按 id 升序排序
func Question1Sub2(employees []*Employee) []*Employee {
	res := stream.OfSlice(employees).Sorted(func(left types.T, right types.T) int {
		e1, e2 := left.(*Employee), right.(*Employee)
		return int(e1.Id - e2.Id)
	}).Limit(10).ToSliceOf(reflect.TypeOf(&Employee{}))
	return res.([]*Employee)
}

// - Q3: - 输入 employees，对于没有手机号为0的数据，随机填写一个
func Question1Sub3(employees []*Employee) []*Employee {
	return stream.OfSlice(employees).ReduceWith(make([]*Employee, 0, len(employees)), func(acc types.R, e types.T) types.R {
		employee := new(Employee)
		if em := e.(*Employee); *(em.Phone) == "" {
			employee.Id = em.Id
			employee.Name = em.Name
			employee.Age = em.Age
			employee.Position = em.Position
			phone := randomdata.PhoneNumber()
			employee.Phone = &phone
			acc = append(acc.([]*Employee), employee)
		} else {
			acc = append(acc.([]*Employee), em)
		}
		return acc
	}).([]*Employee)
}

// - Q4: - 输入 employees ，返回一个map[int][]int，其中 key 为 员工年龄 Age，value 为该年龄段员工ID
func Question1Sub4(employees []*Employee) map[int][]int64 {
	return stream.OfSlice(employees).ReduceWith(map[int][]int64{}, func(acc types.R, e types.T) types.R {
		employee := e.(*Employee)
		m := acc.(map[int][]int64)
		m[*employee.Age] = append(m[*employee.Age], employee.Id)
		return m
	}).(map[int][]int64)
}
