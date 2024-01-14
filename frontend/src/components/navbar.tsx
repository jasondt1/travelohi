import { useState } from "react";
import logo from "../assets/logo.png";
import { VscSearch } from "react-icons/vsc";
import { VscMenu } from "react-icons/vsc";
import { Link } from "react-router-dom";

export default function Navbar() {
  const [open, setOpen] = useState(false);

  const toggleOpen = () => {
    setOpen(!open);
  };

  return (
    <div className="center parrent">
      <VscMenu className="menu" onClick={toggleOpen} />
      <div className="navbar">
        <div className="logo-container">
          <Link to="/">
            {" "}
            <img
              src={logo}
              alt="logo"
              className={open ? "logo open" : "logo"}
            />
          </Link>
        </div>
        <div className={open ? "left-nav open" : "left-nav"}>
          <div className="search-container">
            <input type="text" className="search" placeholder="Search..." />
            <VscSearch className="search-icon" />
          </div>
        </div>
        <div className={open ? "right-nav open" : "right-nav"}>
          <Link className="link" to="/login">
            Log In
          </Link>
          <Link className="link" to="/register">
            Register
          </Link>
        </div>
      </div>
    </div>
  );
}
