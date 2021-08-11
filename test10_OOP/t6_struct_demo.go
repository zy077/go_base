package main

import (
    "encoding/json"
    "fmt"
)

/*
使用“面向对象”的思维方式编写一个学生信息管理系统。
学生有id、姓名、年龄、分数等信息
程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
 */

type Student struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Age uint `json:"age"`
    Score uint `json:"score"`
}

// 枚举学生列表
func (s *Student) ListStudents() []*Student {
    Students := make([]*Student, 10, 10)
    for i:=0; i<10; i++ {
        s := &Student{
            ID: i,
            Name: "haha",
            Age: uint(i*6),
            Score: uint(i*10),
        }
        Students[i] = s
    }
    return Students
}

// 添加学生
func (s *Student) AddStudent(data string) *Student {
    // 反序列化
    stu := &Student{}
    err := json.Unmarshal([]byte(data), stu)
    if err != nil {
        fmt.Printf("Error: %s", err)
    }
    fmt.Printf("保存数据：%s \n", stu)
    return stu
}

// 编辑学生
func (s *Student) EditStudent(data string, id int) *Student {
    stuData := &Student{}
    // 反序列化json数据
    err := json.Unmarshal([]byte(data), &stuData)
    if err != nil {
        fmt.Printf("Error: %s", err)
    }
    // 找到要修改的学生
    stu := &Student{
        ID: id,
    }
    // 更新
    stu.Name = stuData.Name
    stu.Age = stuData.Age
    stu.Score = stuData.Score
    // 保存数据
    fmt.Printf("保存学生数据：%s \n", stu)
    return stu
}

// 删除学生
func (s *Student) DeleteStudent(id int) *Student {
    // 找到要删除的学生对象
    stu := &Student{
        ID: id,
        Name: "张三",
        Age: uint(18),
        Score: uint(80),
    }
    fmt.Printf("删除学生数据：%s \n", stu)
    return stu
}

func main(){
    var stu Student
    // 添加
    data := `{"id": 10, "name": "ZhangSan", "age": 20, "score": 80}`
    s01 := stu.AddStudent(data)
    fmt.Printf("ID：%d \n", s01.ID)
    fmt.Printf("Name：%s \n", s01.Name)
    fmt.Printf("Age：%d \n", s01.Age)
    fmt.Printf("Scroe：%d \n", s01.Score)
    // 删除
    s02 := stu.DeleteStudent(20)
    fmt.Println("删除学生对象：", s02)
    // 修改
    updateData := `{"name": "LiSi", "age": 10, "score": 90}`
    s03 := stu.EditStudent(updateData, 20)
    fmt.Println("修改后的学生对象：", s03)
    // 枚举学生
    var s04 []*Student
    s04 = stu.ListStudents()
    for _, s := range s04 {
        fmt.Println("学生：", s)
    }

}