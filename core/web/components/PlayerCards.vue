<template>
  <div>
    <ol class="list">
      <li v-for="player of $props.players" :key="player.id">
        <Card
          :content="player.isSelected ? isShown ? player.selectedNumber : '✔︎' : '?' "
          :is-highlighted="player.isSelected && !isShown"
          :is-shown="isShown">
        </Card>
      </li>
    </ol>
  </div>
</template>

<script>
/**
 * {
 *   id: "hogehuga",
 *   isSelected: false,
 *   selectedNumber: null,
 *   isControllable: true,
 * }
 */

export default {
  props: {
    players: {
      type: Array,
      default() { return [] }
    },
    isShown: {
      type: Boolean,
      default() {
        return false;
      }
    }
  },
}
</script>

<style>
.list {
  list-style: none;
  display: flex;
}
</style>
