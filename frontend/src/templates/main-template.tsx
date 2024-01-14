import Navbar from "../components/navbar";
import { IChildren } from "../interfaces/children-interface";

export default function MainTemplate({ children }: IChildren) {
  return (
    <div className="w-full bg">
      <Navbar />
      <div className="parent-container">
        <div className="inner-container">{children}</div>
      </div>
    </div>
  );
}
