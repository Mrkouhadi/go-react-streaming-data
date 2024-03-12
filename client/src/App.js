// App.js
import React, { useEffect, useState } from 'react';

function App() {
  const [events, setEvents] = useState([]);

  useEffect(() => {
    const eventSource = new EventSource("http://localhost:8080/events");

    // Event listener for messages from the server
    eventSource.onmessage = (event) => {
      setEvents(prevEvents => [...prevEvents, event.data]);
    };

    // Event listener for connection errors
    eventSource.onerror = (error) => {
      console.error("EventSource failed:", error);
      eventSource.close();
    };

    // Cleanup function to close the connection when the component unmounts
    return () => {
      eventSource.close();
    };
  }, []);

  return (
    <div>
      <h1>Server-Sent Events (SSE) Example</h1>
      <ul>
        {events.map((event, index) => (
          <li key={index}>{event}</li>
        ))}
      </ul>
    </div>
  );
}

export default App;
