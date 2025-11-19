import React from "react";
import { Outlet } from "react-router-dom";

const LOGIN_APP_URL = import.meta.env.VITE_LOGIN_APP_URL;

const ProtectedRoute = ({ isAuthenticated, isLoading }) => {
  if (isLoading) {
    return <div>Loading session...</div>;
  }

  if (!isAuthenticated) {
    window.location.href = LOGIN_APP_URL;
    return null;
  }

  return <Outlet />;
};

export default ProtectedRoute;
