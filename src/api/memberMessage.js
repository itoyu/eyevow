const msgs = [
  {
    id: 1,
    text: '誓いをたててくれてありがとう'
  },
  {
    id: 2,
    text: 'あなたの「誓い」、達成を願っている'
  },
  {
    id: 3,
    text: '<a href="https://thebase.in/" class="txtlink" target="_blank">ECサイト</a>のグッズもチェックしてね'
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
