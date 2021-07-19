<template>
  <button class="card" @click="$props.onClickHandler">
    {{ $props.content }}
  </button>
</template>

<script>

export default {
  props: {
    isHighlighted: Boolean,
    content: {
      type: String,
      default: "0"
    },
    onClickHandler: {
      type: Function,
      default() {
        return function () {};
      }
    },
  }
}

</script>

<style>
.card {
  width: 64px;
  height: 103px;
  background-color: beige;
  font-size: 1.5rem;
  margin: 5px;
}
</style>
