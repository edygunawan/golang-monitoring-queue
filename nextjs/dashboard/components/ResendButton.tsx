import React from 'react';

interface ButtonProps {
  onClick: () => void;
  label?: string;
}

const ResendButton: React.FC<ButtonProps> = ({ onClick, label = "Resend" }) => {
  return (
    <button
      onClick={onClick}
      className="flex items-center px-4 py-2 bg-blue-500 text-white rounded-md shadow-md hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-400 focus:ring-offset-2 transition"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
        strokeWidth={2}
        stroke="currentColor"
        className="w-5 h-5 mr-2"
      >
        <path
          strokeLinecap="round"
          strokeLinejoin="round"
          d="M4.5 9.75a7.5 7.5 0 0114.298-2.28M19.5 14.25a7.5 7.5 0 01-14.298 2.28M12 6v3m0 6v3m3-6H9"
        />
      </svg>
      {label}
    </button>
  );
};

export default ResendButton;
