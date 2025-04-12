import FacultyViewAllCourses from "../../pages/FacultyViewAllCourses"
import FacultyViewGrades from "../../pages/FacultyViewGrades"
import Login from "../../pages/LoginPage"
import ViewAvailableCourses from "../../pages/ViewAvailableCourses"
import ViewGrades from "../../pages/ViewGrades"
import { RouteType } from "../../types/router.types"

const routes: Array<RouteType> = [
  { path: "/", element: <Login /> },
  { path: "/view-grades", element: <ViewGrades /> },
  { path: "/view-all-courses", element: <FacultyViewAllCourses /> },
  { path: "/view-all-grades", element: <FacultyViewGrades /> },
  { path: "/view-available-courses", element: <ViewAvailableCourses /> }
]

export default routes
