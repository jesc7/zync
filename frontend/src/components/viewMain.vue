<template>
    <div></div>
</template>

<script setup>
import { onMounted, onUnmounted, ref } from 'vue'
import { EventsOn, EventsOff } from '@wailsjs/runtime/runtime'
import { useQuasar } from 'quasar'

const $q = useQuasar()
const time = ref()

onMounted(() => {
  // Подписываемся на событие "changeTime"
  EventsOn('changeTime', () => {
    time.value = runtime
    
    // Показываем уведомление Quasar при каждом событии
    $q.notify({
      message: `Получены данные из Go: ${data}`,
      color: 'positive',
      timeout: 500
    })
  })
})

</script>
