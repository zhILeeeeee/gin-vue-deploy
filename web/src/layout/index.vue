<script setup>

import SideBar from "@/layout/components/sideBar.vue";
</script>

<template>
  <div class="page-container">
    <side-bar class="sideBar"></side-bar>
    <div
        class="main-container"
    >
    </div>

  </div>
</template>

<style scoped>

.page-container {
  width: 100%;
  height: 100vh;
  position: relative;
  overflow-y: hidden;
  background-color: #eff1f2;

  .sideBar {
    transition: width 0.28s;
    height: 100%;
    position: fixed;
    top: 0;
    bottom: 0;
    left: 0;
    display: flex;
    flex-direction: column;
    overflow: hidden;
    //background: linear-gradient(137deg, #000000, #2f2f2f 97%);
    border-radius: 0 24px 24px 0;
  }


}

</style>