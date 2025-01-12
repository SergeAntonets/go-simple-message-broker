import React, { useState, useEffect } from "react";
// import { getExchanges } from "../services/api";

const TopicList = () => {
  const [topics, setTopics] = useState([
      {
          name: "Test",
      }
  ]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

//   useEffect(() => {
//     const fetchTopics = async () => {
//       setLoading(true);
//       try {
//         const data = await getExchanges();
//         setTopics(data.filter((topic) => topic.type === "topic"));
//       } catch (err) {
//         setError(err.message);
//       } finally {
//         setLoading(false);
//       }
//     };

//     fetchTopics();
//   }, []);

  if (loading)
    return (
      <div className="flex justify-center items-center h-screen">
        <p className="text-gray-700 text-lg">Loading...</p>
      </div>
    );
  if (error)
    return (
      <div className="flex justify-center items-center h-screen">
        <p className="text-red-500 text-lg">Error: {error}</p>
      </div>
    );

  return (
    <div className="max-w-4xl mx-auto py-8">
      <h1 className="text-2xl font-bold text-center text-gray-800 mb-6">
        Topics
      </h1>
      {topics.length > 0 ? (
        <div className="bg-white shadow-md rounded-lg p-4">
          <ul className="divide-y divide-gray-200">
            {topics.map((topic) => (
              <li key={topic.name} className="py-3">
                <div className="flex justify-between items-center">
                  <span className="text-gray-700 font-medium">{topic.name}</span>
                  <span className="text-sm text-gray-500 uppercase">
                    {topic.type}
                  </span>
                </div>
              </li>
            ))}
          </ul>
        </div>
      ) : (
        <p className="text-center text-gray-600">No topics found.</p>
      )}
    </div>
  );
};

export default TopicList;
