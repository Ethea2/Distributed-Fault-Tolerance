import { useEffect, useState } from "react";
import FacultyNavbar from "../components/FacultyNavbar";
import { IStudentGrades } from "../types/grades.types";
import useFetchWithAuth from "../hooks/useFetchWithAuth";
import StudentGrades from "../components/StudentGrades";

export default function FacultyViewGrades() {
  const [data, setData] = useState<IStudentGrades[]>()
  const { data: studentGradesData, loading } = useFetchWithAuth("/grade/all_students")

  useEffect(() => {
    if (!loading) {
      if (studentGradesData) {
        setData(studentGradesData.student_grades as IStudentGrades[])
      }
    }
  }, [studentGradesData, loading])

  return (
    <>
      <FacultyNavbar />
      <div className="flex justify-center items-center text-3xl font-bold text-secondary w-full p-5">STUDENTS</div>
      <div className="w-full min-h-screen flex flex-col justify-start items-center gap-5 p-5">{
        data &&
        data.map((student) => {
          return (
            <StudentGrades student={student} user_id={student.user_id} />
          )
        })
      }</div>
    </>
  )
}
