import { useEffect, useState } from "react"
import { ICourses } from "../types/courses.types"
import useFetch from "../hooks/useFetch"
import FacultyNavbar from "../components/FacultyNavbar"

export default function FacultyViewAllCourses() {
  const [data, setData] = useState<ICourses | null>(null)
  const { data: courseData, loading } = useFetch("/course")

  useEffect(() => {
    if (!loading) {
      setData(courseData as ICourses)
    }
  }, [loading, courseData])

  return (
    <>
      <FacultyNavbar />
      <div className="w-full min-h-screen flex flex-wrap justify-center items-start gap-10 p-10">
        {
          data &&
          data.courses.map((course, index) => {
            return (
              <div key={index} className="card card-dash bg-base-100 w-1/3 border-2 border-primary">
                <div className="card-body">
                  <h2 className="card-title">{course.course_name}</h2>
                  <p>{course.course_code}</p>
                  <p>{course.course_description}</p>
                </div>
                <p className={`w-full text-center py-2 ${course.available ? "text-secondary" : "text-green-400"}`}>{course.available ? "Ongoing" : "Done"}</p>
              </div>
            )
          })
        }
      </div >
    </>
  )
}
