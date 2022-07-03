namespace go Server

struct LoginRequest {
    1: required string username
    2: required string password
}

struct LoginResponse {
    1: required string Message
    2: optional i32 authority
}

struct ChangePasswordRequenst {
    1: required string username
    2: required string NewPassword
}

struct ChangePasswordReponse {
    1: required string Message
}

struct SelectCourseRequest {
    1: required string CourseId
    2: required string StudentId
}

struct SelectCourseResponse {
    1: required string Message
}

struct StudentQueryScoreRequest {
    1: required string StudentId
}

struct StudentScoreResponse {
    1: string CourseId
    2: string CourseName
    3: double Credit
    4: double Score
}

struct StudentQueryScoreResponse {
    1: required list<StudentScoreResponse> CourseScore
}

struct StudentShowCourseRequest {
    1: optional string StudentId
}

struct ShowCourseResponse {
    1: string CourseId
    2: string CourseName
    3: string TeacherName
    4: double Credit
}

struct StudentShowCourseReponse {
    1: required list<ShowCourseResponse> Courses
}

struct StudentQuerySelectionRequest {
    1: required string StudentId
}

struct StudentQuerySelectionResponse {
    1: required list<ShowCourseResponse> Courses
}

struct StudentEvaluateRequest {
    1: required string StudentId
    2: required string CourseId
}

struct StudentEvaluateResponse {
    1: required string Message
}

struct ShowCourse2Teacher {
    1: string CourseId
    2: string CourseName
    3: string Credit
}

struct TeacherQueryCourseRequest {
    1: required string TeacherId
}

struct TeacherQueryCourseResponse {
    1: required list<ShowCourse2Teacher> Courses
}

struct StudentCourseInfo {
    1: string StudentId
    2: string StudentName
    3: double Score
}

struct ShowStudentInfoRequest {
    1: required string CourseId
}

struct ShowStudentInfoResponse {
    1: required list<StudentCourseInfo> Students
}

struct TeacherInputScoreRequest {
    1: required string StudentId
    2: required string CourseId
    3: required double Score
}

struct TeacherInputScoreResponse {
    1: required string Message
}

struct TeacherModifyScoreRequest {
    1: required string StudentId
    2: required string CourseId
    3: required double Score
}

struct TeacherModifyScoreResponse {
    1: required string Message
}

struct StudentInfo {
    1: string StudentId
    2: string StudentName
    3: string ClassAndGrade
}

struct TeacherInfo {
    1: string TeacherId
    2: string TeacherName
}

struct CourseInfo {
    1: string CourseId
    2: string CourseName
    3: string TeacherName
    4: string Credit
}

struct AdminQueryStudentInfoResponse {
    1: required list<StudentInfo> Students
}

struct AdminQueryTeacherInfoResponse {
    1: required list<TeacherInfo> Teachers
}

struct AdminQueryCourseInfoResponse {
    1: required list<CourseInfo> Courses
}

struct AdminAddStudentInfoRequest {
    1: required string StudentId
    2: required string StudentName
    3: required string ClassAndGrade
}

struct AdminAddStudentInfoResponse {
    1: required string Message
}

struct AdminDeleteStudentInfoRequest {
    1: required string StudentId
}

struct AdminDeleteStudentInfoResponse {
    1: required string Message
}

struct AdminAddTeacherInfoRequest {
    1: required string TeacherId
    2: required string TeacherName
}

struct AdminAddTeacherInfoResponse {
    1: required string Message
}

struct AdminDeleteTeacherInfoRequest {
    1: required string TeacherId
}

struct AdminDeleteTeacherInfoResponse {
    1: required string Message
}

struct AdminAddCourseInfoRequest {
    1: required string CourseId
    2: required string CourseName
    3: required string TeacherId
    4: required string Credit
}

struct AdminAddCourseInfoResponse {
    1: required string Message
}

struct AdminDeleteCourseInfoRequest {
    1: required string CourseId
}

struct AdminDeleteCourseInfoResponse {
    1: required string Message
}

service Service {
    LoginResponse Login(1: LoginRequest req) // 登录
    ChangePasswordReponse ChangePassword(1: ChangePasswordRequenst req) // 退出

    StudentShowCourseReponse ShowCourse(1: StudentShowCourseRequest req) // 展示课程信息，供选课
    SelectCourseResponse SelectCourse(1: SelectCourseRequest req) // 选课
    StudentQueryScoreResponse QueryScore(1: StudentQueryScoreRequest req) // 查询课程成绩
    StudentQuerySelectionResponse QuerySelection(1: StudentQuerySelectionRequest req) // 查看选课
    StudentEvaluateResponse EvaluateRequest(1: StudentEvaluateRequest req) // 评教

    TeacherQueryCourseResponse ShowCourseToTeacher(1: TeacherQueryCourseRequest req) // 显示课程
    ShowStudentInfoResponse ShowStudentInfo(1: ShowStudentInfoRequest req) // 显示学生
    TeacherInputScoreResponse InputScore(1: TeacherInputScoreRequest req) // 输入成绩
    TeacherQueryCourseResponse ModifyShowCourse(1: TeacherQueryCourseRequest req)
    ShowStudentInfoResponse ModifyShowStudent(1: ShowStudentInfoRequest req)
    TeacherModifyScoreResponse ModifyScore(1: TeacherModifyScoreRequest req) // 修改成绩
    TeacherQueryCourseResponse QueryCourse(1: TeacherQueryCourseRequest req)
    ShowStudentInfoResponse ShowStudentScore(1: ShowStudentInfoRequest req) // 查看学生成绩
    TeacherQueryCourseResponse ShowCourseSelection(1: TeacherQueryCourseRequest req)
    ShowStudentInfoResponse StudentCourseSelection(1: ShowStudentInfoRequest req) // 查看选课学生

    AdminQueryStudentInfoResponse QueryStudentInfo() // 显示学生信息
    AdminQueryTeacherInfoResponse QueryTeacherInfo() // 显示教师信息
    AdminQueryCourseInfoResponse QueryCourseInfo() // 显示课程信息
    AdminAddStudentInfoResponse AddStudent(1: AdminAddStudentInfoRequest req) // 添加学生
    AdminDeleteStudentInfoResponse DeleteStudent(1: AdminDeleteStudentInfoRequest req) // 删除学生
    AdminAddTeacherInfoResponse AddTeacher(1: AdminAddTeacherInfoRequest req) // 添加教师
    AdminDeleteTeacherInfoResponse DeleteTeacher(1: AdminDeleteTeacherInfoRequest req) // 删除教师
    AdminAddCourseInfoResponse AddCourse(1: AdminAddCourseInfoRequest req) // 添加教程
    AdminDeleteCourseInfoResponse DeleteCourse(1: AdminDeleteCourseInfoRequest req) // 删除课程
}