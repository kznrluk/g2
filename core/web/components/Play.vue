<template>
  <div v-if="state">
    <SelectableCards :handler="OnSelectCard" :pack="state.pack"></SelectableCards>
    <PlayerCards :players="state.players" :is-shown="state.isShown"></PlayerCards>
  </div>
</template>

<script>
export default {
  data() {
    return {
      myId: null,
      state: null,
    }
  },

  methods: {
    OnSelectCard(number) {
      console.log(number)
    },

    // DEBUG
    ctrl() {
      //const d = this.state.players[0].isControllable;
      //this.$set(this.state.players[0], "isControllable", !d);
    },
    show() {
      //const d = this.state.isShown;
      //this.$set(this.state, "isShown", !d);
    }
  },

  beforeMount () {
    const ws = new WebSocket("ws://localhost:8090/ws")

    ws.addEventListener('open', (socket) => {
      console.log('Connected')
      console.log(socket)
      ws.send(JSON.stringify({
        api: "Ping"
      }))
    })

    ws.addEventListener('message', (event) => {
      console.log('Message from server ', event.data);
    })
  }
}
</script>
