import { useEffect, useState } from "react";
import StudentNavbar from "../components/StudentNavbar";
import { IGrade } from "../types/grades.types";
import useFetchWithAuth from "../hooks/useFetchWithAuth";

export default function ViewGrades() {
  const [data, setData] = useState<IGrade[]>()
  const { data: pastGrades, loading } = useFetchWithAuth("/grade")

  useEffect(() => {
    if (!loading) {
      if (pastGrades) {
        setData(pastGrades.grades as IGrade[])
      }
    }
  }, [loading, pastGrades])
  return (
    <>
      <StudentNavbar />
      <div className="w-full min-h-screen flex flex-col items-center justify-start p-10">
        {
          data &&
          data.map((courseGrades) => {
            return (
              <>
                <div className="card card-dash bg-base-100 w-full border-2 border-primary my-5">
                  <div className="flex justify-between items-center p-5">
                    <h2 className="card-title w-1/3 text-center">{courseGrades.course_name}</h2>
                    <p className="w-1/3 text-center">{courseGrades.course_code}</p>
                    <div className="w-1/3 flex justify-end gap-2">
                      <p className="w-full flex justify-end items-center">Final Grade: <span className="mx-5 w-5 font-bold">{courseGrades.grade}</span></p>
                    </div>
                  </div>
                  <p className={`w-full text-center py-2 ${courseGrades.available ? "text-secondary" : "text-green-400"}`}>{courseGrades.available ? "Ongoing" : "Done"}</p>
                </div>

              </>
            )
          })
        }
      </div>
    </>
  )
}
