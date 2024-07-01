<template>
    <h4>Import</h4>
    <el-input v-model="dbName"  placeholder="Please input DbName" clearable/>
    <el-input v-model="colName"  placeholder="Please input ColName" clearable/>
    <el-input v-model="url"  placeholder="Please input Url" clearable/>
    <el-input v-model="type"  placeholder="Please input Type" clearable/>
    <el-input v-model="value"  placeholder="Please input pay Value" clearable/>
    <h4>WEI</h4>
    <h4 v-if="reqID.valueOf()!=0" style="margin-left: 60px">ReqID: {{reqID}}</h4>
    <el-button type="primary" round @click="callImport">Call</el-button>

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
const url = ref('')
const type = ref('')
const value = ref(baseGasFee)
const reqID = ref(0)
/*
url:
	urlcsv    = "http://127.0.0.1:8080/data.csv"
	urlndjson = "http://127.0.0.1:8080/ndjson"
 */
async function callImport(){
    try {
        const amount = ethers.utils.parseUnits(value.value,"wei")
        const tx = await oracle.Import_from_url(dbName.value,colName.value,url.value,type.value, { value: amount });
        const receipt = await tx.wait();
        reqID.value=Number(receipt.events[0].args[0]._hex)
        ElMessage({
            showClose: true,
            message: 'Call Import Success !',
            type: 'success',
        })
    } catch (error){
        ElMessage({
            showClose: true,
            message: 'Error Call Import'+ error.toString(),
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
    width: 180px;
    float: left;
    margin: 0 10px;
}
.el-button{
    width: 100px;
    position: absolute;
    right: 10px;
}
</style>