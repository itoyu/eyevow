const msgs = [
  {
    id: 1,
    text: 'はじめまして、私はeyevow。'
  },
  {
    id: 2,
    text: 'あなたの「誓い」を応援するためにやってきた。'
  },
  {
    id: 3,
    text: '人類の未来のために、あなたには達成したい目標や叶えたい願いを誓って欲しい。'
  },
  {
    id: 4,
    text: '"Vow" ボタンで、誓いを立てることができるよ。'
  },
  {
    id: 5,
    text: 'もし私のことや、このプロジェクトについて気になったのなら"About"を覗いてみて。'
  },
  {
    id: 6,
    text: 'あと、"Cheer"では他のひとの誓いを応援することができるの。'
  },
  {
    id: 7,
    text: 'それじゃ、あなたと人類が希望に溢れることを願ってる。'
  }
];

export default {
  fetch() { return msgs },
  find(id)  { return msgs.find(el => el.id === id) },
  asyncFind(id, callback) {
    setTimeout(() => {
      callback(msgs.find(el => el.id === id))
    }, 100)
  }
}
