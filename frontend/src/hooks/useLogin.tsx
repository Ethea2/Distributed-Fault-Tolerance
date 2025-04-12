import { useRef } from "react"
import { useNavigate } from "react-router"
import { Id, toast } from "react-toastify"

const useLogin = () => {
  const toastID = useRef<Id>(null)
  const router = useNavigate()

  const loginUser = async (username: string, password: string) => {
    toastID.current = toast.loading("Logging you in...")
    if (username === "" || password === "") {
      toast.update(toastID.current, {
        render: "Please complete all fields!",
        autoClose: 5000,
        hideProgressBar: false,
        closeOnClick: false,
        pauseOnHover: true,
        draggable: true,
        progress: undefined,
        theme: "dark",
        type: "error"
      })
      return
    }

    const res = await fetch(`${import.meta.env.VITE_API_URL}/auth/login`, {
      method: "POST",
      body: JSON.stringify({ username, password })
    })

    const json = await res.json()

    if (!res.ok) {
      console.log(json)
      toast.update(toastID.current, {
        render: "Something went wrong!",
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
    }

    localStorage.setItem("user", JSON.stringify({ token: json.token, role: json.role }))
    toast.update(toastID.current, {
      render: "Logged you in!",
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
    if (json.role === "student") {
      router("/view-available-courses")
    } else {
      router("/view-all-courses")
    }
  }

  return { loginUser }
}

export default useLogin
