import { useEffect, useRef, useState } from "react";
import StudentNavbar from "../components/StudentNavbar";
import { ICourses } from "../types/courses.types";
import useFetchWithAuth from "../hooks/useFetchWithAuth";
import { Id, toast } from "react-toastify";

export default function ViewAvailableCourses() {
  const [data, setData] = useState<ICourses | null>(null)
  const { data: courseData, loading, refetch } = useFetchWithAuth("/course/available")
  const toastID = useRef<Id>(null)

  useEffect(() => {
    if (!loading) {
      if (courseData) {
        setData(courseData as ICourses)
      }
    }
  }, [loading, courseData])

  const enroll = async (course_id: number) => {
    toastID.current = toast.loading("Enrolling you now!")
    // type CourseEnrollRequest struct {
    // 	CourseID int `json:"course_id"`
    // }

    const res = await fetch(`${import.meta.env.VITE_API_URL}/course/enroll`, {
      method: "POST",
      body: JSON.stringify({ course_id }),
      headers: {
        "Authorization": `Bearer ${JSON.parse(localStorage.getItem("user")!).token}`
      }
    })

    const json = await res.json()

    if (!res.ok) {
      toast.update(toastID.current, {
        render: "Unable to enroll you",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: false,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "dark",
        type: "error",
        isLoading: false
      })
      toastID.current = null
      return
    }

    toast.update(toastID.current, {
      render: json.message,
      autoClose: 5000,
      hideProgressBar: false,
      closeOnClick: false,
      pauseOnHover: true,
      draggable: true,
      progress: undefined,
      theme: "dark",
      type: "success",
      isLoading: false
    })
    refetch()
  }
  return (
    <>
      <StudentNavbar />
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
                <div className="w-full flex justify-between items-center px-5 py-2">
                  <p className={`w-1/2 text-start py-2 ${course.available ? "text-secondary" : "text-green-400"}`}>{course.available ? "Ongoing" : "Done"}</p>
                  <button onClick={() => enroll(course.id)} className="btn btn-primary">Enroll</button>
                </div>
              </div>
            )
          })
        }
      </div >
    </>
  )
}
