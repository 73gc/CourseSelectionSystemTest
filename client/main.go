package main

import (
	"context"
	server "courseselection/kitex_gen/Server"
	"courseselection/kitex_gen/Server/service"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/howeyc/gopass"
)

var cli service.Client
var err error

type User struct{}
type Admin struct{}
type Teacher struct{}
type Student struct{}
type UI struct{}

var ui *UI
var user *User
var admin *Admin
var teacher *Teacher
var student *Student

var Authority int32

func Clear() {
	switch runtime.GOOS {
	case "linux":
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func (s *User) Login(Username, Password string) string {
	// 登录接口测试
	req := &server.LoginRequest{
		Username: Username,
		Password: Password,
	}
	resp, err := cli.Login(context.Background(), req)
	if resp.Authority == nil {
		Authority = -1
	} else {
		Authority = *resp.Authority
	}
	if err != nil {
		log.Println(err.Error())
		return resp.Message
	}
	return resp.Message
}

func (s *User) ChangePassword(Username, NewPassword string) string {
	// 修改密码接口测试
	req := &server.ChangePasswordRequenst{
		Username:     Username,
		NewPassword_: NewPassword,
	}
	resp, err := cli.ChangePassword(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return resp.Message
	}
	return resp.Message
}

// func getPassword() string {
// 	var Password string
// 	fmt.Scanf("%s", &Password)
// 	for i := 0; i < len(Password); i++ {
// 		fmt.Printf("\b")
// 	}
// 	for i := 0; i < len(Password); i++ {
// 		fmt.Printf("*")
// 	}
// 	return Password
// }

func (s *UI) LoginUI() {
	Clear()
	var Username, Password string
	fmt.Printf("用户名: ")
	fmt.Scanf("%s", &Username)
	passwd, _ := gopass.GetPasswdPrompt("密码: ", true, os.Stdin, os.Stdout)
	Password = string(passwd)
	var Message string
	Message = user.Login(Username, Password)
	fmt.Println(Message)
	if Message == "登陆成功" {
		time.Sleep(5 * time.Second)
		return
	} else {
		time.Sleep(5 * time.Second)
		s.LoginUI()
		return
	}
}

func main() {
	cli, err = service.NewClient("course.selection", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	ui = &UI{}
	user = &User{}
	admin = &Admin{}
	teacher = &Teacher{}
	student = &Student{}
	ui.LoginUI()

	// 测试管理员查询学生信息接口
	// resp, err := cli.QueryStudentInfo(context.Background())
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试管理员查询教师信息接口
	// resp, err := cli.QueryTeacherInfo(context.Background())
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试管理员查询课程信息接口
	// resp, err := cli.QueryCourseInfo(context.Background())
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试管理员添加学生接口
	// req := &server.AdminAddStudentInfoRequest{
	// 	StudentId:     "22070303001",
	// 	StudentName:   "张同学",
	// 	ClassAndGrade: "1902",
	// }
	// resp, err := cli.AddStudent(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp.Message)

	// 测试管理员删除学生接口
	// req := &server.AdminDeleteStudentInfoRequest{
	// 	StudentId: "22070303001",
	// }
	// resp, err := cli.DeleteStudent(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp.Message)

	// 测试管理员添加教师接口
	// req := &server.AdminAddTeacherInfoRequest{
	// 	TeacherId:   "22070302002",
	// 	TeacherName: "教师2",
	// }
	// resp, err := cli.AddTeacher(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp.Message)

	// 测试管理员删除教师接口
	// req := &server.AdminDeleteTeacherInfoRequest{
	// 	TeacherId: "22070302002",
	// }
	// resp, err := cli.DeleteTeacher(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp.Message)

	// 测试管理员添加课程接口
	// req := &server.AdminAddCourseInfoRequest{
	// 	CourseId:   "22070304001",
	// 	CourseName: "编译原理",
	// 	TeacherId:  "22070302001",
	// 	Credit:     3,
	// }
	// resp, err := cli.AddCourse(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp.Message)

	// 测试管理员删除课程接口
	// req := &server.AdminDeleteCourseInfoRequest{
	// 	CourseId: "22070304001",
	// }
	// resp, err := cli.DeleteCourse(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp.Message)

	// 测试向学生展示课程的接口
	// studentId := "01907010109"
	// req := &server.StudentShowCourseRequest{
	// 	StudentId: &studentId,
	// }
	// resp, err := cli.ShowCourse(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试学生选课接口
	// req := &server.SelectCourseRequest{
	// 	StudentId: "01907010109",
	// 	CourseId:  "22070304001",
	// }
	// resp, err := cli.SelectCourse(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// fmt.Println(resp.Message)

	// 测试学生查询成绩接口
	// req := &server.StudentQueryScoreRequest{
	// 	StudentId: "01907010109",
	// }
	// resp, err := cli.QueryScore(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试学生查询选课信息接口
	// req := &server.StudentQuerySelectionRequest{
	// 	StudentId: "01907010109",
	// }
	// resp, err := cli.QuerySelection(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试学生评教接口
	// req := &server.StudentEvaluateRequest{
	// 	StudentId: "01907010109",
	// 	CourseId:  "22070304001",
	// 	Score:     100,
	// }
	// resp, err := cli.EvaluateRequest(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp.Message)

	// 测试向老师展示未录入成绩课程接口
	// req := &server.TeacherQueryCourseRequest{
	// 	TeacherId: "22070302001",
	// }
	// resp, err := cli.ShowCourseToTeacher(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试向老师展示未录入成绩的学生的信息
	// req := &server.ShowStudentInfoRequest{
	// 	CourseId: "22070304001",
	// }
	// resp, err := cli.ShowStudentInfo(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试老师录入成绩接口
	// req := &server.TeacherInputScoreRequest{
	// 	StudentId: "01907010109",
	// 	CourseId:  "22070304001",
	// 	Score:     90,
	// }
	// resp, err := cli.InputScore(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp.Message)

	// 测试向老师展示已录入成绩课程接口
	// req := &server.TeacherQueryCourseRequest{
	// 	TeacherId: "22070302001",
	// }
	// resp, err := cli.ModifyShowCourse(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试向老师展示已录入成绩的学生的信息
	// req := &server.ShowStudentInfoRequest{
	// 	CourseId: "22070304001",
	// }
	// resp, err := cli.ModifyShowStudent(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试老师修改成绩接口
	// req := &server.TeacherModifyScoreRequest{
	// 	StudentId: "01907010109",
	// 	CourseId:  "22070304001",
	// 	Score:     80,
	// }
	// resp, err := cli.ModifyScore(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp.Message)

	// 测试教师查看选课信息接口
	// req := &server.TeacherQueryCourseRequest{
	// 	TeacherId: "22070302001",
	// }
	// resp, err := cli.ShowCourseSelection(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

	// 测试教师查看选课学生接口
	// req := &server.ShowStudentInfoRequest{
	// 	CourseId: "22070304001",
	// }
	// resp, err := cli.StudentCourseSelection(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)
}
