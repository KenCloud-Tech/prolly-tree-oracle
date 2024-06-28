<template>

    <el-container >
        <el-header>
            <connect :sendOracle="getOracle"></connect>
        </el-header>
        <el-container id="body">
            <el-aside>
                <event-listen v-if="isConnected" :oracle="oracle"></event-listen>
            </el-aside>
            <el-main>
                <div v-if="isConnected">
                    <el-row>
                        <allow-write :oracle="oracle" :baseGasFee="baseGasFee"></allow-write>
                    </el-row>
                    <el-row>
                        <create :oracle="oracle" :baseGasFee="baseGasFee"></create>
                    </el-row>
                    <el-row>
                        <put :oracle="oracle" :gasPerByte="gasPerByte" :baseGasFee="baseGasFee"></put>
                    </el-row>
                    <el-row>
                        <get :oracle="oracle" :baseGasFee="baseGasFee"></get>
                    </el-row>
                    <el-row>
                        <creat-index :oracle="oracle" :baseGasFee="baseGasFee"></creat-index>
                    </el-row>
                    <el-row>
                        <search :oracle="oracle" :baseGasFee="baseGasFee"></search>
                    </el-row>
                    <el-row>
                        <import-unchain :oracle="oracle" :baseGasFee="baseGasFee"></import-unchain>
                    </el-row>
                </div>
            </el-main>
        </el-container>
    </el-container>


</template>
<script setup>
import Connect from "@/components/connect.vue";
import EventListen from "@/components/eventListen.vue";

import {ref} from "vue";
import Create from "@/components/create.vue";
import AllowWrite from "@/components/allowWrite.vue";
import Get from "@/components/get.vue";
import Put from "@/components/put.vue";
import CreatIndex from "@/components/creatIndex.vue";
import Search from "@/components/search.vue";
import ImportUnchain from "@/components/importUnchain.vue";
let oracle

const isConnected = ref(false)
let baseGasFee
let gasPerByte
async function getOracle(contract, connected,base,perByte) {
    oracle = contract
    isConnected.value = connected
    baseGasFee=base
    gasPerByte=perByte
    console.log(oracle)
}


</script>

<style scoped>

.el-header {
    height: 15vh;
    background-color: #337ecc;

}


#body{
    height: 85vh;
}
.el-main {
    background-color: #F2F6FC;
    padding: 0;
}
.el-aside{
    width: 28vw;
    background-color: #c8c9cc;
}
.el-row{
    width: 100%;
    margin: 10px;
}

</style>
