// /app/api/forward-post/route.ts
import { NextResponse } from 'next/server';
import { Queue } from "@/interfaces/queue";
import { getQueues } from './get-queue';

// Handle POST requests
export async function POST(request: Request) {
  try {
    // Parse data from the request body
    const data = await request.json();
    console.log("data", data);

    // Define the REST API endpoint
    const apiBaseUrl = process.env.REST_API_BASE_URL;
    const restApiUrl = `${apiBaseUrl}/queue`;

    // Forward the data to the external REST API
    const response = await fetch(restApiUrl, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });

    if (!response.ok) {
      throw new Error(`Failed to forward data: ${response.statusText}`);
    }

    // Return the response from the external REST API
    const responseData = await response.json();
    return NextResponse.json(responseData, { status: response.status });
  } catch (error) {
    console.log(error);
    return NextResponse.json(
      { error: 'An error occurred' },
      { status: 500 }
    );
  }
}

// Handle GET requests
export async function GET() {
  try {
    const data: Queue[] = await getQueues();

    const out = data.map((queue) => ({
      message: queue.message,
      status: queue.status,
    }));

    return NextResponse.json(out, { status: 200 });
  } catch (error) {
    console.log(error);
    return NextResponse.json({ error: 'An error occurred' }, { status: 500 });
  }
}
