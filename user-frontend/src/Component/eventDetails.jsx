import React, { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import axiosInstance from "./interceptor";
import logo from "../assets/eventflow-logo-removebg.png";

function EventDetails() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [event, setEvent] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchEventDetails = async () => {
      try {
        const API_GATEWAY_URL = import.meta.env.VITE_API_GATEWAY_URL;
        const res = await axiosInstance.get(`${API_GATEWAY_URL}/events/${id}`);
        setEvent(res.data);
      } catch (err) {
        console.error("Error fetching event details:", err);
        setError("Failed to load event details.");
      } finally {
        setLoading(false);
      }
    };

    if (id) {
      fetchEventDetails();
    }
  }, [id]);

  if (loading)
    return <div className="text-center mt-20">Loading event details...</div>;
  if (error)
    return <div className="text-center mt-20 text-red-500">{error}</div>;
  if (!event) return null;

  return (
    <div className="min-h-screen bg-gray-50 pb-10">
      <div className="flex justify-center py-6 bg-white shadow-sm mb-6">
        <img
          src={logo}
          className="w-40 h-auto cursor-pointer"
          alt="Logo"
          onClick={() => navigate("/")}
        />
      </div>

      <div className="max-w-5xl mx-auto bg-white shadow-lg rounded-lg overflow-hidden">
        {event.bannerImageUrl ? (
          <img
            src={event.bannerImageUrl}
            alt={event.title}
            className="w-full h-full object-contain md:h-96"
          />
        ) : (
          <div className="w-full h-64 bg-gray-200 flex items-center justify-center text-gray-400">
            No Image Available
          </div>
        )}

        <div className="p-8">
          <div className="flex flex-col md:flex-row justify-between items-start md:items-center border-b pb-6 mb-6">
            <div>
              <span className="bg-blue-100 text-blue-800 text-xs font-semibold px-2.5 py-0.5 rounded uppercase tracking-wide">
                {event.category}
              </span>
              <h1 className="text-3xl font-bold mt-2 text-gray-900">
                {event.title}
              </h1>
              <p className="text-gray-500 mt-1 flex items-center">
                📍 {event.venueName}, {event.address}, {event.city},{" "}
                {event.country}
              </p>
            </div>

            <div className="mt-4 md:mt-0 text-right">
              <div className="text-lg font-semibold text-blue-600">
                {new Date(event.startDateTime).toLocaleDateString(undefined, {
                  weekday: "long",
                  year: "numeric",
                  month: "long",
                  day: "numeric",
                })}
              </div>
              <div className="text-gray-600">
                {new Date(event.startDateTime).toLocaleTimeString([], {
                  hour: "2-digit",
                  minute: "2-digit",
                })}
              </div>
            </div>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div className="md:col-span-2">
              <h2 className="text-xl font-bold mb-4">About Event</h2>
              <p className="text-gray-700 whitespace-pre-line leading-relaxed">
                {event.description}
              </p>

              {event.tags && event.tags.length > 0 && (
                <div className="mt-6">
                  <h3 className="text-sm font-semibold text-gray-500 mb-2">
                    TAGS
                  </h3>
                  <div className="flex flex-wrap gap-2">
                    {event.tags.map((tag, index) => (
                      <span
                        key={index}
                        className="px-3 py-1 bg-gray-100 text-gray-600 rounded-full text-sm"
                      >
                        #{tag}
                      </span>
                    ))}
                  </div>
                </div>
              )}
            </div>

            <div className="bg-gray-50 p-6 rounded-xl h-fit">
              <h3 className="text-lg font-bold mb-4">Event Info</h3>

              <div className="space-y-4">
                <div>
                  <p className="text-sm text-gray-500">Language</p>
                  <p className="font-medium">
                    {event.language || "Not specified"}
                  </p>
                </div>
                <div>
                  <p className="text-sm text-gray-500">Age Restriction</p>
                  <p className="font-medium">
                    {event.ageRestriction || "None"}
                  </p>
                </div>
                <div>
                  <p className="text-sm text-gray-500">Refund Policy</p>
                  <p className="font-medium text-red-500">
                    {event.refundPolicy || "Non-refundable"}
                  </p>
                </div>
                <div>
                  <p className="text-sm text-gray-500">Max Tickets Per User</p>
                  <p className="font-medium">{event.maxTicketsPerUser || 10}</p>
                </div>
              </div>

              <button
                className="w-full mt-6 bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-4 rounded transition duration-200"
                onClick={() =>
                  navigate(`/categories/${id}`, {
                    state: {
                      eventTitle: event.title,
                      eventDate: event.startDateTime,
                    },
                  })
                }
              >
                Book Tickets
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default EventDetails;
