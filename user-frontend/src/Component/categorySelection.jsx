import React, { useEffect, useState } from "react";
import { useParams, useNavigate, useLocation } from "react-router-dom";
import axiosInstance from "./interceptor";
import logo from "../assets/eventflow-logo-removebg.png";

function CategorySelection() {
  const { eventId } = useParams();
  const navigate = useNavigate();
  const location = useLocation();

  const { eventTitle, eventDate, eventVenue } = location.state || {};

  const [categories, setCategories] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchCategories = async () => {
      try {
        debugger;
        const API_GATEWAY_URL = import.meta.env.VITE_API_GATEWAY_URL;
        const res = await axiosInstance.get(
          `${API_GATEWAY_URL}/tickets/categories/${eventId}`
        );
        console.log(res.data);
        setCategories(res.data);
      } catch (err) {
        console.error("Error fetching categories:", err);
        setError("Failed to load categories.");
      } finally {
        setLoading(false);
      }
    };

    if (eventId) fetchCategories();
  }, [eventId]);

  const handleCategorySelect = (category) => {
    navigate(`/events/${eventId}/categories/${category.category_name}`, {
      state: {
        eventTitle,
        eventDate,
        eventVenue,
        categoryName: category.category_name,
        categoryPrice: category.price,
      },
    });
  };

  if (loading)
    return <div className="text-center mt-20">Loading categories...</div>;
  if (error)
    return <div className="text-center mt-20 text-red-500">{error}</div>;

  return (
    <div className="min-h-screen bg-gray-50 pb-10">
      <div className="flex justify-center py-6 bg-white shadow-sm mb-6 sticky top-0 z-10">
        <img
          src={logo}
          className="w-32 h-auto cursor-pointer"
          alt="Logo"
          onClick={() => navigate("/")}
        />
      </div>

      <div className="max-w-3xl mx-auto px-4">
        <div className="mb-8">
          <h1 className="text-2xl font-bold text-gray-900">
            {eventTitle || "Select Category"}
          </h1>
          {eventDate && (
            <p className="text-gray-500">
              {new Date(eventDate).toLocaleDateString()} • {eventVenue}
            </p>
          )}
        </div>

        <div className="mb-10 text-center">
          <div className="mx-auto w-3/4 h-10 bg-gray-800 text-gray-200 rounded-b-xl flex items-center justify-center shadow-md">
            <span className="text-xs font-bold tracking-[0.4em]">STAGE</span>
          </div>
        </div>

        <div className="grid gap-4">
          {categories.map((cat) => {
            const isSoldOut = cat.available_tickets <= 0;

            return (
              <button
                key={cat.category_name}
                disabled={isSoldOut}
                onClick={() => handleCategorySelect(cat)}
                className={`w-full text-left bg-white border border-gray-200 rounded-xl p-6 shadow-sm flex justify-between items-center transition-all group ${
                  isSoldOut
                    ? "opacity-60 cursor-not-allowed"
                    : "hover:border-blue-500 hover:shadow-md active:scale-[0.98]"
                }`}
              >
                <div>
                  <h3 className="text-xl font-bold text-gray-800 group-hover:text-blue-600 transition-colors">
                    {cat.category_name}
                  </h3>
                  <p className="text-gray-500 text-sm mt-1">
                    {cat.available_tickets} seats left out of{" "}
                    {cat.total_tickets}
                  </p>
                </div>

                <div className="text-right">
                  {isSoldOut ? (
                    <span className="text-red-500 font-bold uppercase text-sm tracking-widest">
                      Sold Out
                    </span>
                  ) : (
                    <>
                      <div className="text-2xl font-bold text-gray-900">
                        ₹{cat.price.toLocaleString()}
                      </div>
                      <div className="text-blue-600 text-sm font-semibold flex items-center justify-end mt-1">
                        Select Category{" "}
                        <span className="ml-1 group-hover:translate-x-1 transition-transform">
                          →
                        </span>
                      </div>
                    </>
                  )}
                </div>
              </button>
            );
          })}
        </div>
      </div>
    </div>
  );
}

export default CategorySelection;
