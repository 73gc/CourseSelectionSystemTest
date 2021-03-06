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
    3: required double Score
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
    4: required double Credit
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
    LoginResponse Login(1: LoginRequest req) // ??????
    ChangePasswordReponse ChangePassword(1: ChangePasswordRequenst req) // ??????

    StudentShowCourseReponse ShowCourse(1: StudentShowCourseRequest req) // ??????????????????????????????
    SelectCourseResponse SelectCourse(1: SelectCourseRequest req) // ??????
    StudentQueryScoreResponse QueryScore(1: StudentQueryScoreRequest req) // ??????????????????
    StudentQuerySelectionResponse QuerySelection(1: StudentQuerySelectionRequest req) // ????????????
    StudentEvaluateResponse EvaluateRequest(1: StudentEvaluateRequest req) // ??????

    TeacherQueryCourseResponse ShowCourseToTeacher(1: TeacherQueryCourseRequest req) // ????????????
    ShowStudentInfoResponse ShowStudentInfo(1: ShowStudentInfoRequest req) // ????????????
    TeacherInputScoreResponse InputScore(1: TeacherInputScoreRequest req) // ????????????
    TeacherQueryCourseResponse ModifyShowCourse(1: TeacherQueryCourseRequest req)
    ShowStudentInfoResponse ModifyShowStudent(1: ShowStudentInfoRequest req)
    TeacherModifyScoreResponse ModifyScore(1: TeacherModifyScoreRequest req) // ????????????
    TeacherQueryCourseResponse QueryCourse(1: TeacherQueryCourseRequest req)
    ShowStudentInfoResponse ShowStudentScore(1: ShowStudentInfoRequest req) // ??????????????????
    TeacherQueryCourseResponse ShowCourseSelection(1: TeacherQueryCourseRequest req)
    ShowStudentInfoResponse StudentCourseSelection(1: ShowStudentInfoRequest req) // ??????????????????

    AdminQueryStudentInfoResponse QueryStudentInfo() // ??????????????????
    AdminQueryTeacherInfoResponse QueryTeacherInfo() // ??????????????????
    AdminQueryCourseInfoResponse QueryCourseInfo() // ??????????????????
    AdminAddStudentInfoResponse AddStudent(1: AdminAddStudentInfoRequest req) // ????????????
    AdminDeleteStudentInfoResponse DeleteStudent(1: AdminDeleteStudentInfoRequest req) // ????????????
    AdminAddTeacherInfoResponse AddTeacher(1: AdminAddTeacherInfoRequest req) // ????????????
    AdminDeleteTeacherInfoResponse DeleteTeacher(1: AdminDeleteTeacherInfoRequest req) // ????????????
    AdminAddCourseInfoResponse AddCourse(1: AdminAddCourseInfoRequest req) // ????????????
    AdminDeleteCourseInfoResponse DeleteCourse(1: AdminDeleteCourseInfoRequest req) // ????????????
}