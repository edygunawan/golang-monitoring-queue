import { Queue } from "@/interfaces/queue";

export const getQueues = async (): Promise<Queue[]> => {
  const apiBaseUrl = process.env.REST_API_BASE_URL;
  const response = await fetch(`${apiBaseUrl}/queue-messages`);

  if (!response.ok) {
    throw new Error(`Error fetching data: ${response.statusText}`);
  }

  const data: Queue[] = await response.json();

  return data.map((queue) => ({
    message: queue.message,
    status: queue.status,
  }));
};
