import { BrowserRouter as Router } from 'react-router'
import './App.css'
import PageRouter from './components/router/PageRouter'

function App() {

  return (
    <Router>
      <PageRouter />
    </Router>
  )
}

export default App
