import { useRef, useState } from "react";
import { IGrade } from "../types/grades.types";
import { Id, toast } from "react-toastify";

export default function StudentGradesContainer({ courseGrades, user_id }: {
  courseGrades: IGrade
  user_id: number
}) {
  const [grade, setGrade] = useState<number>(courseGrades.grade)
  const toastID = useRef<Id>(null)
  const submit = async () => {
    toastID.current = toast.loading("Submitting changes now...")

    // type ChangeGradeRequest struct {
    // 	UserID     int     `json:"user_id"`
    // 	CourseCode string  `json:"course_code"`
    // 	Grade      float32 `json:"grade"`
    // }
    const res = await fetch(`${import.meta.env.VITE_API_URL}/grade`, {
      method: "POST",
      body: JSON.stringify({ user_id: user_id, course_code: courseGrades.course_code, grade: grade }),
      headers: {
        "Authorization": `Bearer ${JSON.parse(localStorage.getItem("user")!).token}`
      }
    })

    const json = await res.json()

    if (!res.ok) {
      toast.update(toastID.current, {
        render: "Something went wrong with changing the grade.",
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
  }
  return (
    <>
      <div className="card card-dash bg-base-100 w-full border-2 border-primary my-5">
        <div className="flex justify-between items-center p-5">
          <h2 className="card-title w-1/3 text-center">{courseGrades.course_name}</h2>
          <p className="w-1/3 text-center">{courseGrades.course_code}</p>
          <div className="w-1/3 flex justify-end gap-2">
            <select
              className="border-2 border-info rounded-lg px-3 py-2"
              value={grade}
              onChange={(e) => setGrade(Number(e.target.value))}
            >
              <option value="-1">-1.0</option>
              <option value="0">0.0</option>
              <option value="1">1.0</option>
              <option value="1.5">1.5</option>
              <option value="2">2.0</option>
              <option value="2.5">2.5</option>
              <option value="3">3.0</option>
              <option value="3.5">3.5</option>
              <option value="4">4.0</option>
            </select>
            <button onClick={submit} className="btn btn-secondary w-1/4">Change Grade</button>
          </div>
        </div>
        <p className={`w-full text-center py-2 ${courseGrades.available ? "text-secondary" : "text-green-400"}`}>{courseGrades.available ? "Ongoing" : "Done"}</p>
      </div>
    </>
  )
}
