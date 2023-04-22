import { AiOutlineUser } from "react-icons/ai";
import { BsPencil } from "react-icons/bs";
import { Sidebar } from "./Sidebar";
import { MenuItem } from "./MenuItem";

export const Menu: React.FC<{}> = () => {
  return (
    <div className="flex flex-row py-8 sm:py-4 space-x-6 justify-center ">
      <Sidebar />
      <MenuItem path="/posts/new">
        <BsPencil className="w-6 h-6 text-blue" />
      </MenuItem>
    </div>
  );
};
