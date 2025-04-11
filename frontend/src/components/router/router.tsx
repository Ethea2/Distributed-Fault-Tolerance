import Login from "../../pages/LoginPage"
import ViewGrades from "../../pages/ViewGrades"
import { RouteType } from "../../types/router.types"

const routes: Array<RouteType> = [
  { path: "/", element: <Login /> },
  { path: "/view-grades", element: <ViewGrades /> }
]

export default routes
