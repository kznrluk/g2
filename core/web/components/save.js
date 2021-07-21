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

console.log(STATE, yourId)
