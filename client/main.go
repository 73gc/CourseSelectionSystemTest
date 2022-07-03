package main

import (
	"courseselection/kitex_gen/Server/service"
	"log"

	"github.com/cloudwego/kitex/client"
)

func main() {
	cli, err := service.NewClient("course.selection", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Println(err.Error())
		return
	}
	// 登录接口测试
	// req := &server.LoginRequest{
	// 	Username: "22070303001",
	// 	Password: "22070303001",
	// }
	// resp, err := cli.Login(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp.Message)
	// fmt.Println(*resp.Authority)

	// 修改密码接口测试
	// req := &server.ChangePasswordRequenst{
	// 	Username:     "22070301001",
	// 	NewPassword_: "acmicpc",
	// }
	// resp, err := cli.ChangePassword(context.Background(), req)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	return
	// }
	// fmt.Println(resp.Message)

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
}
