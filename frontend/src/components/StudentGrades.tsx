import { IStudentGrades } from "../types/grades.types";
import StudentGradesContainer from "./StudentGradesContainer";

export default function StudentGrades({ student, user_id }: {
  student: IStudentGrades
  user_id: number
}) {
  return (
    <>
      <div className="collapse collapse-plus bg-base-100 border-primary border-2">
        <input type="checkbox" />
        <div className="collapse-title font-semibold">{student.username}</div>
        <div className="collapse-content text-sm">
          {
            [...student.grades]
              .sort((a, b) => {
                if (a.available === b.available) return 0;
                return a.available ? -1 : 1;
              })
              .map((grade) => (
                <StudentGradesContainer courseGrades={grade} user_id={user_id} />
              ))
          }
        </div>
      </div>
    </>
  )
}
