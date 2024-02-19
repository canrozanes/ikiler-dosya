interface PageFooterHyperlinkProps {
  children: React.ReactNode;
  path: string;
}

export const PageFooterHyperlink = ({
  children,
  path,
}: PageFooterHyperlinkProps) => {
  return (
    <a
      className="page-footer__hyperlink"
      href={path}
      target="_blank"
      rel="noopener noreferrer"
    >
      {children}
    </a>
  );
};
