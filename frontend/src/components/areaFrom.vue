<template>
  <areaBlank 
    v-model="params"
    @click="app.SelectSrc().then()"
    text="кликните или перетащите базу-источник"
  />
</template>

<script setup>
import areaBlank from './areaBlank.vue'

const app = window.go.backend.App
const src = defineModel()
console.log(src)

const params = defineModel('params')
console.log(params)


//const select = () => { app.SelectSrc().then() }
</script>
