import { Link } from 'react-router-dom';
import React from 'react';

interface ILinkWrapperProps {
  link: string;
  children: React.ReactElement;
}

// LinkWrapper component can be used to wrap another component inside a link. This can be used for selectedable cards,
// which should navigate the user to another location. This is to prefer over an onClick handler, so that the user can
// decide if he wants the link in a new tab or not.
const LinkWrapper: React.FunctionComponent<ILinkWrapperProps> = ({ children, link }: ILinkWrapperProps) => {
  return (
    <Link
      to={link}
      target={link.startsWith('http') ? '_blank' : undefined}
      style={{ color: 'inherit', textDecoration: 'inherit' }}
    >
      {children}
    </Link>
  );
};

export default LinkWrapper;
