export default function Login() {

  return (
    <div className="w-full min-h-screen flex justify-center items-center">
      <div className="card w-1/4 bg-base-100 card-xl shadow-sm border-2">
        <div className="card-body flex items-center justify-center">
          <h2 className="card-title">School Center of Enrollments</h2>
          <input type="text" placeholder="Username" className="input input-primary" />
          <input type="password" placeholder="Password" className="input input-primary" />
          <div className="justify-end card-actions">
            <button className="btn btn-primary">Login</button>
          </div>
        </div>
      </div>
    </div>
  )
}
