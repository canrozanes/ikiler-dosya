interface MobileMenuToggleButtonProps {
  icon: string;
  handleClick: () => void;
}

export const MobileMenuToggleButton = ({
  icon,
  handleClick,
}: MobileMenuToggleButtonProps) => {
  return (
    <span
      className="mobile-nav-bar__toggle material-icons"
      id="mobile-menu-toggle-button"
      onClick={handleClick}
    >
      {icon}
    </span>
  );
};
