<template>
  <div>
    <SelectableCards :handler="OnSelectCard" :pack="state.pack"></SelectableCards>
    <PlayerCards :players="state.players" :is-shown="state.isShown"></PlayerCards>
    <Control v-if="state.players[0].isControllable"></Control>
    <button @click="ctrl">CONTROL: {{ state.players[0].isControllable }}</button>
    <button @click="show">IS_SHOWN: {{ state.isShown }}</button>
  </div>
</template>

<script>
const yourId = "hogehuga";
const STATE = {
  roomNumber: 19283,
  selectablePack: [ "Mountain Goat pack" ],
  pack: ["0", "1/2", "1", "2", "3", "5", "8", "13", "20", "40", "100", "?", "☕️"],
  players: [
    {
      id: "hogehuga",
      isSelected: false,
      selectedNumber: "3",
      isControllable: true,
    },
    {
      id: "hogehuga2",
      isSelected: true,
      selectedNumber: "9",
      isControllable: true,
    },
    {
      id: "hogehuga3",
      isSelected: false,
      selectedNumber: "300",
      isControllable: true,
    },
  ],
  isShown: false,
}

export default {
  data() {
    return {
      myId: yourId,
      state: { ...STATE },
    }
  },

  methods: {
    OnSelectCard(number) {
      console.log(number)
    },

    // DEBUG
    ctrl() {
      const d = this.state.players[0].isControllable;
      this.$set(this.state.players[0], "isControllable", !d);
    },
    show() {
      const d = this.state.isShown;
      this.$set(this.state, "isShown", !d);
    }
  },

  mounted () {
    console.log('HI!')
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
