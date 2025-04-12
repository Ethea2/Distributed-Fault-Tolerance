package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

type LoginModel struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseToken struct {
	Token   string `json:"token"`
	Message string `json:"message"`
	Role    string `json:"role"`
}

type Course struct {
	ID                int    `json:"id"`
	CourseName        string `json:"course_name"`
	CourseCode        string `json:"course_code"`
	CourseDescription string `json:"course_description"`
	Availabie         bool   `json:"available"`
}

type CourseEnrollRequest struct {
	CourseID int `json:"course_id"`
}

type CourseResponse struct {
	Courses []Course `json:"courses"`
}

type GeneralResponse struct {
	Message string `json:"message"`
}

type Grade struct {
	CourseName       string  `json:"course_name"`
	CourseCode       string  `json:"course_code"`
	Available        bool    `json:"available"`
	EnrollmentStatus string  `json:"enrollment_status"`
	Grade            float32 `json:"grade"`
}

type GradeResponse struct {
	Grades []Grade `json:"grades"`
}

type StudentGrades struct {
	UserID   int     `json:"user_id"`
	Username string  `json:"username"`
	Grades   []Grade `json:"grades"`
}

type StudentGradesResponse struct {
	StudentGrades []StudentGrades `json:"student_grades"`
}

type ChangeGradeRequest struct {
	UserID     int     `json:"user_id"`
	CourseCode string  `json:"course_code"`
	Grade      float32 `json:"grade"`
}
