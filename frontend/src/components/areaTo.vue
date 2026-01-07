<template>
  <areaBlank @click="update" text="to..." />
</template>

<script setup>
import areaBlank from './areaBlank.vue'

const app = window.go.backend.App
const update = () => {
  app.SelectDst().then((data) => {
    dst.path = data.path
    console.log("The data is : ", data)
  })
}
</script>
