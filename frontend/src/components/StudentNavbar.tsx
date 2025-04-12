import { useEffect } from "react"
import { NavLink, useLocation, useNavigate } from "react-router"
import { toast } from "react-toastify"

export default function StudentNavbar() {
  const location = useLocation()
  const router = useNavigate()

  useEffect(() => {
    if (!localStorage.getItem("user")) {
      toast("Please login first!", {
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: false,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "dark",
        type: "error"
      })
      router("/")
      return
    }
  }, [])

  const logoutUser = () => {
    localStorage.removeItem("user")
    router("/")
  }

  return (
    <div className="navbar z-40 sticky top-0 bg-base-100 shadow-sm">
      <div className="flex-1">
        <a className="btn btn-ghost text-xl">Student Enrollment</a>
      </div>
      <div className="flex-none">
        <ul className="menu menu-horizontal px-1 flex justify-center items-center gap-4">
          <li><NavLink className={`${location.pathname === "/view-available-courses" ? "text-primary" : ""}`} to="/view-available-courses" end>Available Courses</NavLink></li>
          <li><NavLink className={`${location.pathname === "/view-grades" ? "text-primary" : ""}`} to="/view-grades" end>Grades</NavLink></li>
          <li><button onClick={logoutUser} className="btn btn-soft btn-secondary">LOGOUT</button></li>
        </ul>
      </div>
    </div>
  )
}
