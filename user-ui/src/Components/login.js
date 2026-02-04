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
        withCredentials: true,
      });

      const user = res.data.user;
      if (user.role === "user") {
        window.location.href = `${User_Dashboard_URL}`;
      } else {
        window.location.href = `${User_Dashboard_URL}/admin`;
      }
    } catch (err) {
      console.error(err.response?.data || err.message);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-slate-950 text-slate-200">
      <div className="w-full max-w-md px-4">
        <form
          onSubmit={handleSubmit}
          className="bg-slate-900 border border-slate-800 p-10 rounded-xl shadow-2xl"
        >
          <div className="mb-8 text-center">
            <h2 className="text-3xl font-extrabold tracking-tight text-white mb-2">
              Sign In
            </h2>
            <p className="text-slate-400 text-sm">
              Enter your credentials to access the ticketing dashboard
            </p>
          </div>

          <div className="space-y-5">
            <div>
              <label className="block text-xs font-semibold uppercase tracking-wider text-slate-500 mb-2 ml-1">
                Name
              </label>
              <input
                type="text"
                name="name"
                placeholder="Devashish"
                value={form.name}
                onChange={handleChange}
                className="w-full bg-slate-800 border border-slate-700 p-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all placeholder-slate-600"
              />
            </div>

            <div>
              <label className="block text-xs font-semibold uppercase tracking-wider text-slate-500 mb-2 ml-1">
                Password
              </label>
              <input
                type="password"
                name="password"
                placeholder="••••••••"
                value={form.password}
                onChange={handleChange}
                className="w-full bg-slate-800 border border-slate-700 p-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent transition-all placeholder-slate-600"
              />
            </div>

            <button
              type="submit"
              className="w-full bg-indigo-600 text-white py-3 mt-4 rounded-lg hover:bg-indigo-500 active:transform active:scale-[0.98] transition-all font-bold shadow-lg shadow-indigo-500/20"
            >
              Log In
            </button>
          </div>

          <div className="mt-8 text-center border-t border-slate-800 pt-6">
            <p className="text-sm text-slate-500">
              New to the platform?{" "}
              <button
                type="button"
                onClick={() => navigate("/")}
                className="text-indigo-400 font-semibold hover:text-indigo-300 transition"
              >
                Create an account
              </button>
            </p>
          </div>
        </form>
      </div>
    </div>
  );
}
