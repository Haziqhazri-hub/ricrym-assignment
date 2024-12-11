import React, { useState, useEffect } from "react";

const GameRankingDashboard = () => {
  const [players, setPlayers] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1); // Total pages will be fetched from API
  const [loading, setLoading] = useState(true);

  // Fetch data for a specific page
  const fetchData = async (page) => {
    try {
      setLoading(true);
      const response = await fetch(
        `https://ricrym-assignment.onrender.com/pagination/${page}`
      );
      const data = await response.json();
      setPlayers(data.data); // Assuming API returns players in `players`
      setTotalPages(data.totalPages); // Assuming API returns total pages in `totalPages`
      //   console.log(data.data);
    } catch (error) {
      console.error("Error fetching data:", error);
    } finally {
      setLoading(false);
    }
  };

  // Fetch initial data on page load
  useEffect(() => {
    fetchData(currentPage);
  }, [currentPage]);

  // Change the page and fetch new data
  const handlePageChange = (newPage) => {
    if (newPage > 0 && newPage <= totalPages) {
      setCurrentPage(newPage);
    }
  };

  return (
    <div className="min-h-screen bg-gray-900 text-white flex justify-center items-center">
      <div className="w-full max-w-4xl p-6 bg-gray-800 rounded-lg shadow-lg">
        <h1 className="text-3xl font-semibold text-center mb-6">
          Players Ranking
        </h1>

        {loading ? (
          <div className="text-center text-xl">Loading...</div>
        ) : (
          <>
            <div className="overflow-x-auto">
              <table className="min-w-full table-auto">
                <thead>
                  <tr>
                    <th className="px-4 py-2 text-left text-lg">Rank</th>
                    <th className="px-4 py-2 text-left text-lg">Username</th>
                    <th className="px-4 py-2 text-left text-lg">Score</th>
                  </tr>
                </thead>
                <tbody>
                  {players.map((player) => (
                    <tr key={player.rank} className="border-b border-gray-700">
                      <td className="px-4 py-2">{player.rank}</td>
                      <td className="px-4 py-2">{player.username}</td>
                      <td className="px-4 py-2">{player.total_score}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>

            {/* Pagination Controls */}
            <div className="flex justify-between mt-6">
              <button
                className="px-4 py-2 bg-blue-500 text-white rounded disabled:bg-gray-500"
                onClick={() => handlePageChange(currentPage - 1)}
                disabled={currentPage === 1}
              >
                Previous
              </button>
              <span className="self-center text-lg">
                Page {currentPage} of {totalPages}
              </span>
              <button
                className="px-4 py-2 bg-blue-500 text-white rounded disabled:bg-gray-500"
                onClick={() => handlePageChange(currentPage + 1)}
                disabled={currentPage === totalPages}
              >
                Next
              </button>
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default GameRankingDashboard;
