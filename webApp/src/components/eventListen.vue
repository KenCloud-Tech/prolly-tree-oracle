<template>
<h4 style="text-align: center">EVENT LISTENING...</h4>
  <div style="margin: 20px">
      Received event ReqState with:
      <div style="background-color: white; height: 180px; width: 100%; white-space: pre-wrap;">
          {{ReqState}}
      </div>
  </div>
    <div style="margin: 20px">
        Received event CatchData with:
        <div style="background-color: white; height: 180px; width: 100%; white-space: pre-wrap;">
            {{CatchData}}
        </div>
    </div>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue';
import {ethers} from "ethers";
const ReqState = ref('')
const CatchData = ref('')
const { oracle } = defineProps<{ oracle: any }>();

onMounted(() => {
    oracle.on("ReqState", (reqID, user, statement, info, event) => {
        ReqState.value = `Request ID: ${reqID.toString()}\nUser: ${user}\nStatement: ${statement}\nInfo: ${info}`
    });
    oracle.on("CatchData", (reqID, data, event) => {

        const byteArray = hexToUint8Array(data).slice(1);
        const decoder1 = new TextDecoder('utf-8');
        const stringData = decoder1.decode(byteArray);

        const jsonData = JSON.parse(stringData);
        let str = JSON.stringify(jsonData)
        
        str = str.replace(/\\n/g, '')
            .replace(/\\"/g, '"')
            .replace(/\},\{/g, '},\n{')
        CatchData.value = `Request ID: ${reqID.toString()}\nData: ${str}`
    });

});
function hexToUint8Array(hex) {
    const array = new Uint8Array(Math.ceil(hex.length / 2));
    for (let i = 0; i < hex.length; i += 2) {
        array[Math.floor(i / 2)] = parseInt(hex.substring(i, i + 2), 16);
    }
    return array;
}
</script>

<style scoped>
</style>