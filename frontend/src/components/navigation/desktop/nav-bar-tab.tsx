import { NavLink } from "react-router-dom";

interface NavBarTabProps {
  path: string;
  label: string;
}

export const NavBarTab = ({ path, label }: NavBarTabProps) => {
  return (
    <NavLink
      to={path}
      end
      className={({ isActive }) =>
        "nav-bar__tab " + (isActive ? "nav-bar__tab--active" : "")
      }
    >
      {label}
    </NavLink>
  );
};
