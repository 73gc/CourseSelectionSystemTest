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

func (s *Admin) AddCourse() {
	Clear()
	// 测试管理员添加课程接口
	req := server.NewAdminAddCourseInfoRequest()
	fmt.Printf("请输入课程号: ")
	fmt.Scanf("%s", &req.CourseId)
	fmt.Printf("请输入课程名称: ")
	fmt.Scanf("%s", &req.CourseName)
	fmt.Printf("请输入任课教师职工号: ")
	fmt.Scanf("%s", &req.TeacherId)
	fmt.Printf("请输入课程学分: ")
	fmt.Scanf("%g", &req.Credit)
	resp, err := cli.AddCourse(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(resp.Message)
	time.Sleep(3 * time.Second)
}

func (s *Admin) showCourseToAdmin() []*server.CourseInfo {
	Clear()
	// 测试管理员查询课程信息接口
	resp, err := cli.QueryCourseInfo(context.Background())
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	// fmt.Println(resp)
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("序号", "课程号", "课程名", "任课教师", "学分")
	for i, v := range resp.Courses {
		table.AddRow(i+1, v.CourseId, v.CourseName, v.TeacherName, v.Credit)
	}
	fmt.Println(table)
	return resp.Courses
}

func chooseCourse() int {
	fmt.Printf("选择课程序号或退出(0): ")
	var op int
	fmt.Scanf("%d", &op)
	return op
}

func (s *Admin) ModifyCourse() {
	Clear()
	for {
		Courses := s.showCourseToAdmin()
		op := chooseCourse()
		for ; op < 0 || op > len(Courses); op = chooseCourse() {
			fmt.Println("无效操作, 请重新选择")
		}
		if op == 0 {
			return
		}
		// 测试管理员删除课程接口
		req := &server.AdminDeleteCourseInfoRequest{
			CourseId: Courses[op-1].CourseId,
		}
		_, err := cli.DeleteCourse(context.Background(), req)
		if err != nil {
			log.Println(err.Error())
			return
		}
		req1 := server.NewAdminAddCourseInfoRequest()
		fmt.Printf("请输入课程号: ")
		fmt.Scanf("%s", &req1.CourseId)
		fmt.Printf("请输入课程名称: ")
		fmt.Scanf("%s", &req1.CourseName)
		fmt.Printf("请输入任课教师职工号: ")
		fmt.Scanf("%s", &req1.TeacherId)
		fmt.Printf("请输入课程学分: ")
		fmt.Scanf("%g", &req1.Credit)
		resp, err := cli.AddCourse(context.Background(), req1)
		if err != nil {
			log.Println(err.Error())
		}
		if resp.Message == "添加成功" {
			fmt.Println("修改成功")
		}
		time.Sleep(3 * time.Second)
	}
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
	case 4:
		admin.AddCourse()
	case 5:
		admin.ModifyCourse()
	case 0:
		return true
	}
	return false
}

