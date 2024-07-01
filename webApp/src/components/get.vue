<template>
    <h4>Get</h4>
    <el-input v-model="dbName"  placeholder="Please input DbName" clearable/>
    <el-input v-model="colName"  placeholder="Please input ColName" clearable/>
    <el-input v-model="recordID"  placeholder="Please input RecordID" clearable/>
    <el-input v-model="value"  placeholder="Please input pay Value" clearable/>
    <h4>WEI</h4>
    <h4 v-if="reqID.valueOf()!=0" style="margin-left: 60px">ReqID: {{reqID}}</h4>
    <el-button type="primary" round @click="callGet">Call</el-button>

</template>

<script lang="ts" setup>

import {ref} from 'vue'
import {ethers} from "ethers";
import {ElMessage} from "element-plus";

const handel = defineProps(['oracle','baseGasFee'])
const oracle = handel.oracle
const baseGasFee = handel.baseGasFee.toString()

const dbName = ref('')
const colName = ref('')
const recordID = ref('')
const value = ref(baseGasFee)
const reqID = ref(0)
/*
recordID:
[129, 99, 66, 111, 98]
 */
async function callGet(){
    try {
        const amount = ethers.utils.parseUnits(value.value,"wei")
        let byteArray = JSON.parse(recordID.value);
        let dataB = new Uint8Array(byteArray);
        const tx = await oracle.Get(dbName.value,colName.value,dataB,'', { value: amount });
        const receipt = await tx.wait();
        reqID.value=Number(receipt.events[0].args[0]._hex)
        ElMessage({
            showClose: true,
            message: 'Call Get Success !',
            type: 'success',
        })
    } catch (error){
        ElMessage({
            showClose: true,
            message: 'Error Call Get'+error.data.message,
            type: 'error',
        })
    }
}


</script>

<style scoped>
h4{
    margin: 5px 0;
}
.el-input{
    width: 200px;
    float: left;
    margin: 0 10px;
}
.el-button{
    width: 100px;
    position: absolute;
    right: 10px;
}
</style>