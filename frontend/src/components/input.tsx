interface InputProps {
  label: string;
  value: string;
  placeholder?: string;
  name: string;
  type?: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
}

export default function Input({
  label,
  value,
  placeholder,
  name,
  type,
  onChange,
}: InputProps) {
  return (
    <div className="input-container">
      <label htmlFor={name}>{label}</label>
      <input
        type={type ? type : "text"}
        value={value}
        name={name}
        placeholder={placeholder}
        onChange={onChange}
        id={name}
      />
    </div>
  );
}
