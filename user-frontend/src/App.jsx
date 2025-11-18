import React, { useState, useEffect } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import axios from "axios";
import ProtectedRoute from "./Component/protected";
import Dashboard from "./Component/dashboard";

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  const API_GATEWAY_URL = import.meta.env.VITE_API_GATEWAY_URL;

  useEffect(() => {
    const checkAuth = async () => {
      try {
        await axios.get(`${API_GATEWAY_URL}/auth/refresh`, {
          withCredentials: true,
        });
        setIsAuthenticated(true);
      } catch (error) {
        setIsAuthenticated(false);
      } finally {
        setIsLoading(false);
      }
    };

    checkAuth();
  }, []);
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route
            element={
              <ProtectedRoute
                isAuthenticated={isAuthenticated}
                isLoading={isLoading}
              />
            }
          >
            <Route path="/dashboard" element={<Dashboard />} />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
