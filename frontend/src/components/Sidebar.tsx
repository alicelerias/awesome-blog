import { useState } from "react";
import { AiOutlineUser } from "react-icons/ai";
import { useMutation } from "react-query";
import { useNavigate } from "react-router-dom";
import { logout } from "../api/mutations";

export const Sidebar = () => {
  const [isOpen, setOpen] = useState(false);

  const handleDropDown = () => {
    setOpen(!isOpen);
  };

  const navigate = useNavigate();

  const { mutate } = useMutation(() => logout(), {
    onSuccess: () => {
      setTimeout(() => {
        navigate("/login");
      }, 2000);
    },
  });

  const onLogout = () => {
    mutate();
  };

  return (
    <div className="w-6">
      <button
        className="text-white bg-blue-700 hover:text-blue items-center"
        onClick={handleDropDown}
      >
        <AiOutlineUser className="w-6 h-6" />
      </button>

      <div
        id="dropdown"
        className={`z-10 w-44 bg-white rounded divide-y divide-gray-100 shadow ${
          isOpen ? "block" : "hidden"
        }`}
      >
        <ul className=" z-10 w-44 bg-white rounded divide-y divide-gray-100 shadow ">
          <li>
            <a
              href="/profile"
              className="text-black block py-2 px-4 hover:text-blue"
            >
              Profile
            </a>
            <a
              href="/favorites"
              className="text-black block py-2 px-4 hover:text-blue border-b border-box-color"
            >
              Favorites
            </a>
            <p
              onClick={onLogout}
              className="text-black block py-2 px-4 hover:text-blue cursor-pointer"
            >
              Logout
            </p>
          </li>
        </ul>
      </div>
    </div>
  );
};
