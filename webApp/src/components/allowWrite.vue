<template>
    <div class="title">
        <h4>Address：Allow the user's wallet address</h4>
        <h4 v-if="reqID.valueOf()!=0" style="margin-left: 60px">ReqID: {{reqID}}</h4>
        <el-button type="primary" round @click="callAllow">Call</el-button>
    </div>
    <div class="content">
        <li>
            <p>Please input DbName<span>(String)</span></p>
            <el-input v-model="dbName"  placeholder="Please input DbName" @blur="mouseBlur" clearable/>
        </li>
        <li>
            <p>Please input Address<span>(String)</span></p>
            <el-input v-model="addr"  placeholder="Please input Address" clearable/>
        </li>
        <li>
            <p>Please input pay Value<span>(Number)</span></p>
            <el-input v-model="value"  placeholder="Please input pay Value" clearable/>
            <h4>WEI</h4>
        </li>
    </div>
</template>

<script lang="ts" setup>

import {ref,onMounted} from 'vue'
import {ethers} from "ethers";
import {ElMessage} from "element-plus";

const handel = defineProps(['oracle','baseGasFee'])
const oracle = handel.oracle
const baseGasFee = handel.baseGasFee.toString()
const addr = ref('0xDF5Ec19a07F5Fd4136658Af5a1F7D531C8fEA842')
const dbName  = ref('')
const value = ref(baseGasFee)
const reqID = ref(0)


onMounted(() => {
    dbName.value = localStorage.getItem('dbName')
})

async function callAllow(){
    try {
        const amount = ethers.utils.parseUnits(value.value,"wei")
        const tx = await oracle.AllowWrite(dbName.value, addr.value, { value: amount });
        const receipt = await tx.wait();
        reqID.value=Number(receipt.events[0].args[0]._hex)
        console.log("reqID.value",reqID.value);
        
        ElMessage({
            showClose: true,
            message: 'Call AllowWrite Success !',
            type: 'success',
        })
    } catch (error){
        ElMessage({
            showClose: true,
            message: 'Error Call AllowWrite'+error.data.message,
            type: 'error',
        })
    }
}


</script>

<style scoped>

.el-button{
    width: 100px;
}
.el-input{
    width: 240px;
    height: 30px;
    border: none !important;
}
.title{
    width: 100%;
    display: flex;
    justify-content: space-between;
    align-items: center;
    height: 50px;
    border-bottom: 1px solid #e0e0e0;
    box-shadow: 0 8px 6px rgba(0, 0, 0, 0.1);
}
.content{
    padding: 30px;
}
.content > li{
    display: flex;
    align-items: center;
    list-style: none;
    gap:10px;
}
.content > li p{
    width: 400px;
    white-space: pre-wrap;
}
.content > li p span{
    color: #a1a1a1;
    white-space: pre-wrap;
}
</style>