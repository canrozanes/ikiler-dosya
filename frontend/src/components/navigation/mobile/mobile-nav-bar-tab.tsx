import { NavLink } from "react-router-dom";

interface MobileNavBarTabProps {
  path: string;
  label: string;
  handleClick: () => void;
}

export const MobileNavBarTab = ({
  path,
  label,
  handleClick,
}: MobileNavBarTabProps) => {
  return (
    <NavLink
      onClick={handleClick}
      to={path}
      end
      className={({ isActive }) =>
        "mobile-nav-bar__tab " + (isActive ? "mobile-nav-bar__tab--active" : "")
      }
    >
      {label}
    </NavLink>
  );
};
