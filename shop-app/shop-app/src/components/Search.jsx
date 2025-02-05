import React, { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import axios from "axios";

const Searchcomp = () => {
  const [results, setResults] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(false);

  const location = useLocation();
  const query = new URLSearchParams(location.search).get("query"); // ดึงคำค้นหาจาก URL

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      setError(false);
      setResults([]); // ล้างผลลัพธ์เก่าเมื่อมีการค้นหาใหม่

      if (query) {
        try {
          const response = await axios.get(`http://localhost:8080/api/v1/products/search?word=A4&sort=asc${query}`);
          if (response.data.length > 0) {
            setResults(response.data);
          } else {
            setError(true); // หากไม่พบข้อมูล ให้ตั้งค่า error เป็น true
          }
        } catch (error) {
          console.error("Error fetching data:", error);
          setError(true);
        }
      }
      setLoading(false);
    };

    fetchData();
  }, [query]);

  return (
    <div className="Searchcomp-results-container">
      {loading && <p>Loading...</p>}
      {!loading && error && <h2>404 Error <br />  No results found for "{query}"</h2>}
      {!loading && !error && (
        <div className="Searchcomp-results">
          {results.map((result, index) => (
            <div key={index} className="result-item">
              <h3>{result.title}</h3>
              <p>{result.description}</p>
            </div>
          ))}
        </div>
      )}
    </div>
  );
};

export default Searchcomp;