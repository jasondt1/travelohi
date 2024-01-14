import { useState } from "react";
import Calendar from "react-calendar";
import "react-calendar/dist/Calendar.css";

interface DatePickerProps {
  selectedDate: Date | null;
  setSelectedDate: (date: Date | null) => void;
  placeholder?: string;
}

export default function DatePicker({
  selectedDate,
  setSelectedDate,
  placeholder,
}: DatePickerProps) {
  const [isCalendarOpen, setIsCalendarOpen] = useState(false);

  const handleInputFocus = () => {
    setIsCalendarOpen(true);
  };

  const handleCalendarChange = (date: any) => {
    setSelectedDate(date as Date);
    setIsCalendarOpen(false);
  };

  return (
    <div className="input-container">
      <label>Date of Birth</label>
      <input
        placeholder={placeholder}
        value={selectedDate ? selectedDate.toLocaleDateString("en-US") : ""}
        onFocus={handleInputFocus}
        readOnly
      />
      {isCalendarOpen && (
        <Calendar
          className="calendar"
          onChange={handleCalendarChange}
          value={selectedDate}
        />
      )}
    </div>
  );
}
