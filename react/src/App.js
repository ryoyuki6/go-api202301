// axios を import
import axios from 'axios';

function App() {

  // サーバーにリクエストしてレスポンスを待っている状態
  const fetchData = async () => {
    const response = await axios.get(

      // レスポンスデータの取得、交換
      // Prommise が Fulfilled または Rejected となるまで待つ
      // Promiseインスタンスの状態が変わったら処理を再開する
      // "https://jsonplaceholder.typicode.com/users"
      "http://localhost:8080/health_data"
    );

    // 非同期処理が完了したら返される
    return console.log(response.data);
  };
  fetchData();

  return (
    <p>axios!!!</p>
  );
}

export default App;
