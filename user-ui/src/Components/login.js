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
  // 1) Login -> backend will SET cookies (access_token, refresh_token)
    const USER_SERVICE_URL = process.env.REACT_APP_USER_SERVICE_URL;
    const API_GATEWAY_URL = process.env.REACT_APP_API_GATEWAY_URL;
  await axios.post(`${USER_SERVICE_URL}/users/login`, form, {
    withCredentials: true, // <-- send/receive cookies across origins
  });

  // 2) Ask API Gateway where to go; cookies will be sent automatically
  const routeRes = await axios.get(`${API_GATEWAY_URL}/gateway/dashboard`, {
    withCredentials: true, // <-- include cookies on this request too
  });

  const { dashboardUrl } = routeRes.data;
  window.location.href = dashboardUrl;
} catch (err) {
      console.error(err.response?.data || err.message);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <form
        onSubmit={handleSubmit}
        className="bg-white p-6 rounded-2xl shadow-md w-96"
      >
        <h2 className="text-2xl font-semibold mb-4">Login</h2>

        <input
          type="text"
          name="name"
          placeholder="Name"
          value={form.name}
          onChange={handleChange}
          className="w-full p-2 mb-3 border rounded"
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          value={form.password}
          onChange={handleChange}
          className="w-full p-2 mb-4 border rounded"
        />

        <button
          type="submit"
          className="w-full bg-indigo-600 text-white py-2 rounded-xl hover:bg-indigo-700"
        >
          Login
        </button>
      </form>
    </div>
  );
}
