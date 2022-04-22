<template>
  <div />
</template>

<script>
export default {
  name: 'IndexPage',
  layout: 'default',
  beforeMount () {
    this.$axios.get('/api/user/auth/oauth2').then((resp) => {
      if (resp.data.data.enabled) {
        window.location.replace(resp.data.data.redirect_url)
      } else {
        this.$router.push('/auth')
      }
    })
  }
}
</script>
