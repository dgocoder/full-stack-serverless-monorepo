import { type ReactNode } from 'react';

export type ButtonProps = {
  children: ReactNode;
  onClick: () => void;
};
export const Button = ({ children, onClick }: ButtonProps) => {
  return <button onClick={onClick}>{children}</button>;
};
