import { BrowserRouter as Router } from 'react-router'
import './App.css'
import PageRouter from './components/router/PageRouter'
import { Bounce, ToastContainer } from 'react-toastify'
import "react-toastify/ReactToastify.css"

function App() {

  return (
    <Router>
      <ToastContainer
        position="bottom-center"
        autoClose={5000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick={false}
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
        theme="dark"
        transition={Bounce}
      />
      <PageRouter />
    </Router>
  )
}

export default App
