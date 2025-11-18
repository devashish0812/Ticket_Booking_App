import logo from "../assets/eventflow-logo-removebg.png";
import { EVENT_CATEGORIES } from "../constants/eventCategories";
import { useEffect } from "react";
import { useState } from "react";
import { useRef } from "react";
import axiosInstance from "./interceptor.jsx";
function Dashboard() {
  const dateInputRef = useRef(null);
  const [events, setEvents] = useState([]);
  const [filters, setFilters] = useState({
    category: "",
    date: "",
    sortBy: "date",
    page: "1",
    order: "asc",
  });

  useEffect(() => {
    const fetchEvents = async (filters) => {
      try {
        const API_GATEWAY_URL = import.meta.env.VITE_API_GATEWAY_URL;
        const res = await axiosInstance.get(
          `${API_GATEWAY_URL}/dashboard?category=${filters.category}&date=${filters.date}&page=${filters.page}&order=${filters.order}`,
          {
            withCredentials: true,
          }
        );

        if (res.data === null || res.data.length === 0) {
          alert("No events found for this category");
          res.data = [];
        }
        setEvents(res.data.events);
        console.log(res.data.totalCount);
      } catch (error) {
        console.error("Error fetching events:", error);
        alert("Failed to fetch events for the selected category.");
      }
    };

    fetchEvents(filters);
  }, [filters]);

  const handleFilterChange = (key, value) => {
    setFilters((prevFilters) => ({
      ...prevFilters,
      [key]: value,
    }));
    // fetchEvents(filters);
  };

  const handleCategoryClick = (category) => {
    handleFilterChange("category", category);
  };
  const handleApplyFilters = () => {
    const committedDate = dateInputRef.current
      ? dateInputRef.current.value
      : "";
    handleFilterChange("date", committedDate);
  };

  return (
    <>
      <div className="flex justify-center py-6">
        <img src={logo} className="w-40 h-auto" alt="Logo" />
      </div>

      <div className="flex gap-8 p-8 items-start">
        <div className="w-1/6 bg-gray-100 p-4 rounded-lg shadow-sm">
          <h2 className="font-bold mb-4 text-lg">Filters</h2>

          <label className="block mb-2 text-sm font-semibold">Category</label>
          <div className="flex flex-wrap gap-2">
            {EVENT_CATEGORIES.map((category) => (
              <button
                key={category}
                className={`px-3 py-1 text-sm font-semibold rounded-full border transition-all duration-150
                ${
                  filters.category === category
                    ? "bg-blue-600 text-white border-blue-600 shadow"
                    : "bg-white text-gray-700 border-gray-300 hover:bg-gray-200"
                }`}
                onClick={
                  filters.category === category
                    ? () => handleCategoryClick("")
                    : () => handleCategoryClick(category)
                }
              >
                {category}
              </button>
            ))}
          </div>

          <div className="mt-6">
            <label className="block mb-2 text-sm font-semibold">Date</label>
            <div className="flex flex-col gap-2">
              <input
                type="date"
                ref={dateInputRef}
                defaultValue={filters.date}
                className="w-full p-2 border rounded focus:ring-2 focus:ring-blue-400"
              />
              <button
                onClick={handleApplyFilters}
                className="w-full px-4 py-2 bg-green-500 text-white rounded-lg hover:bg-green-600 transition-colors"
              >
                Apply
              </button>
            </div>
          </div>

          <div className="mt-6">
            <label className="block mb-2 text-sm font-semibold">Sort By</label>
            <select
              className="w-full p-2 border border-gray-300 rounded focus:ring-2 focus:ring-blue-400"
              onChange={(e) => {
                const [sortBy, order] = e.target.value.split("-");
                handleFilterChange("sortBy", sortBy);
                handleFilterChange("order", order);
              }}
              value={`${filters.sortBy}-${filters.order}`}
            >
              <option value="date-asc">Date (Ascending) 📅</option>
              <option value="date-desc">Date (Descending) ⬇️</option>
              <option value="category-asc">Category Name (A-Z) 🅰️</option>
              <option value="category-desc">Category Name (Z-A) 🇿</option>
              <option value="price-asc">Price (Low to High) 💰</option>
            </select>
          </div>
        </div>

        <div className="flex-1 w-1/3">
          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
            {events.map((event) => (
              <div
                key={event.id}
                className="bg-white shadow-md p-4 rounded-lg hover:shadow-lg transition-shadow"
              >
                <img
                  src={event.bannerImageUrl}
                  alt={event.title}
                  className="w-full aspect-[17/20] object-cover rounded"
                />
                <h3 className="text-lg font-bold mt-2">{event.title}</h3>
                <h4 className="text-gray-600 text-sm">
                  {event.venueName}, {event.city}
                </h4>
                <p className="text-gray-500 text-sm">{event.category}</p>
                <p className="text-gray-500 text-sm">
                  {new Date(event.startDateTime).toLocaleDateString()}
                </p>
              </div>
            ))}
          </div>
        </div>

        <div className="w-1/6"></div>
      </div>
    </>
  );
}
export default Dashboard;
