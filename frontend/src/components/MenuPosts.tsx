import { MenuItem } from "./MenuItem";

export const MenuPosts: React.FC<{}> = () => {
  return (
    <div className="flex flex-row justify-around gap-two bg-box-color text-sm p-one">
      <MenuItem path={"/posts"}> FOLLOWING </MenuItem>

      <MenuItem path={"/"}> FOR YOU </MenuItem>

      <MenuItem path={"/posts/you"}> YOUR POSTS </MenuItem>
    </div>
  );
};
