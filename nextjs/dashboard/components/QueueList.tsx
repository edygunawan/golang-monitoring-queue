'use client'

import React from 'react';
import { Queue } from '@/interfaces/queue';
import ResendButton from './ResendButton';

interface TableProps {
  data: Queue[];
}

const QueueList: React.FC<TableProps> = ({ data }) => {
  const handleResend = async (queue: Queue) => {
    const response = await fetch('/api/queue', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        message: queue.message
      }),
    });

    if (response.ok) {
      alert("Queue has been resend");
    } else {
      alert("Can't resend queue");
    }
  };

  return (
    <table className="min-w-full table-auto border-collapse border border-gray-300">
      <thead>
        <tr>
          <th className="px-4 py-2 text-left border-b border-gray-300">Message</th>
          <th className="px-4 py-2 text-left border-b border-gray-300">Status</th>
          <th className="px-4 py-2 border-b border-gray-300"></th>
        </tr>
      </thead>
      <tbody>
        {data.map((item, idx) => (
          <tr key={idx}>
            <td className="px-4 py-2 text-left border-b border-gray-300">{item.message}</td>
            <td className="px-4 py-2 text-left border-b border-gray-300">{item.status}</td>
            <td className="px-4 py-2 border-b border-gray-300">
              {item.status === 'failed' && (
                <ResendButton onClick={() => handleResend(item)} label="Resend" />
              )}
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default QueueList;
