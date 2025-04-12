
// type Grade struct {
// 	CourseName       string  `json:"course_name"`
// 	CourseCode       string  `json:"course_code"`
// 	Available        bool    `json:"available"`
// 	EnrollmentStatus string  `json:"enrollment_status"`
// 	Grade            float32 `json:"grade"`
// }
//
//


export interface IGrade {
  course_name: string
  course_code: string
  available: boolean
  enrollment_status: string
  grade: number
}


// type StudentGrades struct {
// 	UserID   int     `json:"user_id"`
// 	Username string  `json:"username"`
// 	Grades   []Grade `json:"grades"`
// }

export interface IStudentGrades {
  user_id: number
  username: string
  grades: IGrade[]
}

// type StudentGradesResponse struct {
// 	StudentGrades []StudentGrades `json:"student_grades"`
// }

export interface IStudentGradesResponse {
  student_grades: IStudentGrades[]
}

export interface IGrades {
  grades: IGrade[]
}
