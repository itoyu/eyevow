const db = [
  {
    id: 1,
    name: 'ITEM01',
    price: 100,
    desc: 'Description 01 ...'
  },
  {
    id: 2,
    name: 'ITEM02',
    price: 200,
    desc: 'Description 02 ...'
  },
  {
    id: 3,
    name: 'ITEM03',
    price: 300,
    desc: 'Description 03 ...'
  }
];

export default {
  fetch(id) { return db },
  find(id)  { return db.find(el => el.id === id) },
  asyncFind(id, callback) {
    setTimeout(() => {
      callback(db.find(el => el.id === id))
    }, 100)
  }
}
