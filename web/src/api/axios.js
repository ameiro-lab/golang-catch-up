/** axio設定ファイル */
import axios from 'axios';

// axiosインスタンスの作成
const axiosInstance = axios.create({
  baseURL: 'http://localhost:8080', // SpringBootのポート
  timeout: 10000, // タイムアウトの設定
  headers: {
    'Content-Type': 'application/json',
  },
});

export default axiosInstance;