
// type Course struct {
// 	ID                int    `json:"id"`
// 	CourseName        string `json:"course_name"`
// 	CourseCode        string `json:"course_code"`
// 	CourseDescription string `json:"course_description"`
// 	Availabie         bool   `json:"available"`
// }
//
export interface ICourse {
  id: number
  course_name: string
  course_code: string
  course_description: string
  available: string
}

export interface ICourses {
  courses: Array<ICourse>
}
