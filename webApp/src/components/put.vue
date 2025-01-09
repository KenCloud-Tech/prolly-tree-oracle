<template>
    <div class="title">
        <h4>Put</h4>
        <h4 v-if="reqID.valueOf()!=0" style="margin-left: 60px">ReqID: {{reqID}}</h4>
        <el-button type="primary" round @click="callPut">Call</el-button>
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
        <li>
            <p>Please input Data<span>(String)</span></p>
            <el-input v-model="data" @input="onShow" style="width: 240px" :autosize="{ minRows: 1, maxRows: 100 }" type="textarea" placeholder="Please input Data" />
            <el-button v-show="isShow" type="text" style="margin-left: 10px;" @click="pushContent">Fill in with one click</el-button>
        </li>
        <li>
            <p>Please input pay Value<span>(Number)</span></p>
            <el-input v-model="value"  placeholder="Please input pay Value" clearable/>
            <h4>WEI</h4>
        </li>
    </div>
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
const isShow = ref(true)
/*
data:
[
{"name":"Alice", "age": 18},
{"name":"Bob", "age": 19},
{"name":"Albert", "age": 20},
{"name":"Clearance and Steve", "age":18}
]
 */
onMounted(() => {
    dbName.value = localStorage.getItem('dbName')
    colName.value = localStorage.getItem('colName')
    console.log('putdbName',dbName.value);
})
const onShow = () => {
    if (data.value !== '') {
        isShow.value = false;
    } else {
        isShow.value = true; 
    }
}
const pushContent = () => {
    data.value = '[{"name":"Alice", "age": 18},{"name":"Bob", "age": 19},{"name":"Albert", "age": 20},{"name":"Clearance and Steve", "age":18}]'
    onShow()
}

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
            message: 'Error Call Put'+error.data,
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