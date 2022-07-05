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
	"github.com/gosuri/uitable"
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
var Usrid string

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

func (s *UI) LoginUI() {
	Clear()
	var Username, Password string
	fmt.Printf("用户名: ")
	fmt.Scanf("%s", &Username)
	Usrid = Username
	passwd, _ := gopass.GetPasswdPrompt("密码: ", true, os.Stdin, os.Stdout)
	Password = string(passwd)
	var Message string
	Message = user.Login(Username, Password)
	fmt.Println(Message)
	if Message == "登录成功" {
		time.Sleep(5 * time.Second)
		switch Authority {
		case 1:
			for {
				ex := s.AdminUI()
				if ex {
					return
				}
			}
		case 2:
			for {
				ex := s.TeacherUI()
				if ex {
					return
				}
			}
		case 3:
			for {
				ex := s.StudentUI()
				if ex {
					return
				}
			}
		}
	} else {
		time.Sleep(5 * time.Second)
		s.LoginUI()
	}
	return
}

func (s *UI) ChangePasswdUI() {
	Clear()
	oldPasswd, _ := gopass.GetPasswdPrompt("原密码: ", true, os.Stdin, os.Stdout)
	newPasswd, _ := gopass.GetPasswdPrompt("新密码: ", true, os.Stdin, os.Stdout)
	ReNewPasswd, _ := gopass.GetPasswdPrompt("确认新密码", true, os.Stdin, os.Stdout)
	if string(newPasswd) != string(ReNewPasswd) {
		fmt.Println("两次输入密码不一致")
		time.Sleep(5 * time.Second)
		return
	}
	resp := user.Login(Usrid, string(oldPasswd))
	if resp != "登录成功" {
		fmt.Println("原密码错误")
		time.Sleep(5 * time.Second)
		return
	}
	Message := user.ChangePassword(Usrid, string(newPasswd))
	if Message != "修改成功" {
		fmt.Println(Message)
		return
	}
	fmt.Println("修改成功")
	time.Sleep(5 * time.Second)
	return
}

func ChooseOp() int {
	fmt.Printf("选择你要进行的操作: ")
	var op int
	fmt.Scanf("%d", &op)
	return op
}

