import QueueList from '../components/QueueList';
import { getQueues } from './api/queue/get-queue';

export default async function Home() {
  const queues = await getQueues();

  return (
    <div className="w-[1024px] h-screen mx-auto">
      <main>
        <h1 className="mb-5 mt-2">Queue List</h1>
        <QueueList data={queues} />
      </main>
    </div>
  );
}
