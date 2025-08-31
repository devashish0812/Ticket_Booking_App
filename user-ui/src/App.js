import { BrowserRouter, Routes, Route } from "react-router-dom";
import Login from "./Components/login";
import Signup from "./Components/signup";
//import Dashboard from "./components/Dashboard";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<Signup />} />
        {/* <Route path="/dashboard" element={<Dashboard />} />
        <Route path="*" element={<Login />} /> */}
      </Routes>
    </BrowserRouter>
  );
}

export default App;
