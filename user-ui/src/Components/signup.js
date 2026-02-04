import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Signup() {
  const [form, setForm] = useState({
    name: "",
    email: "",
    password: "",
    role: "user",
    orgName: "",
    contactNo: "",
    address: "",
  });

  const navigate = useNavigate();

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const API_GATEWAY_URL = process.env.REACT_APP_API_GATEWAY_URL;
      const res = await fetch(`${API_GATEWAY_URL}/auth/signup`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(form),
      });

      if (res.ok) {
        navigate("/login");
      } else {
        alert("Signup failed");
      }
    } catch (err) {
      console.error(err);
    }
  };

  return (
    <div className="flex justify-center items-center min-h-screen bg-slate-950 text-slate-200 p-4">
      <div className="w-full max-w-md">
        <form
          onSubmit={handleSubmit}
          className="bg-slate-900 border border-slate-800 p-10 rounded-xl shadow-2xl"
        >
          <div className="mb-8 text-center">
            <h2 className="text-3xl font-extrabold tracking-tight text-white mb-2">
              Create Account
            </h2>
            <p className="text-slate-400 text-sm">
              Join the high-concurrency ticketing platform
            </p>
          </div>

          <div className="space-y-4">
            <div>
              <label className="block text-xs font-semibold uppercase tracking-wider text-slate-500 mb-2 ml-1">
                Full Name
              </label>
              <input
                type="text"
                name="name"
                placeholder="John Doe"
                value={form.name}
                onChange={handleChange}
                className="w-full bg-slate-800 border border-slate-700 p-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition-all placeholder-slate-600"
              />
            </div>

            <div>
              <label className="block text-xs font-semibold uppercase tracking-wider text-slate-500 mb-2 ml-1">
                Email Address
              </label>
              <input
                type="email"
                name="email"
                placeholder="john@example.com"
                value={form.email}
                onChange={handleChange}
                className="w-full bg-slate-800 border border-slate-700 p-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition-all placeholder-slate-600"
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
                className="w-full bg-slate-800 border border-slate-700 p-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition-all placeholder-slate-600"
              />
            </div>

            <div>
              <label className="block text-xs font-semibold uppercase tracking-wider text-slate-500 mb-2 ml-1">
                Account Type
              </label>
              <select
                name="role"
                value={form.role}
                onChange={handleChange}
                className="w-full bg-slate-800 border border-slate-700 p-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition-all text-slate-300"
              >
                <option value="user">Standard User</option>
                <option value="organizer">Event Organizer</option>
              </select>
            </div>

            {form.role === "organizer" && (
              <div className="space-y-4 pt-2 border-t border-slate-800 mt-4">
                <input
                  type="text"
                  name="orgName"
                  placeholder="Organization Name"
                  value={form.orgName}
                  onChange={handleChange}
                  className="w-full bg-slate-800 border border-slate-700 p-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition-all placeholder-slate-600"
                />
                <input
                  type="text"
                  name="address"
                  placeholder="Official Address"
                  value={form.address}
                  onChange={handleChange}
                  className="w-full bg-slate-800 border border-slate-700 p-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition-all placeholder-slate-600"
                />
                <input
                  type="text"
                  name="contactNo"
                  placeholder="Contact Number"
                  value={form.contactNo}
                  onChange={handleChange}
                  className="w-full bg-slate-800 border border-slate-700 p-3 rounded-lg focus:outline-none focus:ring-2 focus:ring-indigo-500 transition-all placeholder-slate-600"
                />
              </div>
            )}

            <button
              type="submit"
              className="w-full bg-indigo-600 text-white py-3 mt-6 rounded-lg hover:bg-indigo-500 active:transform active:scale-[0.98] transition-all font-bold shadow-lg shadow-indigo-500/20"
            >
              Create Account
            </button>
          </div>

          <div className="mt-8 text-center border-t border-slate-800 pt-6">
            <p className="text-sm text-slate-500">
              Already have an account?{" "}
              <button
                type="button"
                onClick={() => navigate("/login")}
                className="text-indigo-400 font-semibold hover:text-indigo-300 transition"
              >
                Sign In
              </button>
            </p>
          </div>
        </form>
      </div>
    </div>
  );
}
