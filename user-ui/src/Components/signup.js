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
    <div className="flex justify-center items-center min-h-screen bg-gradient-to-br from-indigo-500 via-purple-500 to-pink-500">
      <form
        onSubmit={handleSubmit}
        className="bg-white p-8 rounded-2xl shadow-2xl w-full max-w-md"
      >
        <h2 className="text-3xl font-bold mb-6 text-center text-gray-800">
          Create Account
        </h2>

        <input
          type="text"
          name="name"
          placeholder="Name"
          value={form.name}
          onChange={handleChange}
          className="w-full p-3 mb-3 border rounded-lg focus:ring-2 focus:ring-indigo-400"
        />

        <input
          type="email"
          name="email"
          placeholder="Email"
          value={form.email}
          onChange={handleChange}
          className="w-full p-3 mb-3 border rounded-lg focus:ring-2 focus:ring-indigo-400"
        />

        <input
          type="password"
          name="password"
          placeholder="Password"
          value={form.password}
          onChange={handleChange}
          className="w-full p-3 mb-3 border rounded-lg focus:ring-2 focus:ring-indigo-400"
        />

        <select
          name="role"
          value={form.role}
          onChange={handleChange}
          className="w-full p-3 mb-4 border rounded-lg focus:ring-2 focus:ring-indigo-400"
        >
          <option value="user">User</option>
          <option value="organizer">Organizer</option>
        </select>

        {form.role === "organizer" && (
          <>
            <input
              type="text"
              name="orgName"
              placeholder="Organisation Name"
              value={form.orgName}
              onChange={handleChange}
              className="w-full p-3 mb-3 border rounded-lg focus:ring-2 focus:ring-indigo-400"
            />
            <input
              type="text"
              name="address"
              placeholder="Organisation Address"
              value={form.address}
              onChange={handleChange}
              className="w-full p-3 mb-3 border rounded-lg focus:ring-2 focus:ring-indigo-400"
            />
            <input
              type="text"
              name="contactNo"
              placeholder="Contact Number"
              value={form.contactNo}
              onChange={handleChange}
              className="w-full p-3 mb-3 border rounded-lg focus:ring-2 focus:ring-indigo-400"
            />
          </>
        )}

        <button
          type="submit"
          className="w-full bg-indigo-600 text-white py-3 rounded-xl hover:bg-indigo-700 transition font-semibold"
        >
          Signup
        </button>

        <p className="text-center mt-4 text-gray-600">
          Already a user?{" "}
          <button
            type="button"
            onClick={() => navigate("/login")}
            className="text-indigo-600 font-medium hover:underline"
          >
            Login
          </button>
        </p>
      </form>
    </div>
  );
}
