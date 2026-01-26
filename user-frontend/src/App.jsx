import React, { useState, useEffect } from "react";
import { BrowserRouter, Routes, Route } from "react-router-dom";
import axios from "axios";
import ProtectedRoute from "./Component/protected";
import Dashboard from "./Component/dashboard";
import EventDetails from "./Component/eventDetails";
import CategorySelection from "./Component/categorySelection";
import SectionSelection from "./Component/sectionSelection";
import SeatSelection from "./Component/seatSelection";
function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isLoading, setIsLoading] = useState(true);

  const API_GATEWAY_URL = import.meta.env.VITE_API_GATEWAY_URL;

  useEffect(() => {
    const checkAuth = async () => {
      try {
        await axios.post(
          `${API_GATEWAY_URL}/auth/refresh`,
          {},
          { withCredentials: true }
        );

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
            <Route path="/event/:id" element={<EventDetails />} />
            <Route
              path="/categories/:eventId"
              element={<CategorySelection />}
            />
            <Route
              path="/events/:eventId/categories/:categoryName"
              element={<SectionSelection />}
            />
            <Route
              path="/events/:eventId/sections/:sectionName/seats"
              element={<SeatSelection />}
            />
          </Route>
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
