<template>
    <h4>Search</h4>
    <el-input v-model="dbName"  placeholder="Please input DbName" clearable/>
    <el-input v-model="colName"  placeholder="Please input ColName" clearable/>
    <el-input
            v-model="query"
            style="width: 240px"
            :autosize="{ minRows: 1, maxRows: 100 }"
            type="textarea"
            placeholder="Please input Data"
    />
    <el-input v-model="value"  placeholder="Please input pay Value" clearable/>
    <h4>WEI</h4>
    <h4 v-if="reqID.valueOf()!=0" style="margin-left: 60px">ReqID: {{reqID}}</h4>
    <el-button type="primary" round @click="callSearch">Call</el-button>

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
const query = ref('')
const value = ref(baseGasFee)
const reqID = ref(0)
/*
query:
[
{
  "Equal": {
    "IndexName": "name",
    "IndexVal": "Bob"
  },
  "Compare": {
    "Op": "",
    "IndexName": "",
    "IndexVal": null
  },
  "Sort": "",
  "Limit": 0,
  "Skip": 0
}
]
 */
async function callSearch(){
    try {
        const amount = ethers.utils.parseUnits(value.value,"wei")
        const dataB = ethers.utils.toUtf8Bytes(query.value)

        const tx = await oracle.Search(dbName.value,colName.value,dataB,'', { value: amount });
        const receipt = await tx.wait();
        reqID.value=Number(receipt.events[0].args[0]._hex)
        ElMessage({
            showClose: true,
            message: 'Call Search Success !',
            type: 'success',
        })
    } catch (error){
        ElMessage({
            showClose: true,
            message: 'Error Call Search'+error.data.message,
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