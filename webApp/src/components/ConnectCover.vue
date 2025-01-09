<template>
  <div class="contentBox">

  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'

</script>

<style scoped>
.contentBox{
    width: 80vw;
    margin: 0 auto;
    display: flex;
    justify-content: space-between;
}
.text{
    width: 45%;
    height: 100%;
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
    margin-bottom: 15%;
}
</style>