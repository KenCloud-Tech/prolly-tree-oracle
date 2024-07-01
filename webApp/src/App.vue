<template>

    <el-container >
        <el-header>
            <div id="user" v-if="isConnected">User: {{ wallet.address.slice(0,10)+'...' }}</div>
            <div id="user" v-if="oracleIsConnected">BaseGasFee: {{ baseGasFee.valueOf() }} Wei</div>
            <div id="user" v-if="oracleIsConnected">GasPerByte: {{ gasPerByte.valueOf() }} Wei/Byte</div>
            <el-button id="connect_button" type="primary" @click="onClickConnectButton()">{{ isConnected ? 'Disconnect Wallet' : 'Connect Wallet' }}</el-button>
            <VueDappModal dark auto-connect />
        </el-header>
        <el-container id="body">
            <el-aside>
                <event-listen v-if="oracleIsConnected" :oracle="oracle"/>
            </el-aside>
            <el-main style="position:relative;">
                <div v-if="isConnected && !oracleIsConnected" id="oracleList">
                    <div v-for="oracleName in oracleList">
                        <el-button class="oracle" @click="onClickOracle(oracleName)" color="#626aef" :dark="isDark" plain>{{oracleName}}</el-button>
                    </div>
                </div>
                <div v-if="oracleIsConnected">
                    <el-row>
                        <allow-write :oracle="oracle" :baseGasFee="baseGasFee"/>
                    </el-row>
                    <el-row>
                        <create :oracle="oracle" :baseGasFee="baseGasFee"/>
                    </el-row>
                    <el-row>
                        <put :oracle="oracle" :gasPerByte="gasPerByte" :baseGasFee="baseGasFee"/>
                    </el-row>
                    <el-row>
                        <get :oracle="oracle" :baseGasFee="baseGasFee"/>
                    </el-row>
                    <el-row>
                        <creat-index :oracle="oracle" :baseGasFee="baseGasFee"/>
                    </el-row>
                    <el-row>
                        <search :oracle="oracle" :baseGasFee="baseGasFee"/>
                    </el-row>
                    <el-row>
                        <import-unchain :oracle="oracle" :baseGasFee="baseGasFee"/>
                    </el-row>
                </div>
            </el-main>
        </el-container>
    </el-container>


</template>
<script setup>
import EventListen from "@/components/eventListen.vue";
import {ref, watch} from "vue";
import Create from "@/components/create.vue";
import AllowWrite from "@/components/allowWrite.vue";
import Get from "@/components/get.vue";
import Put from "@/components/put.vue";
import CreatIndex from "@/components/creatIndex.vue";
import Search from "@/components/search.vue";
import ImportUnchain from "@/components/importUnchain.vue";
import { BrowserWalletConnector, useVueDapp } from '@vue-dapp/core'
import { VueDappModal, useVueDappModal } from '@vue-dapp/modal'
import '@vue-dapp/modal/dist/style.css'
import { ElMessage } from 'element-plus'
import {ethers} from "ethers";
import {oracleAbi} from "./oracle";
import {registerAbi,registerAddress} from "./register";


const { addConnectors, isConnected, wallet, disconnect } = useVueDapp()
addConnectors([new BrowserWalletConnector()])


let register
let oracle
let oracleIsConnected = ref(false)
const oracleList = ref([])
const baseGasFee = ref(0)
const gasPerByte = ref(0)
function onClickConnectButton() {
    if (isConnected.value) disconnect()
    else {
        const { open } = useVueDappModal()
        open()
    }

}
async function onClickOracle(oracleName){
    const oracleAddress = await register.getValue(oracleName)
    const provider = new ethers.providers.Web3Provider(window.ethereum)
    const signer = provider.getSigner()
    oracle = new ethers.Contract(oracleAddress, oracleAbi, signer)
    baseGasFee.value = Number((await oracle.baseGasFee())._hex)
    gasPerByte.value = Number((await oracle.gasPerByte())._hex)
    oracleIsConnected.value = true
}

watch(isConnected,async (newv, oldv) => {
    if (wallet.error) {
        ElMessage({
            showClose: true,
            message: 'Oops, We get an ERROR !',
            type: 'error',
        })
    } else if (newv) {
        ElMessage({
            showClose: true,
            message: 'Connect Wallet Success !',
            type: 'success',
        })
        const provider = new ethers.providers.Web3Provider(window.ethereum)
        register = new ethers.Contract(registerAddress, registerAbi, provider)
        oracleList.value = await register.getKeys()
    } else {
        ElMessage({
            showClose: true,
            message: 'Disconnect Wallet Success !',
            type: 'success',
        })
        oracle = null
        oracleIsConnected.value = false
    }
})


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

#oracleList{
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}

#connect_button{
    float: right;
    width: 150px;
    height: 50px;
    position: relative;
    top: 50%;
    background-color: #409EFF;
    transform: translate(-50%, -50%);
    box-shadow: var(--el-box-shadow-dark);
}
#user{
    float: left;
    margin-right: 50px;
    color: white;
    position: relative;
    top: 50%;
    transform: translate(0, -50%);
}

.oracle{
    width: 200px;
    height: 50px;
    margin: 10px;
}
</style>
