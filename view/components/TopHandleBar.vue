<template>
  <div class="fx-row top-handle-bar">
    <a href="#">My files</a>
    <div class="fx-row right-content">
      <SearchBar />
      <input type="file" name="file" @change.prevent="handleFileEvent">
      <input type="hidden" name="parent" value="">
      <input type="hidden" name="key" value="">
      <ButtonAction text="Transfer" color="blue" @click.native="login({ username: 'jeff', password: 'azertyuio' })">
        <Icon name="upload" />
      </ButtonAction>
    </div>
  </div>
</template>

<script>
import { mapActions } from 'vuex'

export default {
  name: 'TopHandleBar',
  methods: {
    ...mapActions({
      login: 'account/login',
      fetchFileUpload: 'account/fetchFileUpload'
    }),
    handleFileEvent (e) {
      this.fetchFileUpload({ file: e.target.files[0] })
    }
  }
}
</script>

<style scoped>
  .top-handle-bar {
    box-sizing: border-box;
    width: 100%;
    padding: 0;
    color: transparent;
    align-items: center;
    justify-content: space-between;
  }

  .right-content {
    gap: 1rem;
  }

  a {
    font-style: normal;
    font-weight: 600;
    font-size: 14px;
    line-height: 17px;
  }
</style>
