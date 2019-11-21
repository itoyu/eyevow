

class Vows {
  constructor(db) {
    this.mem = [];
  }

  async ofUser(user_id) {
    for (let v of this.mem) {
      if (v.user_id === user_id) {
        return v;
      }
    }
    return null;
  }

  async post(user_id, text, eyevow) {
    const id = shortid.generate();

    const v = {
      id, user_id,
      text,
      eyevow,
      archived: false, cheer_count: 0
    };

    this.mem.push(v);

    return v;
  }
}

async function init() {
  
  
  module.exports.vows = new Vows(db);
}

module.exports = {
  init,
  vows: null
};