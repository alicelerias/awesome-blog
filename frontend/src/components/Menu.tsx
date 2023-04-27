import { AiOutlineUser } from "react-icons/ai";
import { BsPencil } from "react-icons/bs";
import { Sidebar } from "./Sidebar";
import { MenuItem } from "./MenuItem";
import { NavigateFunction } from "react-router-dom";
import { useMutation } from "react-query";
import { logout } from "../api/mutations";

type props = {
  navigate: NavigateFunction;
};

export const Menu: React.FC<props> = ({ navigate }) => {
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
    <div className="flex flex-row py-8 sm:py-4 space-x-6 justify-center ">
      <Sidebar
        name={<AiOutlineUser className="h-6 w-6" />}
        children={
          <>
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
          </>
        }
      />
      <MenuItem path="/posts/new">
        <BsPencil className="w-6 h-6 text-blue" />
      </MenuItem>
    </div>
  );
};
