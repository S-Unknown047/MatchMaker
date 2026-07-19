import { Routes, Route, Navigate } from 'react-router-dom'
import LoginComponent from './LoginComponent.jsx'
import SignUp from './SignUp.jsx'
import HomePageWithoutLogin from "./HomePageWithoutLogin.jsx";
import HomePage from "./HomePage.jsx";
import RequireAuth from "./RequireAuth.jsx";
import Unauthorized from "./Unauthorized.jsx";

const Roles = {
  'User': 'user',
  'Admin': 'admin'
}

function App() {
  return (
    <Routes>
      <Route path="/" element={<HomePageWithoutLogin />} />
      <Route path="/login" element={<LoginComponent />} />
      <Route path="/SignUp" element={<SignUp />} />
      <Route path="/unauthorized" element={<Unauthorized />} />
      <Route element={<RequireAuth allowedRoles={[Roles['User']]} />}>
        <Route path="/home" element={<HomePage />} />
      </Route>
      <Route path="*" element={<Navigate to="/" replace />} />
    </Routes>
  )
}

export default App
