<template>
    <v-sheet class="pa-12" color="primary" rounded>
        <!-- ログイン -->
        <v-card class="mx-auto px-6 py-8" max-width="350">
            <v-card-title>Golang - CatchUp🐾</v-card-title>
            <v-card-text>{{ text }}</v-card-text>
            <!-- フォーム -->
            <v-text-field
                v-model="userName"
                :readonly="loading"
                label="User Name"
                clearable />
            <v-text-field
                v-model="password"
                :readonly="loading"
                label="Password"
                clearable />
            <!-- ボタン -->
            <v-btn color="blue"
                @click="onClick">button</v-btn>
            
        </v-card>
    </v-sheet>
</template>

<script setup>
import { ref } from 'vue'
import api from '@/api/api'

/** リアクティブデータの定義 */
const text = ref("View the response.");
const form = ref(false);
const loading = ref(false);
const userName = ref(null);
const password = ref(null);

/** メソッドの定義 */
const onClick = async () => {

    // ローディングを開始する。
    loading.value = true;

    // リクエストボディを作成する。
    const data = {
        'userName': userName.value,
        'password': password.value
    };

    // API通信（POST）
    api.login(data).then(res => {
        const response = res.data;

        // レスポンスの確認
        const resStatus = response.status;
        const resData = response.data;
        const resMessage = response.message;
        console.log('status:', resStatus);
        console.log('data:', resData);
        console.log('message:', resMessage);

        if (resStatus === 200) {
            // データをセットする。
            text.value = 'User Name: ' + resData.userName + ' Password: ' + resData.password;
        }

    }).catch(error => {
        console.error("Error:", error);

    });

    // ローディングを終了する。
    loading.value = false;
};


</script>

<style scoped>

</style>