func (s *Admin) addTeacher() {
	Clear()
	// 测试管理员添加教师接口
	var teacherId, teacherName string
	fmt.Println("输入职工号: ")
	fmt.Scanf("%s", &teacherId)
	fmt.Println("输入教师姓名: ")
	fmt.Scanf("%s", &teacherName)
	req := &server.AdminAddTeacherInfoRequest{
		TeacherId:   teacherId,
		TeacherName: teacherName,
	}
	resp, err := cli.AddTeacher(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(resp.Message)
}

func (s *Admin) addStudent() {
	Clear()
	// 测试管理员添加学生接口
	var studentId, studentName, studentClass string
	fmt.Println("输入学号: ")
	fmt.Scanf("%s", &studentId)
	fmt.Println("输入学生姓名: ")
	fmt.Scanf("%s", &studentName)
	fmt.Println("输入班级: ")
	fmt.Scanf("%s", &studentClass)
	req := &server.AdminAddStudentInfoRequest{
		StudentId:     studentId,
		StudentName:   studentName,
		ClassAndGrade: studentClass,
	}
	resp, err := cli.AddStudent(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(resp.Message)
}

func getIdentity() int {
	var identity int
	fmt.Printf("用户身份(1教师/2学生): ")
	fmt.Scanf("%d", &identity)
	return identity
}

func (s *Admin) AddUser() {
	Clear()
	identity := getIdentity()
	for ; identity < 1 || identity > 2; identity = getIdentity() {
		fmt.Println("无效身份, 请重新输入")
	}
	if identity == 1 {
		s.addTeacher()
	} else {
		s.addStudent()
	}
	time.Sleep(3 * time.Second)
}

func chooseTeacher() int {
	fmt.Printf("选择教师序号: ")
	var op int
	fmt.Scanf("%d", &op)
	return op
}

func (s *Admin) deleteTeacher() {
	Clear()
	// 测试管理员查询教师信息接口
	resp1, err := cli.QueryTeacherInfo(context.Background())
	if err != nil {
		log.Println(err.Error())
		return
	}
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("", "教师职工号", "教师姓名")
	for i, v := range resp1.Teachers {
		table.AddRow(i+1, v.TeacherId, v.TeacherName)
	}
	fmt.Println(table)
	tCount := len(resp1.Teachers)
	op := chooseTeacher()
	for ; op < 1 || op > tCount; op = chooseTeacher() {
		fmt.Println("教师不存在, 请重新选择")
	}
	// 测试管理员删除教师接口
	req := &server.AdminDeleteTeacherInfoRequest{
		TeacherId: resp1.Teachers[op-1].TeacherId,
	}
	resp, err := cli.DeleteTeacher(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(resp.Message)
}

func chooseStudent() int {
	var op int
	fmt.Printf("选择学生的序号: ")
	fmt.Scanf("%d", &op)
	return op
}

func (s *Admin) deleteStudent() {
	Clear()
	// 测试管理员查询学生信息接口
	resp1, err := cli.QueryStudentInfo(context.Background())
	if err != nil {
		log.Println(err.Error())
		return
	}
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("", "学号", "姓名", "班级")
	for i, v := range resp1.Students {
		table.AddRow(i+1, v.StudentId, v.StudentName, v.ClassAndGrade)
	}
	fmt.Println(table)
	sCount := len(resp1.Students)
	op := chooseStudent()
	for ; op < 1 || op > sCount; op = chooseStudent() {
		fmt.Println("学生不存在, 请重新选择")
	}
	// 测试管理员删除学生接口
	req := &server.AdminDeleteStudentInfoRequest{
		StudentId: resp1.Students[op-1].StudentId,
	}
	resp, err := cli.DeleteStudent(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(resp.Message)
}

func (s *Admin) DeleteUser() {
	Clear()
	identity := getIdentity()
	for ; identity < 1 || identity > 2; identity = getIdentity() {
		fmt.Println("无效身份, 请重新输入")
	}
	if identity == 1 {
		s.deleteTeacher()
	} else {
		s.deleteStudent()
	}
	time.Sleep(3 * time.Second)
}

func (s *UI) AdminUI() bool {
	Clear()
	fmt.Println("管理员用户: ", Usrid)
	fmt.Println("1. 修改密码")
	fmt.Println("2. 添加用户")
	fmt.Println("3. 删除用户")
	fmt.Println("4. 添加课程")
	fmt.Println("5. 修改课程")
	fmt.Println("0. 退出")
	op := ChooseOp()
	for ; op < 0 || op > 5; op = ChooseOp() {
		fmt.Println("无效操作，请重新选择")
	}
	switch op {
	case 1:
		ui.ChangePasswdUI()
		return true
	case 2:
		admin.AddUser()
	case 3:
		admin.DeleteUser()
	case 0:
		return true
	}
	return false
}

func (s *UI) TeacherUI() bool {
	Clear()
	fmt.Println("教师用户: ", Usrid)
	fmt.Println("1. 修改密码")
	fmt.Println("0. 退出")
	op := ChooseOp()
	for ; op < 0 || op > 1; op = ChooseOp() {
		fmt.Println("无效操作，请重新选择")
	}
	switch op {
	case 1:
		ui.ChangePasswdUI()
		return true
	case 0:
		return true
	}
	return false
}

func (s *UI) StudentUI() bool {
	Clear()
	fmt.Println("学生用户: ", Usrid)
	fmt.Println("1. 修改密码")
	fmt.Println("0. 退出")
	op := ChooseOp()
	for ; op < 0 || op > 1; op = ChooseOp() {
		fmt.Println("无效操作，请重新选择")
	}
	switch op {
	case 1:
		ui.ChangePasswdUI()
		return true
	case 0:
		return true
	}
	return false
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

	// 测试管理员查询课程信息接口
	// resp, err := cli.QueryCourseInfo(context.Background())
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp)

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
