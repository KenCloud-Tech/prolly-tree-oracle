<template>
    <div class="common-layout">
        <el-container >
            <el-header>
                <div class="header">
                    <h3>Prolly Tree Indexed Data Oracle for IPLD</h3>
                    <div class="userList">
                        <div id="user" v-if="isConnected">User: {{ wallet.address.slice(0,10)+'...' }}</div>
                        <div id="user" v-if="oracleIsConnected">BaseGasFee: {{ baseGasFee.valueOf() }} Wei</div>
                        <div id="user" v-if="oracleIsConnected">GasPerByte: {{ gasPerByte.valueOf() }} Wei/Byte</div>
                    </div>
                    <el-button id="connect_button" type="primary" @click="onClickConnectButton()">{{ isConnected ? 'Disconnect Wallet' : 'Connect Wallet' }}</el-button>
                </div>
                <VueDappModal dark auto-connect />
            </el-header>
            <el-main>
                <tutorial v-if="!isConnected" />
                <event-listen v-if="oracleIsConnected" :oracle="oracle"/>
                <div v-if="isConnected && !oracleIsConnected" class="CoverBox">
                    <div class="text">
                        <p>
                            A Prolly Tree is a search tree where the number of values stored in each node is determined probabilistically, based on the data which is stored in the tree.<br><br>
                            We introduce the indexing technology of Prolly trees into the IPLD data format and build：<br>
                            1. the golang implementation of probabilistic merkle search trees for IPLD as a library,<br>
                            2. the ipld prolly trees indexer based on this library,One note, our indexing is a separate thing from prolly trees. Prolly trees give us a B+ tree alternative that's content addresssible and determenistic, on top of that we created a document-oriented database engine with collections and indexes.<br>
                            3. the data oracle on Blockchain to interact the IPLD dataset through the indexer.<br><br>
                            Click on Oralce below to begin your journey of discovery.
                        </p>
                        <div v-for="oracleName in oracleList" :key="oracleName.index" class="oracleBox">
                            <!-- <el-button class="oracle" @click="onClickOracle(oracleName)" color="#626aef" :dark="isDark" plain>{{oracleName}}</el-button> -->
                            <el-button class="oracle" @click="onClickOracle(oracleName)" color="#626aef" :dark="isDark" plain>Get Started with Oralce</el-button>
                        </div>
                        <p></p>
                    </div>
                    <div class="videoBox">
                        <img src="./assets/video.jpg" alt="">
                        <div class="play" @click="playVideo">
                            <img style="width: 50px;height: 50px;" src="./assets/play.png" alt="">
                        </div>
                    </div>
                </div>
                <div v-if="videoVisible" class="videoPlay">
                    <el-button class="Close" type="text" circle @click="ClosePlay"><el-icon size="50" color="#939393"><Close /></el-icon></el-button>
                    <iframe width="800" height="450" src="https://www.youtube.com/embed/TblRt1NA39U" title="Efficient P2P Databases with IPLD Prolly Trees - Mauve Signweaver" frameborder="0" allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture; web-share" referrerpolicy="strict-origin-when-cross-origin" allowfullscreen></iframe>
                </div>
                <ul class="nav" v-if="oracleIsConnected">
                    <li :class="{ active: selectedComponent === 'Create' }" @click="selectComponent('Create')">Create</li>
                    <li :class="{ active: selectedComponent === 'Put' }" @click="selectComponent('Put')">Put</li>
                    <li :class="{ active: selectedComponent === 'Get' }" @click="selectComponent('Get')">Get</li>
                    <li :class="{ active: selectedComponent === 'CreateIndex' }" @click="selectComponent('CreateIndex')">CreateIndex</li>
                    <li :class="{ active: selectedComponent === 'getIndex' }" @click="selectComponent('getIndex')">GetRootCid</li>
                    <li :class="{ active: selectedComponent === 'Search' }" @click="selectComponent('Search')">Search</li>
                    <li :class="{ active: selectedComponent === 'Import' }" @click="selectComponent('Import')">Import</li>
                    <li :class="{ active: selectedComponent === 'AllowWrite' }" @click="selectComponent('AllowWrite')">AllowWrite</li>
                </ul>
                <div style="padding: 0px 40px;" v-if="oracleIsConnected">
                    <el-row v-if="selectedComponent === 'Create'">
                        <create :oracle="oracle" :baseGasFee="baseGasFee" />
                    </el-row>
                    <el-row v-if="selectedComponent === 'Put'">
                        <put :oracle="oracle" :gasPerByte="gasPerByte" :baseGasFee="baseGasFee" />
                    </el-row>
                    <el-row v-if="selectedComponent === 'Get'">
                        <get :oracle="oracle" :baseGasFee="baseGasFee" />
                    </el-row>
                    <el-row v-if="selectedComponent === 'CreateIndex'">
                        <creat-index :oracle="oracle" :baseGasFee="baseGasFee" />
                    </el-row>
                    <el-row v-if="selectedComponent === 'Search'">
                        <search :oracle="oracle" :baseGasFee="baseGasFee" />
                    </el-row>
                    <el-row v-if="selectedComponent === 'Import'">
                        <import-unchain :oracle="oracle" :baseGasFee="baseGasFee" />
                    </el-row>
                    <el-row v-if="selectedComponent === 'getIndex'">
                        <getIndex :oracle="oracle" :baseGasFee="baseGasFee" />
                    </el-row>
                    <el-row v-if="selectedComponent === 'AllowWrite'">
                        <allow-write :oracle="oracle" :baseGasFee="baseGasFee" />
                    </el-row>
                </div>
            </el-main>
        </el-container>
    </div>
