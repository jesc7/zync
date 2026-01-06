<template>
  <div
    id="blank"
    class="column justify-center bg-grey-3"
    style="
      height:100%;
      width: 48.5%;
      border-radius: 20px;
      align-items: center;
      user-select: none;
    "
  >
    <p id="plus" style="margin: 0; font-size: 2.5rem;">+</p>
    <p id="hint" style="margin: 0; height:20px; font-size: 0.65rem; font-family: Arial;">{{ props.text }}</p>
  </div>
</template>

<script setup>
const app = window.go.backend.App
//const blank = defineModel()
const props = defineProps(['text'])

//const root = document.querySelector(":root");
//root.style.setProperty("--hover-color", blank.value.color)
</script>

<style scoped lang="scss">
:root {
  --hover-color: #44f;
}

#blank:hover {
  cursor: pointer;
  border-color: var(--hover-color);
  color: var(--hover-color);
}

#blank:hover > #hint {
  visibility: visible;
}

#blank {
  border: 3px dashed;
  border-color: $grey-6;
  color: $grey-6;
}

#blank > #hint {
  visibility: hidden;
}
</style>
