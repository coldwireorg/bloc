<template>
  <div class="fx-column" style="gap: 2rem;">
    <TopHandleBar />
    <div class="fx-column suggestion-container">
      <p>Suggestion</p>
      <div class="fx-row scroll-x-container">
        <FileCard name="" />
        <FileCard name="" />
        <FileCard name="" />
        <FileCard name="" />
        <FileCard name="" />
        <FileCard name="" />
        <FileCard name="" />
      </div>
    </div>
    <FileExplorer />
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  name: 'AppPage',
  layout (context) {
    return context.$device.isMobile ? 'mobile' : 'desktop'
  },
  data () {
    return {
      loginForm: {
        username: '',
        password: ''
      }
    }
  },
  methods: {
    ...mapActions({
      login: 'account/login'
    })
  }
}
</script>

<style scoped>
  .appcontent-main-container {
    position: relative;
    margin: 1rem 2rem;
    width: 100%;
    height: 100%;
    overflow: hidden;
  }

  .suggestion-container {
    gap: 1rem;
  }

  .scroll-x-container {
    overflow-x: scroll;
    scrollbar-width: none;
    width: 100%;
    gap: 1rem;
  }
</style>
