import React from "react";
import Tabs from "./components/Tabs";

const App = () => (
  <div className="bg-gray-100 min-h-screen">
    <header className="bg-blue-600 text-white py-4">
      <h1 className="text-center text-xl font-semibold">RabbitMQ UI</h1>
    </header>
    <main>
      <Tabs />
    </main>
  </div>
);

export default App;
