import React, { memo } from "react";

const Seat = memo(
  ({ seat, isSelected, onToggle }) => {
    const isBooked = seat.status === "booked";
    const isLocked = seat.status === "locked";

    let seatColor =
      "bg-white border-gray-300 hover:border-blue-500 hover:bg-blue-50";

    if (isBooked || isLocked) {
      seatColor =
        "bg-gray-200 border-gray-200 text-gray-400 cursor-not-allowed";
    } else if (isSelected) {
      seatColor =
        "bg-green-500 border-green-600 text-white shadow-md transform scale-105";
    }

    return (
      <div
        onClick={() => !isBooked && !isLocked && onToggle(seat)}
        className={`
        w-10 h-10 m-1 
        border rounded-md 
        flex items-center justify-center 
        text-xs font-bold cursor-pointer transition-all duration-200
        ${seatColor}
      `}
        title={`Row ${seat.row} - Col ${seat.column}`}
      >
        {seat.column}
      </div>
    );
  },
  (prevProps, nextProps) => {
    return (
      prevProps.isSelected === nextProps.isSelected &&
      prevProps.seat.status === nextProps.seat.status
    );
  }
);

export default Seat;
