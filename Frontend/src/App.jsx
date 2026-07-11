import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom'
import LoginComponent from './LoginComponent.jsx'
import SignUp from './SignUp.jsx'
import HomePageWithoutLogin from "./HomePageWithoutLogin.jsx";

function App() {
  return (
    <BrowserRouter>
      <Routes>
          <Route path="/" element={<HomePageWithoutLogin/>}/>
        <Route path="/login" element={<LoginComponent />} />
        <Route path="/SignUp" element={<SignUp />} />
        <Route path="*" element={<Navigate to="/" replace />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App

