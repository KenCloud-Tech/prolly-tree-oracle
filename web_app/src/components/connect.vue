<template>


    <div id="address" v-if="isConnected">User: {{ wallet.address.slice(0,8)+'...' }}</div>
    <div id="address" v-if="isConnected">BaseGasFee: {{ baseGasFee.valueOf() }} Wei</div>
    <div id="address" v-if="isConnected">GasPerByte: {{ gasPerByte.valueOf() }} Wei/Byte</div>
    <el-button id="connect_button" type="primary" @click="onClickConnectButton()">{{ isConnected ? 'Disconnect Wallet' : 'Connect Wallet' }}</el-button>
    <VueDappModal dark auto-connect />

</template>

<script lang="ts" setup>
import { BrowserWalletConnector, useVueDapp } from '@vue-dapp/core'
import { VueDappModal, useVueDappModal } from '@vue-dapp/modal'
import '@vue-dapp/modal/dist/style.css'
import {abi,oracle_address} from "../oracle";
import { ethers } from "ethers";
import { ElMessage } from 'element-plus'
import {ref, watch} from "vue";

const handle = defineProps(['sendOracle','isConnected'])

const { addConnectors, isConnected, wallet, disconnect } = useVueDapp()
addConnectors([new BrowserWalletConnector()])

function onClickConnectButton() {
    if (isConnected.value) disconnect()
    else {
        const { open } = useVueDappModal()
        open()
    }

}

const baseGasFee = ref(0)
const gasPerByte = ref(0)

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
        const signer = provider.getSigner()
        const oracle = new ethers.Contract(oracle_address, abi, signer)
        baseGasFee.value = Number((await oracle.baseGasFee())._hex)
        gasPerByte.value = Number((await oracle.gasPerByte())._hex)
        handle.sendOracle(oracle, true,baseGasFee.value,gasPerByte.value )
    } else {
        ElMessage({
            showClose: true,
            message: 'Disconnect Wallet Success !',
            type: 'success',
        })
        handle.sendOracle(null, false)
    }
})



</script>
<style scoped>

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
#address{
    float: left;
    margin-right: 50px;
    color: white;
    position: relative;
    top: 50%;
    transform: translate(0, -50%);
}
</style>