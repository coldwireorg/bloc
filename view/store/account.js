/* eslint-disable no-console */
export const account = () => ({
  token: null,
  user: null,
  currentDir: {
    id: null,
    path: null
  },
  suggestFiles: [],
  workdirFiles: []
})

export const mutations = {
  setUserInfos (state, { token, user }) {
    state.token = token
    state.user = user
    state.currentDir = user.root
  }
}

export const actions = {
  async login ({ commit }, _payload) {
    await this.$axios({
      url: '/api/auth/login',
      method: 'POST',
      withCredential: true,
      headers: {
        'Content-Type': 'application/json'
      },
      data: _payload
    }).then((res) => {
      console.dir(res.response.content)
    }).catch((err) => {
      if (err.response.status === 300) {
        console.log(err.response.data.content)
        commit('setUserInfos', err.response.data.content)
      }
    })
  },

  async logout () {
    //
  },

  async infos () {
    //
  },

  async register () {
    //
  },

  async fetchFileUpload ({ commit, state }, _payload) {
    const _v = new Uint32Array(25.6)
    const _entropy = crypto.getRandomValues(_v).buffer
    const _key = [...new Uint8Array(_entropy)].map(x => x.toString(16).padStart(2, '0')).join('')

    const _data = new FormData()
    _data.append('file', _payload.file)
    _data.append('parent', state.currentDir)
    _data.append('key', _key)
    await this.$axios({
      url: '/api/file',
      method: 'POST',
      withCredential: true,
      headers: {
        'Content-Type': 'multipart/form-data'
      },
      data: _data
    }).then((res) => {
      console.dir(res)
    }).catch((err) => {
      console.dir(err)
    })
  }
}

export const getters = {
  getCurrentDir (state) {
    return state.currentDir
  }
}
