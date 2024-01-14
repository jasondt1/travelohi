import LoginPage from "./pages/login-page";
import RegisterPage from "./pages/register-page";
import MainTemplate from "./templates/main-template";
import { BrowserRouter, Route, Routes } from "react-router-dom";
import "./styles/main.css";
import "./styles/navbar.css";
import "./styles/scrollbar.css";
import "./styles/calendar.css";
import "react-datepicker/dist/react-datepicker.css";

function App() {
  return (
    <BrowserRouter>
      <MainTemplate>
        <Routes>
          <Route path="/" element={<LoginPage />}></Route>
          <Route path="/login" element={<LoginPage />}></Route>
          <Route path="/register" element={<RegisterPage />}></Route>
        </Routes>
      </MainTemplate>
    </BrowserRouter>
  );
}

export default App;
