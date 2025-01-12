import React, { useState } from "react";
import TopicList from "./TopicList";

const Tabs = () => {
  const [activeTab, setActiveTab] = useState("topics");

  const renderContent = () => {
    switch (activeTab) {
      case "topics":
        return <TopicList />;
      case "queues":
        return <div className="text-center text-gray-500">Queues (Coming Soon)</div>;
      case "connections":
        return <div className="text-center text-gray-500">Connections (Coming Soon)</div>;
      default:
        return null;
    }
  };

  return (
    <div className="max-w-4xl mx-auto py-8">
      <div className="flex border-b border-gray-300 mb-6">
        <button
          className={`flex-1 py-2 text-center ${
            activeTab === "topics"
              ? "border-b-2 border-blue-600 text-blue-600 font-medium"
              : "text-gray-500"
          }`}
          onClick={() => setActiveTab("topics")}
        >
          Topics
        </button>
        <button
          className={`flex-1 py-2 text-center ${
            activeTab === "queues"
              ? "border-b-2 border-blue-600 text-blue-600 font-medium"
              : "text-gray-500"
          }`}
          onClick={() => setActiveTab("queues")}
        >
          Queues
        </button>
        <button
          className={`flex-1 py-2 text-center ${
            activeTab === "connections"
              ? "border-b-2 border-blue-600 text-blue-600 font-medium"
              : "text-gray-500"
          }`}
          onClick={() => setActiveTab("connections")}
        >
          Connections
        </button>
      </div>
      <div>{renderContent()}</div>
    </div>
  );
};

export default Tabs;
