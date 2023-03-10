// axios を import
import React, { useEffect, useState } from "react";
import axios from "axios";

function App() {
  const [getData, setGetData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      const response = await axios.get("http://localhost:8080/health_data");
      setGetData(response.data);
    };
    fetchData();
  }, []);
  console.log("Get data : ", getData);

  return (
    <>
      <ul>
        {getData.map((data) => (
          <li key={data.user_id}>
            month : {data.month}, day : {data.day}, weight : {data.weight}
          </li>
        ))}
      </ul>
    </>
  )
}

export default App;
