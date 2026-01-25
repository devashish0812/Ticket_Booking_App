import { useState } from "react";
import { useNavigate } from "react-router-dom";
import axios from "axios";

export default function Login() {
  const [form, setForm] = useState({ name: "", password: "" });
  const navigate = useNavigate();

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const API_GATEWAY_URL = process.env.REACT_APP_API_GATEWAY_URL;
      const User_Dashboard_URL = process.env.REACT_APP_User_Dashboard_URL;
      const res = await axios.post(`${API_GATEWAY_URL}/auth/login`, form, {
        withCredentials: true, // <-- send/receive cookies across origins
      });

      const user = res.data.user;

      if (user.role === "user") {
        window.location.href = `${User_Dashboard_URL}`;
      } else {
        window.location.href = `${User_Dashboard_URL}/admin`; // admin dashboard
      }
    } catch (err) {
      console.error(err.response?.data || err.message);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gradient-to-br from-blue-500 via-indigo-500 to-purple-600">
      <form
        onSubmit={handleSubmit}
        className="bg-white p-8 rounded-2xl shadow-2xl w-full max-w-md"
      >
        <h2 className="text-3xl font-bold mb-6 text-center text-gray-800">
          Welcome Back
        </h2>

        <input
          type="text"
          name="name"
          placeholder="Name"
          value={form.name}
          onChange={handleChange}
          className="w-full p-3 mb-4 border rounded-lg focus:ring-2 focus:ring-indigo-400"
        />

        <input
          type="password"
          name="password"
          placeholder="Password"
          value={form.password}
          onChange={handleChange}
          className="w-full p-3 mb-6 border rounded-lg focus:ring-2 focus:ring-indigo-400"
        />

        <button
          type="submit"
          className="w-full bg-indigo-600 text-white py-3 rounded-xl hover:bg-indigo-700 transition font-semibold"
        >
          Login
        </button>

        <p className="text-center mt-4 text-gray-600">
          Don’t have an account?{" "}
          <button
            type="button"
            onClick={() => navigate("/")}
            className="text-indigo-600 font-medium hover:underline"
          >
            Signup
          </button>
        </p>
      </form>
    </div>
  );
}
