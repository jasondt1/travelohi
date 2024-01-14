interface DropdownProps {
  options: any;
  label?: string;
  placeholder?: string;
  attribute: string;
  selectedOption: string | null;
  setSelectedOption: any;
}

export default function Dropdown({
  options,
  label,
  placeholder,
  attribute,
  selectedOption,
  setSelectedOption,
}: DropdownProps) {
  const handleSelect = (option: string) => {
    setSelectedOption(option);
  };

  return (
    <div className="input-container">
      <label htmlFor="drop">{label}</label>
      <select
        onChange={(e) => handleSelect(e.target.value)}
        value={selectedOption || ""}
        className="dropdown"
        id="drop"
      >
        <option value="">{placeholder}</option>
        {options.map((option: any, index: any) => (
          <option key={index} value={option[attribute]}>
            {option[attribute]}
          </option>
        ))}
      </select>
    </div>
  );
}
