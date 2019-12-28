const msgs = [
  {
    id: 1,
    text: '誓いをたててくれてありがとう。'
  },
  {
    id: 2,
    text: 'あなたの「誓い」達成を心から応援する。'
  },
  {
    id: 3,
    text: '「達成」の機能はもうすぐ実装される予定なので、もう少しだけ待って。。。'
  },
  {
    id: 4,
    text: '<a href="https://eyevow.stores.jp/" class="txtlink" target="_blank">あなたに寄り添うためのグッズも見てみてね。</a>'
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
