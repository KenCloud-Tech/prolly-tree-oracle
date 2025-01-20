<template>
    <div class="title">
        <h4>Creat</h4>
        <h4 v-if="reqID.valueOf()!=0" style="margin-left: 60px">ReqID: {{reqID}}</h4>
        <el-button type="primary" round @click="callCreate">Call</el-button>
    </div>
    <div class="content">
        <li>
            <p>Please input DbName<span>(String)</span></p>
            <el-input v-model="dbName"  placeholder="Please input DbName" @blur="mouseBlur" clearable/>
        </li>
        <li>
            <p>Please input ColName<span>(String)</span></p>
            <el-input v-model="colName"  placeholder="Please input ColName" @blur="mouseBlur" clearable/>
        </li>
        <li>
            <p>Please input PrimaryKey<span>(String)</span></p>
            <el-input v-model="primaryKey"  placeholder="Please input PrimaryKey" clearable/>
        </li>
        <li>
            <p>Please input pay Value<span>(Number)</span></p>
            <el-input v-model="value"  placeholder="Please input pay Value" clearable/>
            <h4>WEI</h4>
        </li>
    </div>
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
const primaryKey = ref('')
const value = ref(baseGasFee)
const reqID = ref(0)
const mouseBlur = () => {
    localStorage.setItem('dbName', dbName.value)
    localStorage.setItem('colName', colName.value)
    console.log('dbName', dbName.value);
    console.log('colName', colName.value);

}
dbName.value = localStorage.getItem('dbName')
colName.value = localStorage.getItem('colName')

async function callCreate(){
    try {
        const amount = ethers.utils.parseUnits(value.value,"wei")
        const tx = await oracle.Create(dbName.value,colName.value,primaryKey.value, { value: amount });
        const receipt = await tx.wait();
        reqID.value=Number(receipt.events[0].args[0]._hex)
        ElMessage({
            showClose: true,
            message: 'Call Create Success !',
            type: 'success',
        })
    } catch (error){
       var msg = error.message;
        if(error.data){
            msg =  JSON.stringify(error.data)
        }
        ElMessage({
            showClose: true,
            message: 'Error Call Put'+msg,
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