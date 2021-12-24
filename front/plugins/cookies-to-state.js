export default ({app, store}) => {
  const token = app.$cookies.get('jwt-token')
  if (token) {
    store.dispatch('setToken', topken)
  }
}
