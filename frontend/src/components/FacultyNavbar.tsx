import { useEffect } from "react";
import { NavLink, useLocation, useNavigate } from "react-router";
import { toast } from "react-toastify";

export default function FacultyNavbar() {
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

    const role = JSON.parse(localStorage.getItem("user")!).role

    if (role !== "faculty") {
      toast("Students are not allowed here!", {
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: false,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "dark",
        type: "error"
      })
      router("/view-available-courses")
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
        <a className="btn btn-ghost text-xl">Faculty Grader</a>
      </div>
      <div className="flex-none">
        <ul className="menu menu-horizontal px-1 flex justify-center items-center gap-4">
          <li><NavLink className={`${location.pathname === "/view-all-courses" ? "text-primary" : ""}`} to="/view-all-courses" end>Courses</NavLink></li>
          <li><NavLink className={`${location.pathname === "/view-all-grades" ? "text-primary" : ""}`} to="/view-all-grades" end>Grader</NavLink></li>
          <li><button onClick={logoutUser} className="btn btn-soft btn-secondary">LOGOUT</button></li>
        </ul>
      </div>
    </div>
  )
} 
