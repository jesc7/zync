<template>
  <q-layout view="hHh lpR fFf" class="bg-grey-4" style="min-height: 100vh;">
    <q-page-container 
      class="q-mx-md text-grey-8"
      style="
        height: 100vh; 
        font: bold 26px courier;"
    >
      <q-page>
        <div class="row flex-center" style="height: 80px;">
          <appBanner />
        </div>

        <div 
          style="
            height:calc(100vh - 80px);
            user-select: none;"
        >
          <div class="flex q-pb-md" style="height:100%;">

          <areaFrom 
            v-model="src" 
            v-model:params="srcParams" 
          />

          <div style="width: 3%;"></div>

          <areaFrom 
            v-model="src" 
            v-model:params="dstParams" 
          />
        </div>
      </div>

      </q-page>
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref } from 'vue'
import { colors } from 'quasar'
import appBanner from './appBanner.vue'
import areaFrom from './areaFrom.vue'
import areaTo from './areaTo.vue'

const app = window.go.backend.App
const src = ref({
  path:'',
  tables:[],
  profile:0,
  agents:false,
})
app.Src().then(v => { src.value = v })
window.runtime.EventsOn('src.changed', v => { src.value = v; console.log("runtime.EventsOn:", src.value) })

//watch(src, v => { console.log(`watch src path:`, v.value?.path) })

const { getPaletteColor } = colors
const srcParams = ref({
  blank: {
    color: getPaletteColor('blue-8')
  },
})

const dstParams = ref({
  blank: {
    color: getPaletteColor('red-8')
  },
})

//const changeSrc = () => { app.GetSrc().then(v => { src.value = v }) }
//changeSrc()
</script>
