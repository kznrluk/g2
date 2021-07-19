<template>
  <div>
    <ol class="list">
      <li v-for="card in $props.pack" :key="card">
        <Card :on-click-handler="CreateCardClickHandler(card)" :content="card" :is-shown="true"></Card>
      </li>
    </ol>
  </div>
</template>

<script>
export default {
  props: {
    handler: {
      type: Function,
      default() {
        return function () {
          console.log("no prop")
        };
      }
    },
    pack: {
      type: Array, // <String>
      default() { return [] }
    },
  },
  methods: {
    CreateCardClickHandler(num) {
      return () => {
        this.$props.handler(num)
      }
    }
  }
}
</script>

<style>
.list {
  list-style: none;
  display: flex;
}
</style>
