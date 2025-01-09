<template>
    <!-- <h4>GetIndex</h4>
    <el-input v-model="dbName"  placeholder="Please input DbName" clearable/>
    <el-input v-model="colName"  placeholder="Please input ColName" clearable/>
    <el-input v-model="index"  placeholder="Please input Index" clearable/>
    <el-input v-model="value"  placeholder="Please input pay Value" clearable/>
    <h4>WEI</h4>
    <h4 v-if="reqID.valueOf()!=0" style="margin-left: 60px">ReqID: {{reqID}}</h4>
    <el-button type="primary" round @click="callGetIndex">Call</el-button> -->
    <div class="title">
        <h4>GetRootCid</h4>
        <h4 v-if="reqID.valueOf()!=0" style="margin-left: 60px">ReqID: {{reqID}}</h4>
        <el-button type="primary" round @click="callGetRootCid">Call</el-button>
    </div>
    <div class="content">
        <li>
            <p>Please input DbName<span>(String)</span></p>
            <el-input v-model="dbName"  placeholder="Please input DbName" clearable/>
        </li>
        <li>
            <p>Please input ColName<span>(String)</span></p>
            <el-input v-model="colName"  placeholder="Please input ColName" clearable/>
        </li>
        <!-- <li>
            <p>Please input Index<span>(String)</span></p>
            <el-input v-model="index"  placeholder="Please input Index" clearable/>
        </li> -->
        <li>
            <p>Please input pay Value<span>(Number)</span></p>
            <el-input v-model="value"  placeholder="Please input pay Value" clearable/>
            <h4>WEI</h4>
        </li>
    </div>
</template>

<script lang="ts" setup>

import {ref,defineEmits,onMounted} from 'vue'
import {ethers} from "ethers";
import {ElMessage} from "element-plus";

const handel = defineProps(['oracle','baseGasFee'])
const oracle = handel.oracle
const baseGasFee = handel.baseGasFee.toString()

const dbName = ref('')
const colName = ref('')
// const index = ref('')
const value = ref(baseGasFee)
const reqID = ref(0)
const emits = defineEmits(['isOk'])
onMounted(() => {
    dbName.value = localStorage.getItem('dbName')
    colName.value = localStorage.getItem('colName')
})

async function callGetRootCid(){
    try {
        const amount = ethers.utils.parseUnits(value.value,"wei")
        // const tx = await oracle.GetIndex(dbName.value,colName.value,index.value, { value: amount });
        const tx = await oracle.GetRootCid(dbName.value, colName.value, { value: amount });
        const receipt = await tx.wait();
        reqID.value=Number(receipt.events[0].args[0]._hex)
        ElMessage({
            showClose: true,
            message: 'Call GetRootCid Success !',
            type: 'success',
        })
        const dataself={
                id:reqID,
                flagStauts:true,
                inforstauts:'Call GetRootCid Success !'
            }
        emits('isOk',dataself)
    } catch (error){
        ElMessage({
            showClose: true,
            message: 'Error Call GetRootCid'+ error,
            type: 'error',
        })
        const dataself={
                id:reqID,
                flagStauts:false,
                inforstauts:'Error Call GetRootCid'+ error
            }
        emits('isOk',dataself)
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