func (s *Teacher) ShowCourses() []*server.ShowCourse2Teacher {
	Clear()
	// 测试教师查看选课信息接口
	req := &server.TeacherQueryCourseRequest{
		TeacherId: Usrid,
	}
	resp, err := cli.ShowCourseSelection(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("序号", "课程号", "课程名", "学分")
	for i, v := range resp.Courses {
		table.AddRow(i+1, v.CourseId, v.CourseName, v.Credit)
	}
	fmt.Println(table)
	return resp.Courses
}

func (s *Teacher) showStudents(courseId string) {
	Clear()
	// 测试教师查看选课学生接口
	req := &server.ShowStudentInfoRequest{
		CourseId: courseId,
	}
	resp, err := cli.StudentCourseSelection(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("序号", "学号", "姓名")
	for i, v := range resp.Students {
		table.AddRow(i+1, v.StudentId, v.StudentName)
	}
	fmt.Println(table)
	fmt.Printf("回车退出")
	var str string
	fmt.Scanf("%s", &str)
	return
}

func (s *Teacher) ShowCourseSelection() {
	Clear()
	for {
		Courses := s.ShowCourses()
		op := chooseCourse()
		for ; op < 0 || op > len(Courses); op = chooseCourse() {
			fmt.Println("无效操作, 请重新选择")
		}
		if op == 0 {
			return
		}
		s.showStudents(Courses[op-1].CourseId)
	}
}

func (s *Teacher) showUninputCourse() []*server.ShowCourse2Teacher {
	Clear()
	// 测试向老师展示未录入成绩课程接口
	req := &server.TeacherQueryCourseRequest{
		TeacherId: Usrid,
	}
	resp, err := cli.ShowCourseToTeacher(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("序号", "课程号", "课程名", "学分")
	for i, v := range resp.Courses {
		table.AddRow(i+1, v.CourseId, v.CourseName, v.Credit)
	}
	fmt.Println(table)
	return resp.Courses
}

func (s *Teacher) showUninputStudent(courseId string) []*server.StudentCourseInfo {
	Clear()
	// 测试向老师展示未录入成绩的学生的信息
	req := &server.ShowStudentInfoRequest{
		CourseId: courseId,
	}
	resp, err := cli.ShowStudentInfo(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("序号", "学号", "姓名")
	for i, v := range resp.Students {
		table.AddRow(i+1, v.StudentId, v.StudentName)
	}
	fmt.Println(table)
	return resp.Students
}

func (s *Teacher) inputScore(courseId, studentId string) {
	// 测试老师录入成绩接口
	req := &server.TeacherInputScoreRequest{
		StudentId: studentId,
		CourseId:  courseId,
	}
	fmt.Printf("请输入成绩: ")
	fmt.Scanf("%g", &req.Score)
	resp, err := cli.InputScore(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(resp.Message)
	time.Sleep(3 * time.Second)
}

func (s *Teacher) InputScore() {
	Clear()
	for {
		Courses := s.showUninputCourse()
		op := chooseCourse()
		for ; op < 0 || op > len(Courses); op = chooseCourse() {
			fmt.Println("无效操作, 请重新选择")
		}
		if op == 0 {
			return
		}
		for {
			Students := s.showUninputStudent(Courses[op-1].CourseId)
			ops := chooseStudent()
			for ; ops < 0 || ops > len(Students); ops = chooseStudent() {
				fmt.Println("无效操作, 请重新选择")
			}
			if op == 0 {
				break
			}
			s.inputScore(Courses[op-1].CourseId, Students[ops-1].StudentId)
		}
	}
	return
}

func (s *Teacher) showInputedCourse() []*server.ShowCourse2Teacher {
	Clear()
	// 测试向老师展示已录入成绩课程接口
	req := &server.TeacherQueryCourseRequest{
		TeacherId: Usrid,
	}
	resp, err := cli.ModifyShowCourse(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	// fmt.Println(resp)
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("序号", "课程号", "课程名", "学分")
	for i, v := range resp.Courses {
		table.AddRow(i+1, v.CourseId, v.CourseName, v.Credit)
	}
	fmt.Println(table)
	return resp.Courses
}

func (s *Teacher) showInputedStudent(courseId string) []*server.StudentCourseInfo {
	Clear()
	// 测试向老师展示已录入成绩的学生的信息
	req := &server.ShowStudentInfoRequest{
		CourseId: courseId,
	}
	resp, err := cli.ModifyShowStudent(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	// fmt.Println(resp)
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("序号", "学号", "姓名", "成绩")
	for i, v := range resp.Students {
		table.AddRow(i+1, v.StudentId, v.StudentName, v.Score)
	}
	fmt.Println(table)
	return resp.Students
}

func (s *Teacher) modifyScore(studentId, courseId string) {
	// 测试老师修改成绩接口
	req := &server.TeacherModifyScoreRequest{
		StudentId: studentId,
		CourseId:  courseId,
	}
	fmt.Printf("请输入成绩: ")
	fmt.Scanf("%g", &req.Score)
	resp, err := cli.ModifyScore(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	fmt.Println(resp.Message)
	time.Sleep(time.Second)
	return
}

func (s *Teacher) ModifyScore() {
	Clear()
	for {
		Courses := s.showInputedCourse()
		op := chooseCourse()
		for ; op < 0 || op > len(Courses); op = chooseCourse() {
			fmt.Println("无效操作, 请重新选择")
		}
		if op == 0 {
			return
		}
		for {
			Students := s.showInputedStudent(Courses[op-1].CourseId)
			ops := chooseStudent()
			for ; ops < 0 || ops > len(Students); ops = chooseStudent() {
				fmt.Println("无效操作, 请重新选择")
			}
			if op == 0 {
				break
			}
			s.modifyScore(Courses[op-1].CourseId, Students[ops-1].StudentId)
		}
	}
}

func (s *Teacher) ShowScore() {
	Clear()
	for {
		Courses := s.showInputedCourse()
		op := chooseCourse()
		for ; op < 0 || op > len(Courses); op = chooseCourse() {
			fmt.Println("无效操作, 请重新选择")
		}
		if op == 0 {
			return
		}
		s.showInputedStudent(Courses[op-1].CourseId)
		fmt.Printf("回车退出")
		var str string
		fmt.Scanf("%s", &str)
	}
}

func (s *UI) TeacherUI() bool {
	Clear()
	fmt.Println("教师用户: ", Usrid)
	fmt.Println("1. 修改密码")
	fmt.Println("2. 查看选课信息")
	fmt.Println("3. 录入成绩")
	fmt.Println("4. 修改成绩")
	fmt.Println("5. 查询成绩")
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
		teacher.ShowCourseSelection()
	case 3:
		teacher.InputScore()
	case 4:
		teacher.ModifyScore()
	case 5:
		teacher.ShowScore()
	case 0:
		return true
	}
	return false
}

func (s *Student) showCourseToStudent() []*server.ShowCourseResponse {
	Clear()
	// 测试向学生展示课程的接口
	// studentId := "01907010109"
	req := &server.StudentShowCourseRequest{
		StudentId: &Usrid,
	}
	resp, err := cli.ShowCourse(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("序号", "课程号", "课程名", "任课教师", "学分")
	for i, v := range resp.Courses {
		table.AddRow(i+1, v.CourseId, v.CourseName, v.TeacherName, v.Credit)
	}
	fmt.Println(table)
	return resp.Courses
}

func (s *Student) SelectCourse() {
	Clear()
	for {
		Courses := s.showCourseToStudent()
		op := chooseCourse()
		for ; op < 0 || op > len(Courses); op = chooseCourse() {
			fmt.Println("无效操作, 请重新选择")
		}
		if op == 0 {
			return
		}
		// 测试学生选课接口
		req := &server.SelectCourseRequest{
			StudentId: Usrid,
			CourseId:  Courses[op-1].CourseId,
		}
		resp, err := cli.SelectCourse(context.Background(), req)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(resp.Message)
		time.Sleep(3 * time.Second)
	}
}

func (s *Student) QuerySelection() {
	Clear()
	// 测试学生查询选课信息接口
	req := &server.StudentQuerySelectionRequest{
		StudentId: Usrid,
	}
	resp, err := cli.QuerySelection(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// fmt.Println(resp)
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("序号", "课程名", "任课教师", "学分")
	for i, v := range resp.Courses {
		table.AddRow(i+1, v.CourseId, v.CourseName, v.TeacherName, v.Credit)
	}
	fmt.Println(table)
	fmt.Printf("回车退出")
	var str string
	fmt.Scanf("%s", &str)
	return
}

func (s *Student) QueryScore() {
	Clear()
	// 测试学生查询成绩接口
	req := &server.StudentQueryScoreRequest{
		StudentId: Usrid,
	}
	resp, err := cli.QueryScore(context.Background(), req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	// fmt.Println(resp)
	table := uitable.New()
	table.MaxColWidth = 50
	table.AddRow("序号", "课程号", "课程名", "学分", "成绩")
	for i, v := range resp.CourseScore {
		table.AddRow(i, v.CourseId, v.CourseName, v.Credit, v.Score)
	}
	fmt.Println(table)
	fmt.Printf("点击回车退出")
	var str string
	fmt.Scanf("%s", &str)
	return
}

func (s *Student) chooseTeacher() int {
	fmt.Printf("选择教师或退出(0): ")
	var op int
	fmt.Scanf("%d", &op)
	return op
}

func (s *Student) Evaluate() {
	Clear()
	for {
		Courses := s.showCourseToStudent()
		op := s.chooseTeacher()
		for ; op < 0 || op > len(Courses); op = s.chooseTeacher() {
			fmt.Println("无效操作, 请重新选择")
		}
		if op == 0 {
			return
		}
		// 测试学生评教接口
		req := &server.StudentEvaluateRequest{
			StudentId: Usrid,
			CourseId:  Courses[op-1].CourseId,
			Score:     100,
		}
		resp, err := cli.EvaluateRequest(context.Background(), req)
		if err != nil {
			log.Println(err.Error())
			return
		}
		fmt.Println(resp.Message)
		time.Sleep(3 * time.Second)
	}
	return
}

func (s *UI) StudentUI() bool {
	Clear()
	fmt.Println("学生用户: ", Usrid)
	fmt.Println("1. 修改密码")
	fmt.Println("2. 选课")
	fmt.Println("3. 查看选课信息")
	fmt.Println("4. 查询成绩")
	fmt.Println("5. 评教")
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
		student.SelectCourse()
	case 3:
		student.QuerySelection()
	case 4:
		student.QueryScore()
	case 5:
		student.Evaluate()
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
	return
}
