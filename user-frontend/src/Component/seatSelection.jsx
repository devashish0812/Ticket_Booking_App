import React, {
  useEffect,
  useState,
  useMemo,
  useCallback,
  useRef,
} from "react";
import { useParams, useNavigate, useLocation } from "react-router-dom";
import axiosInstance from "./interceptor";
import Seat from "./Seat";
import toast from "react-hot-toast";

const SeatSelection = () => {
  const { eventId, sectionName } = useParams();
  const navigate = useNavigate();
  const location = useLocation();
  const { eventTitle, categoryName, categoryPrice } = location.state || {};

  const [seats, setSeats] = useState([]);
  const [selectedSeats, setSelectedSeats] = useState([]);
  const [loading, setLoading] = useState(true);

  const selectedSeatsRef = useRef(selectedSeats);

  useEffect(() => {
    selectedSeatsRef.current = selectedSeats;
  }, [selectedSeats]);

  useEffect(() => {
    const fetchSeats = async () => {
      try {
        const res = await axiosInstance.get(
          `/tickets/events/${eventId}/sections/${sectionName}/seats`
        );
        setSeats(res.data);
        setLoading(false);
      } catch (err) {
        console.error(err);
        setLoading(false);
      }
    };
    if (eventId && sectionName) fetchSeats();
  }, [eventId, sectionName]);

  const seatsByRow = useMemo(() => {
    const grouped = {};
    seats.forEach((seat) => {
      const rowKey = seat.row;
      if (!grouped[rowKey]) grouped[rowKey] = [];
      grouped[rowKey].push(seat);
    });

    const sortedRowKeys = Object.keys(grouped).sort(
      (a, b) => parseInt(a) - parseInt(b)
    );
    const finalSortedSeats = {};

    sortedRowKeys.forEach((key) => {
      const seatsInRow = grouped[key];
      seatsInRow.sort((a, b) => parseInt(a.column) - parseInt(b.column));
      finalSortedSeats[key] = seatsInRow;
    });

    return finalSortedSeats;
  }, [seats]);

  const handleToggleSeat = useCallback(async (seat) => {
    const currentSelectedSeats = selectedSeatsRef.current;
    const isAlreadySelected = currentSelectedSeats.some(
      (s) => s.id === seat.id
    );

    if (isAlreadySelected) {
      setSelectedSeats((prev) => prev.filter((s) => s.id !== seat.id));
      return;
    }

    if (currentSelectedSeats.length >= 6) {
      toast.error("Max 6 seats allowed", {
        style: { borderRadius: "10px", background: "#333", color: "#fff" },
      });
      return;
    }

    setSelectedSeats((prev) => [...prev, seat]);

    try {
      await axiosInstance.post("/tickets/seats/lock", { seatId: seat.id });
    } catch (error) {
      setSelectedSeats((prev) => prev.filter((s) => s.id !== seat.id));

      setSeats((prevSeats) =>
        prevSeats.map((s) =>
          s.id === seat.id ? { ...s, status: "booked" } : s
        )
      );

      toast.error("Missed it! This seat was just booked.", {
        style: { borderRadius: "10px", background: "#333", color: "#fff" },
      });
    }
  }, []);

  const handleProceed = () => {
    navigate("/checkout", {
      state: {
        ...location.state,
        selectedSeats,
        totalPrice: selectedSeats.length * categoryPrice,
      },
    });
  };

  if (loading) return <div className="p-10 text-center">Loading...</div>;

  return (
    <div className="min-h-screen bg-gray-50 flex flex-col items-center pb-24">
      <div className="w-full bg-white shadow-sm p-4 text-center sticky top-0 z-10">
        <h1 className="text-xl font-bold text-gray-800">{eventTitle}</h1>
        <p className="text-sm text-gray-500">
          {categoryName} • Section {sectionName}
        </p>
      </div>

      <div className="mt-8 mb-12 w-3/4 max-w-lg text-center">
        <div className="h-4 bg-blue-200 rounded-t-full w-full mb-2 shadow-sm"></div>
        <span className="text-xs uppercase tracking-widest text-gray-400">
          Stage
        </span>
      </div>

      <div className="bg-white p-6 rounded-xl shadow-lg border border-gray-100 mx-4 overflow-x-auto">
        {Object.keys(seatsByRow).map((rowKey) => (
          <div key={rowKey} className="flex items-center mb-2 min-w-max">
            <div className="w-8 text-right pr-4 font-bold text-gray-400 text-sm">
              {rowKey}
            </div>
            <div className="flex">
              {seatsByRow[rowKey].map((seat) => {
                const isSelected = selectedSeats.some((s) => s.id === seat.id);

                return (
                  <Seat
                    key={seat.id}
                    seat={seat}
                    isSelected={isSelected}
                    onToggle={handleToggleSeat}
                  />
                );
              })}
            </div>
          </div>
        ))}
      </div>

      {selectedSeats.length > 0 && (
        <div className="fixed bottom-0 w-full bg-white border-t border-gray-200 p-4 shadow-xl flex justify-between items-center max-w-4xl mx-auto md:rounded-t-xl">
          <div>
            <p className="text-2xl font-bold text-gray-900">
              ₹ {(selectedSeats.length * categoryPrice).toLocaleString()}
            </p>
            <p className="text-xs text-gray-400">
              {selectedSeats.length} seats
            </p>
          </div>
          <button
            onClick={handleProceed}
            className="bg-black text-white px-8 py-3 rounded-lg font-bold hover:bg-gray-800"
          >
            Proceed
          </button>
        </div>
      )}
    </div>
  );
};

export default SeatSelection;