</template>
<script setup>
import { ref, watch, onMounted } from "vue";
import { BrowserWalletConnector, useVueDapp } from '@vue-dapp/core'
import { VueDappModal, useVueDappModal } from '@vue-dapp/modal'
import { ElMessage } from 'element-plus'
import {ethers} from "ethers";
import {oracleAbi} from "./oracle";
import { registerAbi, registerAddress } from "./register";
import { Close } from '@element-plus/icons-vue'

import EventListen from "@/components/eventListen.vue";
import Create from "@/components/create.vue";
import AllowWrite from "@/components/allowWrite.vue";
import Get from "@/components/get.vue";
import Put from "@/components/put.vue";
import CreatIndex from "@/components/creatIndex.vue";
import Search from "@/components/search.vue";
import ImportUnchain from "@/components/importUnchain.vue";
import tutorial from "@/components/tutorial.vue";
// @ts-ignore
import getIndex from "@/components/getRootCid.vue";
import ConnectCover from "./components/ConnectCover.vue";

import '@vue-dapp/modal/dist/style.css'

const videoVisible = ref(false);

const videoSrc = ref('');

const playVideo = () => {
  videoSrc.value = "https://www.youtube.com/embed/TU6u_T-s68Y?autoplay=1";
  videoVisible.value = true;
}
const ClosePlay = () => {
    videoSrc.value = "https://www.youtube.com/embed/TU6u_T-s68Y?autoplay=0";
    videoVisible.value = false;
}
const { addConnectors, isConnected, wallet, disconnect } = useVueDapp()
addConnectors([new BrowserWalletConnector()])


let register
let oracle
let oracleIsConnected = ref(false)
const oracleList = ref([])
const baseGasFee = ref(0)
const gasPerByte = ref(0)
// 默认显示的组件
const selectedComponent = ref('Create') 


onMounted(() => {
    localStorage.clear();
});

function selectComponent(component) {
    selectedComponent.value = component;
}

function onClickConnectButton() {
    if (isConnected.value) disconnect()
    else {
        const { open } = useVueDappModal()
        open()
    }

}
async function onClickOracle(oracleName){
    const oracleAddress = "0xF177e3dbE1C3b1cba96EEef0E0E5511910735860"
    // const oracleAddress = await register.getValue(oracleName)
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
        oracleList.value = ['Oracle'];
        // const provider = new ethers.providers.Web3Provider(window.ethereum)
        // register = new ethers.Contract(registerAddress, registerAbi, provider)
        // oracleList.value = await register.getKeys()
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
    height: 9vh;
    background-color: #337ecc;
}

.header{
    height: 100%;
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: white
}
.userList{
    height: 100%;
    display: flex;
    align-items: flex-end;
    gap: 70px;
}
/* #body{
    height: 85vh;
} */
.el-main {
    background-color: #F2F6FC;
    padding: 0;
    padding-top: 40px;
    min-height: 91vh;
    box-sizing: border-box;
}
.el-aside{
    width: 28vw;
    background-color: #c8c9cc;
}
.el-row{
    margin: 10px;
}

/* #oracleList{
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
} */
#connect_button{
    width: 150px;
    height: 70%;
    background-color: #409EFF;
    box-shadow: var(--el-box-shadow-dark);
}
#user{
    width:max-content;
    color: white;
}
.CoverBox{
    width: 80vw;
    min-height: 65vh;
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
    align-items: center;
}
.oracle{
    width: 200px;
    height: 50px;
}
.text{
    width: 45%;
}
.text span{
    font-family: "Nunito", sans-serif;
    line-height: 1.4;
    font-weight: 500;
    font-size: 30px !important;
    color: #161c2d;
}
.text p{
    font-family: "Nunito", sans-serif;
    line-height: 1.6;
    font-size: 18px;
    color: #8492a6;
}
.videoBox{
    width: 48%;
    height: 40vh;
    position: relative;
    overflow: hidden;
    border-radius: 15px;
}
.videoBox img{
    width: 100%;
    height: 100%;
    object-fit: contain;
}
.play{
    width: 80px;
    height: 80px;
    border-radius: 50%;
    background-color: #fff;
    position: absolute;
    top:50%;
    left: 50%;
    transform: translate(-50%,-50%);
    cursor: pointer;
}
.play img{
    position: absolute;
    top:50%;
    left: 50%;
    transform: translate(-45%,-50%);
}
.videoPlay{
    position: absolute;
    top: 0;
    right: 0;
    left: 0;
    width: 100%;
    height: 100vh;
    background-color: rgba(0, 0, 0, 0.8);
    z-index: 9999999;
    display: flex;
    justify-content: center;
    align-items: center;
}
.Close{
    position: absolute;
    top: 20px;
    right: 20px;
}
.nav{
    display: flex;
    gap: 80px;
    padding: 20px 40px;
    box-sizing: border-box;
    border-bottom: 2px solid #e0e0e0;
}
.nav li {
    list-style: none;
    cursor: pointer;
}
.nav li.active {
    font-weight: bold;
    color: #3fa2ff;
}
h3{
    width: 9%;
    overflow: visible; 
    white-space: nowrap; 
}

</style>
