<template>
    <h4>Put</h4>
    <el-input v-model="dbName"  placeholder="Please input DbName" clearable/>
    <el-input v-model="colName"  placeholder="Please input ColName" clearable/>
    <el-input
            v-model="data"
            style="width: 240px"
            :autosize="{ minRows: 1, maxRows: 100 }"
            type="textarea"
            placeholder="Please input Data"
    />
    <el-input v-model="value"  placeholder="Please input pay Value" clearable/>
    <h4>WEI</h4>
    <h4 v-if="reqID.valueOf()!=0" style="margin-left: 60px">ReqID: {{reqID}}</h4>
    <el-button type="primary" round @click="callPut">Call</el-button>

</template>

<script lang="ts" setup>

import {computed, onMounted, ref} from 'vue'
import {ethers} from "ethers";
import {ElMessage} from "element-plus";

const handel = defineProps(['oracle','gasPerByte','baseGasFee'])
const oracle = handel.oracle
const gasPerByte = handel.gasPerByte
const baseGasFee = handel.baseGasFee

const dbName = ref('')
const colName = ref('')
const data = ref('')
/*
data:
[
{"name":"Alice", "age": 18},
{"name":"Bob", "age": 19},
{"name":"Albert", "age": 20},
{"name":"Clearance and Steve", "age":18}
]
 */


const value = computed(()=>{
    const dataB = ethers.utils.toUtf8Bytes(data.value)
    const price = (dataB.length)*gasPerByte+baseGasFee
    return price =='0' ? 1000 : price.toString()
})
const reqID = ref(0)


async function callPut(){
    try {
        const v = value.value.toString()
        const amount = ethers.utils.parseUnits(v,"wei")

        const jsonArray = JSON.parse(data.value);
        let outputStr = '';
        jsonArray.forEach((obj, index) => {
            outputStr += JSON.stringify(obj);
            if (index < jsonArray.length - 1) {
                outputStr += '\n';
            }
        });
        const dataB = ethers.utils.toUtf8Bytes(outputStr)
        const tx = await oracle.Put(dbName.value,colName.value,dataB, { value: amount });
        const receipt = await tx.wait();
        reqID.value=Number(receipt.events[0].args[0]._hex)
        ElMessage({
            showClose: true,
            message: 'Call Put Success !',
            type: 'success',
        })
    } catch (error){
        ElMessage({
            showClose: true,
            message: 'Error Call Put'+error.data.message,
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