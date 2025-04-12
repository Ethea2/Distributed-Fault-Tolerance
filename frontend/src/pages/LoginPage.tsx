import { useState } from "react"
import useLogin from "../hooks/useLogin"

export default function Login() {
  const { loginUser } = useLogin()
  const [username, setUsername] = useState("")
  const [password, setPassword] = useState("")

  const handleLogin = async () => {
    await loginUser(username, password)
  }
  return (
    <div className="w-full min-h-screen flex justify-center items-center">
      <div className="card w-1/4 bg-base-100 card-xl shadow-sm border-2">
        <div className="card-body flex items-center justify-center">
          <h2 className="card-title">School Center of Enrollments</h2>
          <input onChange={(e) => setUsername(e.target.value)} value={username} type="text" placeholder="Username" className="input input-primary" />
          <input onChange={(e) => setPassword(e.target.value)} value={password} type="password" placeholder="Password" className="input input-primary" />
          <div className="justify-end card-actions">
            <button onClick={handleLogin} className="btn btn-primary">Login</button>
          </div>
        </div>
      </div>
    </div>
  )
}
