import React, { useEffect, useState } from "react";
import { useParams, useNavigate, useLocation } from "react-router-dom";
import axiosInstance from "./interceptor";
import logo from "../assets/eventflow-logo-removebg.png";

function SectionSelection() {
  const { eventId, categoryName } = useParams();
  const navigate = useNavigate();
  const location = useLocation();

  const { eventTitle, eventDate, eventVenue, categoryPrice } =
    location.state || {};

  const [sections, setSections] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchSections = async () => {
      try {
        const API_GATEWAY_URL = import.meta.env.VITE_API_GATEWAY_URL;

        const res = await axiosInstance.get(
          `${API_GATEWAY_URL}/tickets/events/${eventId}/categories/${categoryName}`
        );
        console.log(res.data);
        setSections(res.data);
      } catch (err) {
        console.error("Error fetching sections:", err);
        setError("Failed to load sections for this category.");
      } finally {
        setLoading(false);
      }
    };

    if (eventId && categoryName) fetchSections();
  }, [eventId, categoryName]);

  const handleSectionSelect = (section) => {
    navigate(`/events/${eventId}/sections/${section.section_name}/seats`, {
      state: {
        eventTitle,
        eventDate,
        eventVenue,
        categoryName,
        categoryPrice,
        sectionName: section.section_name,
      },
    });
  };

  if (loading)
    return <div className="text-center mt-20">Loading sections...</div>;
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

      <div className="max-w-4xl mx-auto px-4">
        <div className="mb-8">
          <div className="flex items-center text-sm text-blue-600 mb-2">
            <span
              className="cursor-pointer hover:underline"
              onClick={() => navigate(-1)}
            >
              {eventTitle}
            </span>
            <span className="mx-2 text-gray-400">/</span>
            <span className="text-gray-500 font-medium">{categoryName}</span>
          </div>
          <h1 className="text-3xl font-bold text-gray-900">Select Section</h1>
          <p className="text-gray-500 mt-1">
            {eventVenue} • ₹{categoryPrice?.toLocaleString()} per seat
          </p>
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {sections.map((section) => {
            const isSoldOut = section.available_tickets <= 0;
            const fillPercentage =
              (section.available_tickets / section.capacity) * 100;

            return (
              <div
                key={section.section_name}
                onClick={() => !isSoldOut && handleSectionSelect(section)}
                className={`bg-white border-2 rounded-2xl p-6 transition-all relative overflow-hidden flex flex-col justify-between ${
                  isSoldOut
                    ? "opacity-60 grayscale border-gray-200 cursor-not-allowed"
                    : "border-gray-100 hover:border-blue-500 hover:shadow-xl cursor-pointer active:scale-95 shadow-sm"
                }`}
              >
                {!isSoldOut && (
                  <div
                    className="absolute bottom-0 left-0 h-1 bg-blue-500 transition-all duration-500"
                    style={{ width: `${fillPercentage}%` }}
                  />
                )}

                <div>
                  <h3 className="text-xl font-bold text-gray-800 mb-1">
                    {section.section_name}
                  </h3>
                  <p className="text-sm text-gray-500 uppercase tracking-wider font-semibold">
                    {categoryName}
                  </p>
                </div>

                <div className="mt-8">
                  {isSoldOut ? (
                    <div className="text-center py-2 bg-red-50 text-red-600 rounded-lg font-bold text-sm uppercase">
                      Fully Booked
                    </div>
                  ) : (
                    <div className="flex justify-between items-end">
                      <div>
                        <span className="text-2xl font-black text-gray-900">
                          {section.available_tickets}
                        </span>
                        <span className="text-gray-400 text-sm ml-1">
                          / {section.capacity} seats
                        </span>
                      </div>
                      <div className="text-blue-600 font-bold group">
                        Join →
                      </div>
                    </div>
                  )}
                </div>
              </div>
            );
          })}
        </div>
      </div>
    </div>
  );
}

export default SectionSelection;
