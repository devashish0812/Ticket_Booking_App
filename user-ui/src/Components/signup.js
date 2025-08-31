import { useState } from "react";
import { useNavigate } from "react-router-dom";

export default function Signup() {
  const [form, setForm] = useState({ name: "",
     email: "",
     password: "",
     role: "user",
     orgName:"",
     contactNo:"",
     address:"" });
  const navigate = useNavigate();

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const USER_SERVICE_URL = process.env.REACT_APP_USER_SERVICE_URL;
      const res = await fetch(`${USER_SERVICE_URL}/users/signup`, {
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
    <div className="flex justify-center items-center min-h-screen bg-gray-100">
      <form
        onSubmit={handleSubmit}
        className="bg-white p-6 rounded-2xl shadow-md w-96"
      >
        <h2 className="text-2xl font-semibold mb-4">Signup</h2>

        <input
          type="text"
          name="name"
          placeholder="Name"
          value={form.name}
          onChange={handleChange}
          className="w-full p-2 mb-3 border rounded"
        />
        <input
          type="email"
          name="email"
          placeholder="Email"
          value={form.email}
          onChange={handleChange}
          className="w-full p-2 mb-3 border rounded"
        />
        <input
          type="password"
          name="password"
          placeholder="Password"
          value={form.password}
          onChange={handleChange}
          className="w-full p-2 mb-3 border rounded"
        />

        <select
          name="role"
          value={form.role}
          onChange={handleChange}
          className="w-full p-2 mb-4 border rounded"
        >
          <option value="user">User</option>
          <option value="organizer">Organizer</option>
        </select>

        {/* Organizer extra fields */}
        {form.role === "organizer" && (
          <>
            <input
              type="text"
              name="orgName"
              placeholder="Organisation Name"
              value={form.orgName}
              onChange={handleChange}
              className="w-full p-2 mb-3 border rounded"
            />
            <input
              type="text"
              name="address"
              placeholder="Organisation Address"
              value={form.address}
              onChange={handleChange}
              className="w-full p-2 mb-3 border rounded"
            />
            <input
              type="text"
              name="contactnumber"
              placeholder="Contact Number"
              value={form.contactnumber}
              onChange={handleChange}
              className="w-full p-2 mb-3 border rounded"
            />
          </>
        )}

        <button
          type="submit"
          className="w-full bg-indigo-600 text-white py-2 rounded-xl hover:bg-indigo-700"
        >
          Signup
        </button>
      </form>
    </div>
  );

}
