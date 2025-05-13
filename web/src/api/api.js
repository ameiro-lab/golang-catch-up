/** リクエストメソッドの定義ファイル */
import axios from './axios';

export default {

  /**
   * ログインする。
   * @param {*} data 
   * @returns 
   */
  login: (data) => {
    return axios.post('/api/login', data);  // POSTリクエストでボディにデータを送信
  },

    /**
   * GETリクエストでデータを送信する例
   * paramsオプションで、クエリパラメータとして送信する。
   * @returns 
   */
    // name: (data) => {
    //   return axios.get('/api/login', { params: data } );
    // },
};
