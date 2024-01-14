import { useEffect, useState } from "react";
import Input from "../components/input";
import { Link } from "react-router-dom";
import axios from "axios";
import DatePickerComponent from "../components/datepicker";
import Dropdown from "../components/dropdown";
import ReCAPTCHA from "react-google-recaptcha";

export default function RegisterPage() {
  const [form, setForm] = useState({
    email: "",
    pass: "",
    confirmPass: "",
    firstName: "",
    lastName: "",
    gender: "",
    answer: "",
  });

  const [dateOfBirth, setDateOfBirth] = useState<Date | null>(null);
  const [selectedQuestion, setSelectedQuestion] = useState<string>("");
  const [questions, setQuestions] = useState<object[]>([]);
  const [isChecked, setIsChecked] = useState(false);
  const [isCaptchaVerified, setIsCaptchaVerified] = useState(false);

  const [error, setError] = useState("");
  const [page, setPage] = useState(2);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setForm((prevForm) => ({
      ...prevForm,
      [name]: value,
    }));
  };

  const handleGenderChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm((prevForm) => ({
      ...prevForm,
      gender: e.target.value,
    }));
  };

  const handleCheckboxChange = () => {
    setIsChecked(!isChecked);
  };

  const handleCaptchaChange = () => {
    setIsCaptchaVerified(true);
  };

  const handleContinue = async () => {
    const passRegex =
      /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[!@#$%^&*()_+{}\[\]:;<>,.?~\\/-]).{8,30}$/;
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!emailRegex.test(form.email)) {
      setError("Email is invalid");
      return;
    }
    let validateEmail = false;
    try {
      const response = await axios.post(
        "http://localhost:8080/api/check-email",
        { email: form.email }
      );
      validateEmail = response.data.message;
    } catch (error) {
      console.error("Error checking email:", error);
    }
    if (!validateEmail) {
      setError("Email already exists");
      return;
    }
    if (!passRegex.test(form.pass)) {
      setError(
        "Password must contain at least 8 characters, 1 uppercase, 1 lowercase, 1 number, and 1 special character"
      );
      return;
    }
    if (form.pass !== form.confirmPass) {
      setError("Password and confirm password must be the same");
      return;
    }
    setPage(2);
  };

  const fetchSecurityQuestions = async () => {
    try {
      const response = await axios.get(
        "http://localhost:8080/api/get-security-questions"
      );
      setQuestions(response.data.questions);
    } catch (error) {}
  };

  useEffect(() => {
    fetchSecurityQuestions();
  }, []);

  return (
    <div className="child">
      <div className="card">
        <div className="register-form">
          {page === 1 && (
            <>
              <h1 className="title">Register</h1>
              <Input
                label="Email"
                name="email"
                value={form.email}
                onChange={handleChange}
                placeholder="Input email..."
              />
              <Input
                label="Password"
                name="pass"
                value={form.pass}
                onChange={handleChange}
                placeholder="Input password..."
                type="password"
              />
              <Input
                label="Confirm Password"
                name="confirmPass"
                value={form.confirmPass}
                onChange={handleChange}
                placeholder="Input confirm password..."
                type="password"
              />
              <button className="primary-btn" onClick={handleContinue}>
                Continue
              </button>
              <p className="color-red">{error}</p>
              <p>
                Already have an account ?{" "}
                <Link to="/login" className="secondary-btn">
                  {" "}
                  Login
                </Link>
              </p>
            </>
          )}
          {page === 2 && (
            <>
              <h1 className="title">Personal Details</h1>
              <Input
                label="First Name"
                name="firstName"
                value={form.firstName}
                onChange={handleChange}
                placeholder="Input first name..."
              />
              <Input
                label="Last Name"
                name="lastName"
                value={form.lastName}
                onChange={handleChange}
                placeholder="Input last name..."
              />
              <DatePickerComponent
                selectedDate={dateOfBirth}
                setSelectedDate={setDateOfBirth}
                placeholder="Input date of birth..."
              />

              <div className="gender-container">
                <p>Select a Gender</p>
                <label>
                  <input
                    type="radio"
                    name="gender"
                    value="male"
                    checked={form.gender === "male"}
                    onChange={handleGenderChange}
                    className="mr-2"
                  />
                  Male
                </label>

                <label>
                  <input
                    type="radio"
                    name="gender"
                    value="female"
                    checked={form.gender === "female"}
                    onChange={handleGenderChange}
                    className="mr-2"
                  />
                  Female
                </label>
              </div>

              <Dropdown
                options={questions}
                label="Security Question"
                placeholder="Select a security question"
                attribute="question"
                selectedOption={selectedQuestion}
                setSelectedOption={setSelectedQuestion}
              />

              <Input
                label="Answer"
                name="answer"
                value={form.answer}
                onChange={handleChange}
                placeholder="Input answer..."
              />

              <div className="checkbox-container">
                <input
                  type="checkbox"
                  checked={isChecked}
                  onChange={handleCheckboxChange}
                  id="check"
                />
                <label htmlFor="check">Subscribe to newsletter</label>
              </div>

              <ReCAPTCHA
                sitekey="6LfemVApAAAAAJ-4KZyVTKsFnrzmkWSz3DAZ9jD9
                "
                onChange={handleCaptchaChange}
              />

              <button className="primary-btn">Register</button>
              <p className="color-red">{error}</p>
              <p>
                Already have an account ?{" "}
                <Link to="/login" className="secondary-btn">
                  {" "}
                  Login
                </Link>
              </p>
            </>
          )}
        </div>
      </div>
    </div>
  );
}
