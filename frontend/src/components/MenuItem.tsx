import { PropsWithChildren } from "react";
import { Link } from "react-router-dom";

type props = {
  path: string;
};
export const MenuItem: React.FC<PropsWithChildren<props>> = ({
  path,
  children,
}) => {
  return (
    <div className="px-one text-bold hover:text-blue transition duration-150 ease-in">
      <Link to={path}>{children}</Link>
    </div>
  );
};
